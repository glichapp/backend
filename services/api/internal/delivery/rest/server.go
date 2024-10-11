package rest

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvizyx/glich/services/api/internal/config"
	"github.com/kvizyx/glich/shared/logger"
)

type HTTPServer struct {
	server *http.Server
	router *gin.Engine

	logger logger.StructuralLogger
}

func NewHTTPServer(
	config config.Config,
	logger logger.StructuralLogger,
) HTTPServer {
	router := gin.New()

	server := &http.Server{
		Addr: fmt.Sprintf(
			"%s:%d",
			config.HTTP.Host, config.HTTP.Port,
		),
		Handler:           router.Handler(),
		ReadTimeout:       config.HTTP.ReadTimeout,
		ReadHeaderTimeout: config.HTTP.ReadHeaderTimeout,
		WriteTimeout:      config.HTTP.WriteTimeout,
		IdleTimeout:       config.HTTP.IdleTimeout,
		MaxHeaderBytes:    config.HTTP.MaxHeaderBytes,
	}

	return HTTPServer{
		server: server,
		router: router,
		logger: logger,
	}
}

func (s HTTPServer) Start() error {
	s.logger.Info("http server is listening", slog.String("addr", s.server.Addr))

	if err := s.server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("listen and serve: %w", err)
		}
	}

	return nil
}

func (s HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
