package mysql

import "testing"

func TestLastArticle(t *testing.T) {
	CreateArticleDao().LastArticle()
}
