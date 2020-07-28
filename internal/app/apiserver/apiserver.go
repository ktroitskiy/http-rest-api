package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (server *APIServer) Start() error {
	if error := server.configureLogger(); error != nil {
		return error
	}

	server.configureRouter()

	server.logger.Info("starting api server")

	return http.ListenAndServe(server.config.BindAddress, server.router)
}

func (server *APIServer) configureLogger() error {
	level, error := logrus.ParseLevel(server.config.LogLevel)

	if error != nil {
		return error
	}

	server.logger.SetLevel(level)

	return nil
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/hello", server.handleHello())
}

func (server *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello")
	}
}
