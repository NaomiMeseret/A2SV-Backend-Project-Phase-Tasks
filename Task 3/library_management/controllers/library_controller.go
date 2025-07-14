package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)
var reader = bufio.NewReader(os.Stdin)
func ShowMenu(lib *services.Library){
	for{
		fmt.Println("\n📔 Library Menu:")
		fmt.Println("1. Add book")
		fmt.Println("2. Remove book ")
		fmt.Println("3. Borrow book")
		fmt.Println("4. Return book")
		fmt.Println("5. List available books")
		fmt.Println("6. List borrowed books")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice:")
		input, _:= reader.ReadString('\n')
		input= strings.TrimSpace(input)
		choice , err  := strconv.Atoi(input)
		if err!=nil{
			fmt.Println(("Invaild input. please enter a number"))
			continue
		}
		switch choice {
		case 1:
			addBook(lib)
		case 2:
			removeBook(lib)
		case 3:
			borrowBook(lib)
		case 4:
			returnBook(lib)
		case 5:
			listAvailableBooks(lib)
		case 6:
			listBorrowedBooks(lib)
		case 7:
			fmt.Println("👋 See you next time")
			return
		default:
			fmt.Println("Invaild Option")
			
		}
	}
}
func addBook(lib *services.Library){
	fmt.Print("Enter book title:")
	title, _:= reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print("Enter name of author:")
	author,_:=reader.ReadString('\n')
	author = strings.TrimSpace(author)
	book :=models.Book{
		Title : title,
		Author: author,
		Status: "Available",
	}
	lib.AddBook(book)
	fmt.Println("✅ Book added")
}
func removeBook(lib *services.Library){
	fmt.Print("Enter book ID to remove :")
	id , _:=reader.ReadString('\n')
	id = strings.TrimSpace(id)
	bookID , err:= strconv.Atoi(id)
	if err!=nil{
			fmt.Println("Invaild input. please enter a  vaildnumber")
			return
		}
	err=lib.RemoveBook(bookID)
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("✅ Book removed")
	}

}
func borrowBook(lib * services.Library){
	fmt.Println("Available Books")
	books:=lib.ListAvailableBooks()
	for _ , b:=range books{
		fmt.Printf("ID: %d |Title:%s | Author:%s\n", b.ID, b.Title,b.Author)
	}
	fmt.Print("Enter a book ID you want to borrow:")
	id , _:=reader.ReadString('\n')
	id = strings.TrimSpace(id)
	bookID, err:=strconv.Atoi(id)
	if err!=nil{
			fmt.Println(("Invaild input. please enter a  vaild number"))
			return
		}
	fmt.Print("Enter member ID:")
	memberIDStr ,_:=reader.ReadString('\n')
	memberIDStr =strings.TrimSpace(memberIDStr)
	memberID ,err:=strconv.Atoi(memberIDStr)
	if err!=nil{
			fmt.Println("Invalid member ID. please enter a  vaild number")
			return
		}
	err=lib.BorrowBook(bookID,memberID)
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("✅ Book borrowed successfully!")
	}

}
func returnBook(lib *services.Library){
	fmt.Print("Enter a book ID you want to return:")
	id , _:=reader.ReadString('\n')
	id = strings.TrimSpace(id)
	bookID, err:=strconv.Atoi(id)
	if err!=nil{
			fmt.Println("Invalid book ID. please enter a vaild number")
			return
		}
	fmt.Print("Enter member ID:")
	memberIDStr,_:=reader.ReadString('\n')
	memberIDStr =strings.TrimSpace(memberIDStr)
	memberID ,err:=strconv.Atoi(memberIDStr)
	if err!=nil{
			fmt.Println(("Invaild member ID. please enter a vaild number"))
			return
		}
	err =lib.ReturnBook(bookID,memberID)
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println("✅ Book returned successfully!")
	}
}
func listAvailableBooks(lib *services.Library){
	books:=lib.ListAvailableBooks()
	if len(books)==0{
		fmt.Println("No book available")
	}else{
		for _, book:=range books{
			fmt.Printf("ID:%d | Title:%s | Author:%s\n",book.ID, book.Title, book.Author)
		}
	}

}
func listBorrowedBooks(lib *services.Library){
	fmt.Print("Enter member ID to view borrowed books:")
	memberIDStr,_:=reader.ReadString('\n')
	memberIDStr =strings.TrimSpace(memberIDStr)
	memberID ,err:=strconv.Atoi(memberIDStr)
	if err!=nil{
			fmt.Println("Invalid member ID. please enter a  vaild number")
			return
		}
	books:=lib.ListBorrowedBooks(memberID)
	if len(books)==0{
		fmt.Println("No book borrowed")
	}else{
		for _, book:=range books{
			fmt.Printf("ID:%d | Title:%s | Author:%s\n",book.ID, book.Title, book.Author)
		}
	}

}
