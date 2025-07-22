📦 Package: controllers
This package provides the user interface logic for interacting with the library management system from the command line. It handles input/output, menu display, and user-driven operations such as adding, removing, borrowing, and returning books.

🔧 Global Variables
go
Copy
Edit
var reader = bufio.NewReader(os.Stdin)
var Library1 services.Library
reader: Used to read user input from the console.

Library1: Instance of the Library struct from the services package. It stores and manages all book and member records.

🔢 func IntReader() int
Reads a line of input from the user and converts it into an int.

Returns: Parsed integer value entered by the user.

Used in: Menu selection, entering IDs.

🔤 func ReadLine() string
Reads a line of input from the user and returns it as a string after trimming whitespace.

Returns: Trimmed string input from the user.

Used in: Capturing book title, author, and status.

🧾 func SellectYOurChoice()
Displays the main menu options to the user. This function is called at the beginning to show the available operations.

Menu Options:

Add Book

Remove Book

Borrow Book

Return Book

List Available Books

List Borrowed Books

Exit

📖 func CreateBook() models.Book
Prompts the user to enter the details of a new book and constructs a Book struct.

Returns: A models.Book object with user-provided data.

Prompts for: ID, Title, Author, Status.

🔁 func GoForChoice()
Core loop that handles user input and performs the selected operation from the menu.

Reads user's numeric choice.

Executes the corresponding function from the Library1 object:

AddBook

RemoveBook

BorrowBook

ReturnBook

ListAvailableBooks

ListBorrowedBooks

Loops back to the menu after each operation unless the user chooses "Exit".

🧠 func Manipulate()
Initial entry point to the controller logic.

Calls SellectYOurChoice() to display the menu.

Invokes GoForChoice() to begin processing user inputs and looping until exit.

📌 Usage in main.go
To launch the program, the main package should call:

go
Copy
Edit
func main() {
    controllers.Manipulate()
}
✅ Example Interaction
bash
Copy
Edit
Welcome to the Library Management System
1. Add Book
2. Remove Book
...
Enter your choice:
> 1
Book ID:
> 101
Book Title:
> Programming
Book Author:
> Lench
Book Status:
> available

📦 Package: services
The services package provides the core business logic for managing a library. It defines a Library struct and various methods to manipulate books and members, including adding, borrowing, returning, and listing books.

🧱 Structs
🏛️ Library
go
Copy
Edit
type Library struct {
    Books   []models.Book
    Members []models.Member
}
Represents the entire library system.

Books: List of all books in the library.

Members: List of all registered library members.

🧩 Utility Functions
🔢 func IntReader() int
Reads an integer from standard input.

Returns: int — the parsed user input.

🔤 func ReadLine() string
Reads a single line of text from the user input.

Returns: string — trimmed user input.

📘 func CreateBook() models.Book
Prompts the user to input book details and creates a Book object.

Returns: A models.Book with ID, Title, Author, and Status fields set.

📚 Library Methods
➕ func (l *Library) AddBook(book models.Book)
Adds a new book to the library’s book list.

Parameters: book — a models.Book object.

➖ func (l *Library) RemoveBook(ID int)
Removes a book by ID if it’s not currently borrowed.

Parameters: ID — the book's unique ID.

Note: If the book is currently borrowed, removal is blocked and a message is shown.

🔄 Borrowing and Returning
📖 func (l *Library) BorrowBook(bookID, memberID int) string
Handles borrowing a book by a member.

If book is borrowed: Returns a message indicating so.

If book is available and member exists: Adds the book to their borrowed list.

If member doesn't exist: Prompts for name and creates a new member.

Returns: A message indicating the outcome.

📗 func (l *Library) ReturnBook(bookID, memberID int) string
Handles returning a borrowed book.

Finds the book and sets its status to "available".

Removes the book from the member’s BorrowedBooks list.

Returns: Success message or error if the book isn’t found.

📃 Listing Books
✅ func (l *Library) ListAvailableBooks() []models.Book
Lists all books in the library with status "available".

Returns: A slice of models.Book available for borrowing.

📕 func (l *Library) ListBorrowedBooks(memberID int) []models.Book
Returns a list of books currently borrowed by the specified member.

If member found: Returns their BorrowedBooks.

If member not found: Returns a default empty book in a slice.

⚠️ Design Suggestions
The current implementation modifies struct fields of slices (book.Status, member.BorrowedBooks) by value. However, since slices return copies, this may not persist changes outside the loop. Consider:

Using pointers to books and members in []*Book, []*Member.

Refactoring loops to update actual objects via pointers.

Example:

go
Copy
Edit
for i := range l.Books {
    if l.Books[i].ID == bookID {
        l.Books[i].Status = "borrowed"
    }
}git branch -M main
git push -u origin main
