package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{ // In a real application, this list will be fetched from a database or from static files
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
	article{ID: 3, Title: "Article 3", Content: "Article 3 body"},
}

func getAllArticles() []article { // Returns a list of all the articles
	return articleList
}

func getArticleByID(id int) (*article, error) { // Returns an article speficied by an ID
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}
