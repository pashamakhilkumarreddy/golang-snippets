package main

import (
	"log"

	"github.com/pashamakhilkumarreddy/golang-snippets/composition/blog"
)

func main() {
	log.Println("Composition in Golang")
	author1 := blog.NewAuthor("Gwen", "Stacy", "Ultra Instinct")
	blogPost1 := blog.NewPost("Ultra Instinct", "The next level power", author1)
	blogPost1.Details()

	w := blog.NewWebsite(blogPost1, blogPost1)
	w.Contents()
}
