package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/damiannolan/eros/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// ErrShutdownFailed is the error returned when the Server failed to shutdown gracefully
var ErrShutdownFailed = errors.New("ErrShutdownFailed: Failed to shutdown http server. Is the server running?")

// Server encapsulates a HTTP Server lifecycle
type Server interface {
	RegisterMiddleware(middleware ...func(http.Handler) http.Handler)
	RegisterResource(resource Resource)
	Run() error
	Shutdown() error
}

// Resource encapsulates a logical collection of handlers to be bound to a HTTP Server
type Resource interface {
	Path() string
	Routes() http.Handler
}

type server struct {
	config   *Config
	httpSrv  *http.Server
	router   chi.Router
	signalCh chan os.Signal
}

// New is a constructor func which creates and returns a new Server instance
func New(cfg *Config) Server {
	srv := &server{
		config:   cfg,
		router:   chi.NewRouter(),
		signalCh: make(chan os.Signal, 1),
	}

	srv.initialize()
	srv.registerShutdownHook()

	return srv
}

// RegisterMiddleware registers an arbitary number of middleware handlers on the Server
func (s *server) RegisterMiddleware(middleware ...func(http.Handler) http.Handler) {
	s.router.Use(middleware...)
}

// RegisterResource registers a Resource on the server router
func (s *server) RegisterResource(resource Resource) {
	s.router.Mount(s.config.BasePath+resource.Path(), resource.Routes())
}

// Run starts a HTTP server using the configured Server router
func (s *server) Run() error {
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Port),
		Handler: s.router,
	}

	log.WithFields(log.Fields{"name": s.config.Name, "port": s.config.Port}).Info("Starting Server")
	return s.httpSrv.ListenAndServe()
}

// Shutdown gracefully shuts down the HTTP Server and closes it's signal channel
func (s *server) Shutdown() error {
	if s.httpSrv == nil {
		return ErrShutdownFailed
	}

	if err := s.httpSrv.Shutdown(context.Background()); err != nil {
		log.WithError(err).Error("HTTP Server shutdown error")
	}

	log.Info("Server shutdown successfully")
	close(s.signalCh)
	s.httpSrv = nil
	return nil
}

// initialize registers a default set of middleware
func (s *server) initialize() {
	s.RegisterMiddleware(
		middleware.RequestID, // Not necessary - Can write better
		middleware.Logger,    // Not necessary
		middleware.StripSlashes,
		middleware.Recoverer,
	)
}

// registerShutdownHook registers the server signalCh and awaits to be notified of termination signals
func (s *server) registerShutdownHook() {
	signal.Notify(s.signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-s.signalCh
		s.Shutdown()
	}()
}
