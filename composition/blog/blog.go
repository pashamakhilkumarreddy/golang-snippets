package blog

import (
	"fmt"
)

type author struct {
	firstName, lastName, bio string
}

func NewAuthor(firstName, lastName, bio string) author {
	a := author{firstName, lastName, bio}
	return a
}

func (a author) FullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title, content string
	author
}

func NewPost(title, content string, a author) blogPost {
	b := blogPost{title, content, a}
	return b
}

func (b blogPost) Details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.FullName()) // b.author.FullName()
	fmt.Println("Bio: ", b.author.bio)
}

type website struct {
	Posts []blogPost // not possible to anonymously name a slice like []blogPost
}

func NewWebsite(posts ...blogPost) website {
	w := website{
		Posts: posts,
	}
	return w
}

func (w website) Contents() {
	fmt.Printf("\nContent of Websites are as follows\n")
	fmt.Println()
	for _, v := range w.Posts {
		v.Details()
		fmt.Println()
	}
}
