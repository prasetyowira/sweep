package sweep

import (
	"context"
	"database/sql"
	"net/http"

	entsql "github.com/facebookincubator/ent/dialect/sql"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opencensus"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/goph/idgen/ulidgen"
	"github.com/gorilla/mux"
	appkitendpoint "github.com/sagikazarmark/appkit/endpoint"
	appkithttp "github.com/sagikazarmark/appkit/transport/http"
	"github.com/sagikazarmark/kitx/correlation"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
	kitxtransport "github.com/sagikazarmark/kitx/transport"
	kitxhttp "github.com/sagikazarmark/kitx/transport/http"

	"github.com/prasetyowira/sweep/internal/app/sweep/httpbin"
	"github.com/prasetyowira/sweep/internal/app/sweep/landing/landingdriver"
	"github.com/prasetyowira/sweep/internal/app/sweep/product"
	"github.com/prasetyowira/sweep/internal/app/sweep/product/productadapter"
	"github.com/prasetyowira/sweep/internal/app/sweep/product/productadapter/ent"
	"github.com/prasetyowira/sweep/internal/app/sweep/product/productadapter/ent/migrate"
	"github.com/prasetyowira/sweep/internal/app/sweep/product/productdriver"
)

// InitializeApp initializes a new HTTP and a new gRPC application.
func InitializeApp(
	httpRouter *mux.Router,
	db *sql.DB,
	logger Logger,
	errorHandler ErrorHandler,
) {
	endpointMiddleware := []endpoint.Middleware{
		correlation.Middleware(),
		opencensus.TraceEndpoint("", opencensus.WithSpanName(func(ctx context.Context, _ string) string {
			name, _ := kitxendpoint.OperationName(ctx)

			return name
		})),
		appkitendpoint.LoggingMiddleware(logger),
	}

	transportErrorHandler := kitxtransport.NewErrorHandler(errorHandler)

	httpServerOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transportErrorHandler),
		kithttp.ServerErrorEncoder(kitxhttp.NewJSONProblemErrorEncoder(appkithttp.NewDefaultProblemConverter())),
		kithttp.ServerBefore(correlation.HTTPToContext()),
	}

	{
		client := ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))
		err := client.Schema.Create(
			context.Background(),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		)
		if err != nil {
			panic(err)
		}

		store := productadapter.NewEntStore(client)

		service := product.NewService(
			ulidgen.NewGenerator(),
			store,
		)
		service = productdriver.LoggingMiddleware(logger)(service)

		endpoints := productdriver.MakeEndpoints(
			service,
			kitxendpoint.Combine(endpointMiddleware...),
		)

		productdriver.RegisterHTTPHandlers(
			endpoints,
			httpRouter.PathPrefix("/product").Subrouter(),
			kitxhttp.ServerOptions(httpServerOptions),
		)
	}

	landingdriver.RegisterHTTPHandlers(httpRouter)
	httpRouter.PathPrefix("/httpbin").Handler(http.StripPrefix(
		"/httpbin",
		httpbin.MakeHTTPHandler(logger.WithFields(map[string]interface{}{"module": "httpbin"})),
	))
}
