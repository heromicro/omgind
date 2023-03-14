package asyncq

import (
	"github.com/hibiken/asynq"
)

//

func NewAsynqClient() (*asynq.Client, func(), error) {

	client := asynq.NewClient()

	cleanFunc := func() {
		client.Close()
	}

	return client, cleanFunc, nil
}

type Handler struct {
	Pattern string
	asynq.Handler
}

func NewHandler(pattern string, handler asynq.Handler) *Handler {
	return &Handler{Pattern: pattern, Handler: handler}
}
func NewHandlerFunc(pattern string, handler asynq.HandlerFunc) *Handler {
	return &Handler{Pattern: pattern, Handler: handler}
}

func RegisterHandlers(srv *Server, handlers ...*Handler) {
	for _, handler := range handlers {
		srv.Handle(handler.Pattern, handler.Handler)
	}
}
