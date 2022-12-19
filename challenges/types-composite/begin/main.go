// challenges/types-composite/begin/main.go
package main

import (
	"fmt"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	key := b.author.name
	l[key] = append(l[key], b)

}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(author string) []book {
	books, ok := l[author]
	if ok == false {
		return nil
	}
	return books
}

func main() {
	// create a new library
	var l = library{}

	// add 2 books to the library by the same author
	l.addBook(book{title: "Book 1", author: author{name: "author a"}})
	l.addBook(book{title: "Book 2", author: author{name: "author a"}})

	fmt.Println(l)

	// add 1 book to the library by a different author
	l.addBook(book{title: "Book 3", author: author{name: "author c"}})

	// dump the library with spew
	// spew.Dump(l)

	// lookup books by known author in the library
	lp := l.lookupByAuthor("author a")

	fmt.Println(lp)

	// print out the first book's title and its author's name
	if len(lp) != 0 {
		fmt.Println(lp[0].title, lp[0].author.name)
	}

}
