package main

import "testing"

func TestGetAllArticles(t *testing.T) { // Test the function that fetches all articles
	alist := getAllArticles()

	if len(alist) != len(articleList) { // Check that the length of the list of articles returned is the same as the length of the global variable holding the list
		t.Fail()
	}

	for i, v := range alist { // Check that each member is identical
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
