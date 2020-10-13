package productdriver

import (
	"context"
	"encoding/json"
	"net/http"

	"emperror.dev/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	appkithttp "github.com/sagikazarmark/appkit/transport/http"
	kitxhttp "github.com/sagikazarmark/kitx/transport/http"

	api "github.com/prasetyowira/sweep/.gen/api/openapi/product/go"
)

// RegisterHTTPHandlers mounts all of the service endpoints into a router.
func RegisterHTTPHandlers(endpoints Endpoints, router *mux.Router, options ...kithttp.ServerOption) {
	errorEncoder := kitxhttp.NewJSONProblemErrorResponseEncoder(appkithttp.NewDefaultProblemConverter())

	router.Methods(http.MethodPost, http.MethodOptions).Path("").Handler(kithttp.NewServer(
		endpoints.CreateProduct,
		decodeCreatProductHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeCreateProductHTTPResponse, errorEncoder),
		options...,
	))

	router.Methods(http.MethodGet).Path("").Handler(kithttp.NewServer(
		endpoints.ListProducts,
		kithttp.NopRequestDecoder,
		kitxhttp.ErrorResponseEncoder(encodeListProductsHTTPResponse, errorEncoder),
		options...,
	))

	router.Methods(http.MethodGet).Path("/{id}").Handler(kithttp.NewServer(
		endpoints.GetProduct,
		decodeGetProductHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeGetProductHTTPResponse, errorEncoder),
		options...,
	))

	router.Methods(http.MethodGet).Path("/sku/{sku}").Handler(kithttp.NewServer(
		endpoints.GetProductBySKU,
		decodeGetProductBySKUHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeGetProductHTTPResponse, errorEncoder),
		options...,
	))
}

func decodeCreatProductHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var apiRequest api.CreateProductRequest

	err := json.NewDecoder(r.Body).Decode(&apiRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request")
	}

	return CreateProductRequest{
		Sku: apiRequest.Sku,
		Name: apiRequest.Name,
		Expirable: apiRequest.Expirable,
	}, nil
}

func encodeCreateProductHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(CreateProductResponse)

	apiResponse := api.CreateProductResponse{
		Id: resp.Id,
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, kitxhttp.WithStatusCode(apiResponse, http.StatusCreated))
}

func encodeListProductsHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(ListProductsResponse)

	apiResponse := api.ProudctList{}

	for _, product := range resp.Products {
		apiResponse.Products = append(apiResponse.Products, api.Product{
			Id:   product.ID,
			Sku:   product.SKU,
			Name:   product.Name,
			Expirable: product.Expirable,
		})
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, apiResponse)
}

func decodeGetProductHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok || id == "" {
		return nil, errors.NewWithDetails("missing parameter from the URL", "param", "id")
	}

	return GetProductRequest{
		Id: id,
	}, nil
}

func decodeGetProductBySKUHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	sku, ok := vars["sku"]
	if !ok || sku == "" {
		return nil, errors.NewWithDetails("missing parameter from the URL", "param", "sku")
	}

	return GetProductBySKURequest{
		Sku: sku,
	}, nil
}

func encodeGetProductHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetProductResponse)

	apiResponse := api.Product{
		Id: resp.Product.ID,
		Sku: resp.Product.SKU,
		Name: resp.Product.Name,
		Expirable: resp.Product.Expirable,
	}

	return kitxhttp.JSONResponseEncoder(ctx, w, apiResponse)
}
