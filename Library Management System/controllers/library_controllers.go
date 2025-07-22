package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task3/models"
	"task3/services"
)

var reader = bufio.NewReader(os.Stdin)
var Library1 services.Library

func IntReader() int {
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)
	result, _ := strconv.Atoi(value)
	return result
}

func ReadLine() string {
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)
	return value
}

func SellectYOurChoice() {
	fmt.Println("Welcome to the Library Management System")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")
	fmt.Println("7. Exit")
}

func CreateBook() models.Book {
	var book models.Book
	fmt.Println("Book ID: ")
	ID := IntReader()
	book.ID = ID

	fmt.Println("Book Title: ")
	Title := ReadLine()
	book.Title = Title

	fmt.Println("Book Auther: ")
	Author := ReadLine()
	book.Author = Author

	fmt.Println("Book Status: ")
	Status := ReadLine()
	book.Status = Status

	return book
}


func GoForChoice(){
	fmt.Println("Enter your choice: ")
	choice := IntReader()
	switch choice {
	case 1:
		book := CreateBook()
		Library1.AddBook(book)
		GoForChoice()

	case 2:
		fmt.Println("Enter BookID: ")
		bookID := IntReader()
		Library1.RemoveBook(bookID)
		GoForChoice()

	case 3:
		fmt.Println("Enter BookID: ")
		bookID := IntReader()

		fmt.Println("Enter MemberID: ")
		memberID := IntReader()

		msg := Library1.BorrowBook(bookID, memberID)
		fmt.Println(msg)
		GoForChoice()

	case 4:
		fmt.Println("Enter BookID: ")
		bookID := IntReader()

		fmt.Println("Enter MemberID: ")
		memberID := IntReader()

		msg := Library1.ReturnBook(bookID, memberID)
		fmt.Println(msg)
		GoForChoice()

	case 5:
		books := Library1.ListAvailableBooks()
		fmt.Println(books)
		GoForChoice()
	
	case 6:
		fmt.Println("Enter memberID")
		memberID := IntReader()
		bbooks := Library1.ListBorrowedBooks(memberID)

		fmt.Println(bbooks)
		GoForChoice()

	case 7:
		fmt.Println("Good Bye!")
		return
	}
}

func Manipulate() {
	SellectYOurChoice()
	GoForChoice()
}
