package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gambarini/articleapi/internal/repo"
)

type (
	Server struct {
		*http.Server
		Repo repo.IArticleRepository
	}
)

func NewServer(addr string, db repo.IArticleRepository) *Server {

	router := mux.NewRouter()

	server := &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		Repo: db,
	}

	HandleArticle(server)
	HandleTags(server)

	return server
}

func (srv *Server) Router() *mux.Router {
	return srv.Handler.(*mux.Router)
}
