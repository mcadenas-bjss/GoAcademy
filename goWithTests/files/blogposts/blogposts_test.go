package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/mcadenas-bjss/GoAcademy/goWithTests/files/blogposts"
	"github.com/mcadenas-bjss/GoAcademy/goWithTests/files/post"
)

func TestNewBlogPost(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], post.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func assertPost(t *testing.T, got post.Post, want post.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
