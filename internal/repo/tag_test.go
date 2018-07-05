package repo

import (
	"testing"
	"os/exec"
	"github.com/gambarini/articleapi/internal/model"
	"time"
	"reflect"
)

func TestArticleRepository_FindTag(t *testing.T) {

	StartMongoDB()
	defer StopMongoDB()

	repo, err := NewArticleRepository()
	if err != nil {
		t.Fatal(err)
	}

	PopulateArticles(repo)

	tests := []struct {
		name       string
		searchTag  string
		searchDate string
		expCount   int
		expRTags   []string
		expIDs     []string
	}{
		{name: "11 Articles for A", searchDate: "2010-10-10", searchTag: "A", expCount: 11, expRTags: []string{"B", "C", "D"}, expIDs: []string{"11", "10", "9", "8", "7", "6", "5", "4", "3", "2"}},
		{name: "2 Articles for B", searchDate: "2010-10-10", searchTag: "B", expCount: 2, expRTags: []string{"A", "C"}, expIDs: []string{"10", "2"}},
		{name: "No Articles for E", searchDate: "2010-10-10", searchTag: "E", expCount: 0, expRTags: []string{}, expIDs: []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resTag, _ := repo.FindTag(test.searchTag, test.searchDate)
			if err != nil {
				t.Log(err)
				t.Fail()
			}

			if resTag.Count != test.expCount {
				t.Logf("Exp count: %d", test.expCount)
				t.Logf("Res count: %d", resTag.Count)
				t.Fail()
			}

			if !testRTags(test.expRTags, resTag.RelatedTags) {
				t.Logf("Exp rTags: %v", test.expRTags)
				t.Logf("Res rTags: %v", resTag.RelatedTags)
				t.Fail()
			}

			if !reflect.DeepEqual(resTag.Articles, test.expIDs) {
				t.Logf("Exp IDs: %v", test.expIDs)
				t.Logf("Res IDs: %v", resTag.Articles)
				t.Fail()
			}
		})
	}
}

func testRTags(exp []string, res []string) bool {

	if len(exp) != len(res) {
		return false
	}

	var r bool
	for _, tExp := range exp {
		r = false
		for _, tRes := range res {
			if tExp == tRes {
				r = true
				break
			}
		}
		if !r {
			return false
		}
	}

	return true
}

func PopulateArticles(repo *ArticleRepository) {
	date := time.Date(2010, 10, 10, 10, 0, 0, 0, time.UTC)
	repo.Store(model.Article{ID: "1", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "2", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "B"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "3", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "C"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "4", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "D"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "5", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "D"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "6", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "D"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "7", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "D"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "8", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "D"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "9", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "D"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "10", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "B", "C"}, Body: "Body", Title: "Tile"})
	date = date.Add(time.Second)
	repo.Store(model.Article{ID: "11", Date: date.Format("2006-01-02"), DateTime: date, Tags: []string{"A", "D"}, Body: "Body", Title: "Tile"})
}

func StartMongoDB() error {

	cmdStr := "docker run -p 27017:27017 --name mongodb_articleapi_test -d mongo:4.0.0"

	_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()

	if err != nil {
		return err
	}

	return nil

}

func StopMongoDB() error {

	cmdStr := "docker stop mongodb_articleapi_test"
	_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		return err
	}

	cmdStr = "docker rm mongodb_articleapi_test"
	_, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		return err
	}

	return nil
}
