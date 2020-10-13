package watermill

import (
	"time"

	"emperror.dev/errors"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	watermilllog "logur.dev/integration/watermill"
	"logur.dev/logur"
)

// RouterConfig holds information for configuring Watermill router.
type RouterConfig struct {
	CloseTimeout time.Duration
}

// NewRouter returns a new sweep router for sweep subscription logic.
func NewRouter(config RouterConfig, logger logur.Logger) (*message.Router, error) {
	h, err := message.NewRouter(
		message.RouterConfig{
			CloseTimeout: config.CloseTimeout,
		},
		watermilllog.New(logur.WithField(logger, "component", "watermill")),
	)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create sweep router")
	}

	retryMiddleware := middleware.Retry{}
	retryMiddleware.MaxRetries = 1
	retryMiddleware.MaxInterval = time.Millisecond * 10

	h.AddMiddleware(
		// if retries limit was exceeded, sweep is sent to poison queue (poison_queue topic)
		retryMiddleware.Middleware,

		// recovered recovers panic from handlers
		middleware.Recoverer,

		// correlation ID middleware adds to every produced sweep correlation id of consumed sweep,
		// useful for debugging
		middleware.CorrelationID,
	)

	return h, nil
}
