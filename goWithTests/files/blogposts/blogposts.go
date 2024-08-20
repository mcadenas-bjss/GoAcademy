package blogposts

import (
	"io/fs"

	"github.com/mcadenas-bjss/GoAcademy/goWithTests/files/post"
)

func NewPostsFromFs(fileSystem fs.FS) ([]post.Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []post.Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, f string) (post.Post, error) {
	postFile, err := fileSystem.Open(f)
	if err != nil {
		return post.Post{}, err
	}
	defer postFile.Close()
	return post.NewPost(postFile)
}
