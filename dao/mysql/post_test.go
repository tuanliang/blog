package mysql

import (
	"blog/models"
	"testing"
)

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          10,
		AuthorID:    12,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("createpost insert record into mysql failed,err:%v\n", err)
	}
	t.Logf("createpost insert record into mysql success")
}
