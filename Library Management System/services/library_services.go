package services

import (
	"fmt"
	"task3/models"
	"bufio"
	"strconv"
	"strings"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

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


type Library struct{
	Books []models.Book
	Members []models.Member
}

func (l *Library)AddBook(book models.Book) {
	l.Books = append(l.Books, book)
}

func (l *Library)RemoveBook(ID int) {
	index := -1
	for i,book := range l.Books{
		if book.ID == ID && book.Status == "borrowed"{
			fmt.Println("The book is currently borrowed")
			return
		}
		if book.ID == ID{
			index = i
			break
		}
	}

	if index == -1{
		return
	}

	l.Books = append(l.Books[:index], l.Books[index+1:]...)
}

func (l *Library)BorrowBook(bookID, memberID int) string {
	for _, book := range l.Books{
		if bookID == book.ID {
			if book.Status == "borrowed" {
				return "The book is already borrowed."
			}
			book.Status  = "borrowed"
			
			for _,member := range l.Members{
				if member.ID == memberID{
					member.BorrowedBooks = append(member.BorrowedBooks, book)
					return "Book added seccusfully"
				}
			}
			
		}
	}

	var member models.Member

	fmt.Println("Enter member Name: ")
	memberName := ReadLine()
	
	book := CreateBook()
	member.ID = memberID
	member.Name = memberName
	member.BorrowedBooks = []models.Book{book}
	return "Book with new user added seccusfully"
}

func (l *Library)ReturnBook(bookID, memberID int) string{
	for _,book := range l.Books{
		if bookID == book.ID{
			book.Status = "available"

			for _,member := range l.Members{
				if memberID == member.ID{
					for i,b := range member.BorrowedBooks{
						if b.ID == bookID{
							member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]... )
						}
					}
				}
			}
			return "Book returned seccusfully"
			
		}
	}
	return "No such book with that id borowwed"
}

func (l *Library)ListAvailableBooks() []models.Book{
	var availableBooks []models.Book

	for _,b := range l.Books{
		if b.Status == "available"{
			availableBooks = append(availableBooks, b)
		}
	}
	return availableBooks
}

func (l *Library)ListBorrowedBooks(memberID int) []models.Book{
	for _,member := range l.Members{
		if memberID == member.ID{
			return member.BorrowedBooks
		}
	}
	var brwd models.Book
	brwd.ID = 0
	brwd.Author = ""
	brwd.Status = ""
	brwd.Title = ""
	
	var bs []models.Book
	bs = append(bs, brwd)
	return bs
}

