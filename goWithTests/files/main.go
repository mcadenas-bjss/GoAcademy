package main

import (
	"log"
	"os"

	"github.com/mcadenas-bjss/GoAcademy/goWithTests/files/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFs(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
