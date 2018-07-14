package server

import (
	"testing"
	"github.com/gambarini/articleapi/internal/model"
	"net/http/httptest"
	"net/http"
	"fmt"
)

type (
	ArticleRepositoryMock struct {
	}
)

func (repo *ArticleRepositoryMock) Store(article model.Article) error {
	return nil
}
func (repo *ArticleRepositoryMock) Find(id string) (article model.Article, err error) {
	return model.Article{}, nil
}
func (repo *ArticleRepositoryMock) FindTag(tag, date string) (resultTag model.Tag, err error) {

	key := fmt.Sprintf("%s-%s", tag, date)

	return tags[key], nil
}

var (
	tags = make(map[string]model.Tag)
)

func Test_getTag(t *testing.T) {

	srv := NewServer("", &ArticleRepositoryMock{})

	tags["B-2010-10-10"] = model.Tag{
		RelatedTags: []string{"A"},
		Articles: []string{"1"},
		Tag: "B",
		Count: 1,
	}

	tests := []struct {
		name string
		url string
		expCode int
	}{
		{name: "Tag found", url: "/tags/B/2010-10-10", expCode: 200},
		{name: "Tag not found", url: "/tags/A/2010-10-10", expCode: 404},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			req, _ := http.NewRequest("GET", test.url, nil)

			w := httptest.NewRecorder()

			srv.Handler.ServeHTTP(w, req)

			if w.Code != test.expCode {
				t.Logf("Exp code: %d", test.expCode)
				t.Logf("Res code: %d", w.Code)
				t.Fail()
			}

		})
	}
}
