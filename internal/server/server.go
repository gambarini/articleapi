package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gambarini/articleapi/internal/repo"
)

type (
	Server struct {
		*http.Server
		Repo *repo.ArticleRepository
	}
)

func NewServer(addr string, db *repo.ArticleRepository) *Server {

	server := &Server{
		Server: &http.Server{
			Addr: addr,
		},
		Repo: db,
	}

	r := mux.NewRouter()

	http.Handle("/", r)

	HandleArticle(r, server)
	HandleTags(r, server)

	return server
}
