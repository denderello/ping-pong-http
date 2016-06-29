package server

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"gopkg.in/tylerb/graceful.v1"

	"github.com/denderello/ping-pong-http/log"
)

type ServerConfig struct {
	Addr             Addresser
	ShutdownCooldown time.Duration
	Logger           log.Logger
	EnableProfiling  bool
}

type Server struct {
	addr             Addresser
	shutdownCooldown time.Duration
	logger           log.Logger
	router           *mux.Router
	gracefulServer   *graceful.Server
}

func New(sc ServerConfig) *Server {
	r := mux.NewRouter()

	// Register the default http ServeMux for everything under /debug/pprof/debug/pprof  and let Go handle the rest internally
	if sc.EnableProfiling {
		r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	}

	h := handlers.CombinedLoggingHandler(&log.WriterBridge{Logger: sc.Logger}, r)
	return &Server{
		addr:             sc.Addr,
		shutdownCooldown: sc.ShutdownCooldown,
		logger:           sc.Logger,
		router:           r,
		gracefulServer: &graceful.Server{
			NoSignalHandling: true,
			Server: &http.Server{
				Addr:    sc.Addr.Address(),
				Handler: h,
			},
		},
	}
}

func (s *Server) RegisterAPI(api API) {
	s.logger.Debugf("Registering handlers for API %s at path %s", api.Name(), api.PathPrefix())

	r := s.router.PathPrefix(api.PathPrefix()).Subrouter()
	api.RegisterHandlers(r)
}

func (s *Server) Start() error {
	s.logger.Infof("Starting server at %s", s.addr.Address())
	return s.gracefulServer.ListenAndServe()
}

func (s *Server) Stop() {
	s.logger.Info("Stopping server")
	s.gracefulServer.Stop(s.shutdownCooldown)
	s.logger.Debug("Server stopped")
}
