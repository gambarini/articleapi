package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gambarini/articleapi/internal/repo"
	"golang.org/x/net/http2"
	"log"
)

type (
	Server struct {
		*http.Server
		Repo repo.IArticleRepository
	}
)

func NewServer(addr string, db repo.IArticleRepository) *Server {

	router := mux.NewRouter()

	httpServer := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	err := http2.ConfigureServer(httpServer, nil)

	if err != nil {
		log.Fatalf("Error configuring http2 server, %s", err)
	}

	server := &Server{
		Server: httpServer,
		Repo: db,
	}

	HandleArticle(server)
	HandleTags(server)

	return server
}

func (srv *Server) Router() *mux.Router {
	return srv.Handler.(*mux.Router)
}
