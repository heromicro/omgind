package asyncq

import (
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
)

type RedisConfigFunc func() any

func (r RedisConfigFunc) MakeRedisClient() any {
	return r()
}

func NewAsynqClientOpt(r redis.UniversalClient) asynq.RedisConnOpt {
	return RedisConfigFunc(func() any {
		return r
	})
}

func NewAsynqClient(opt asynq.RedisConnOpt) (*asynq.Client, func(), error) {

	client := asynq.NewClient(opt) 

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
