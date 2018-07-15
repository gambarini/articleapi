package repo

import (
	"github.com/gambarini/articleapi/internal/model"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func (repo *ArticleRepository) FindTag(tag, date string) (resultTag model.Tag, err error) {

	session, collection := repo.articleCollection()

	defer session.Close()

	query := bson.M{"$and":
	[]bson.M{
		{"tags": bson.M{"$in": []string{tag}}},
		{"date": date},
	},
	}

	var articles []model.Article
	count := make(chan int)
	tags := make(chan []string)

	go queryCount(collection, query, count)

	go queryTags(collection, query, tag, tags)

	err = collection.Find(query).Select(bson.M{"_id": 1}).Sort("-datetime").Limit(10).All(&articles)

	if err != nil {
		return resultTag, err
	}

	resultTag = model.Tag{
		Tag:         tag,
		Articles:    []string{},
		RelatedTags: []string{},
	}

	for _, art := range articles {

		resultTag.Articles = append(resultTag.Articles, art.ID)

	}

	resultTag.Count = <-count
	resultTag.RelatedTags = <-tags

	return resultTag, nil
}

func queryCount(collection *mgo.Collection, query bson.M, count chan int) {
	c, _ := collection.Find(query).Count()
	count <- c
	close(count)
}

func queryTags(collection *mgo.Collection, query bson.M, tag string, tags chan []string) {
	var ts []string
	resultTs := make([]string, 0)

	collection.Find(query).Distinct("tags", &ts)

	for _, t := range ts {
		if t != tag {
			resultTs = append(resultTs, t)
		}
	}

	tags <- resultTs
	close(tags)
}
