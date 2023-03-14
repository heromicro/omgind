package asyncq

import (
	"context"

	"github.com/go-saas/lazy"
	"github.com/hibiken/asynq"
)

type Server struct {
	*asynq.ServeMux
	server LazyAsynqServer
}

type ServerOption func(opt *asynq.Config)

func NewAsynqServer(opt asynq.RedisConnOpt, opts ...ServerOption) *Server {
	ret := &Server{server: newAsynqServer(opt, opts...), ServeMux: asynq.NewServeMux()}
	return ret
}

func (s *Server) Start(ctx context.Context) error {
	asynqServer, err := s.server.Value(ctx)
	if err != nil {
		return err
	}
	return asynqServer.Start(s)
}

func (s *Server) Stop(ctx context.Context) error {
	asynqServer, err := s.server.Value(ctx)
	if err != nil {
		return err
	}
	asynqServer.Shutdown()
	return nil
}

type LazyAsynqServer lazy.Of[*asynq.Server]

func newAsynqServer(opt asynq.RedisConnOpt, opts ...ServerOption) LazyAsynqServer {
	return lazy.New(func(ctx context.Context) (*asynq.Server, error) {
		cfg := asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{},
			BaseContext: func() context.Context {
				return ctx
			},
			// See the godoc for other configuration options
		}
		for _, option := range opts {
			option(&cfg)
		}
		return asynq.NewServer(
			opt,
			cfg,
		), nil
	})

}
