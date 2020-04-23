package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main1()  {
	fmt.Println(Books{"go语言", "www", "语言教程", 6454407})

	fmt.Println(Books{title:"go语言", author:"www", subject:"语言教程", book_id:6454407})
}

func main()  {
	var Book1 Books
	var Book2 Books

	Book1.title = "gogo"
	Book1.author = "chen"
	Book1.subject = "go language"
	Book1.book_id = 67688834

	fmt.Println(Book1)

	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	printBook(Book2)
}

func printBook(books1 Books)  {
	fmt.Printf( "Book 2 title : %s\n", books1.title)
	fmt.Printf( "Book 2 author : %s\n", books1.author)
	fmt.Printf( "Book 2 subject : %s\n", books1.subject)
	fmt.Printf( "Book 2 book_id : %d\n", books1.book_id)
}