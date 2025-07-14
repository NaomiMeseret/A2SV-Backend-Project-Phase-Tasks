package services

import (
	"errors"
	"library_management/models"
)
type LibraryManager interface{
	AddBook(book models.Book)
	RemoveBook(bookID int)error
	BorrowBook(bookID,memberID int)error
	ReturnBook(bookID, memberID int)error
	ListAvailableBooks()[]models.Book
	ListBorrowedBooks(memberID int)[]models.Book
}
type Library struct{
	books map[int]models.Book
	members map[int]models.Member
	nextID int
}
func NewLibrary()*Library{
	return&Library{
		books : make(map[int]models.Book),
		members : map[int]models.Member{
			001:{ID:101 , Name:"Naomi"},
		},
		nextID : 1,
	}
}
func(l *Library) AddBook(book models.Book){
	book.ID = l.nextID
	book.Status = "Available"
	l.books[book.ID] =book
	l.nextID++
}
func (l *Library) RemoveBook(bookID int)error{
	if _ , ok:=l.books[bookID];!ok{
		return errors.New("book not found")
	}
	delete(l.books , bookID)
	return nil

}
func (l *Library) BorrowBook(bookID int , memberID int)error{
	book , ok := l.books[bookID]
	if !ok{
		return errors.New("book not found")
	}
	if book.Status == "Borrowed"{
		return errors.New("book already Borrowed")
	}
	book.Status = "Borrowed"
	l.books[bookID] = book
	member, ok := l.members[memberID]
	if !ok{
		return errors.New("member not found")
	}
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member
	return nil
}
func (l *Library) ReturnBook(bookID int , memberID int)error{
	if book , ok:=l.books[bookID];ok{
		book.Status = "Available"
		l.books[bookID] = book
	}else{
		return errors.New("book not found")
	}
	if member , ok:=l.members[memberID];ok{
		newList := []models.Book{}
		for _ , b:=range member.BorrowedBooks{
			if b.ID != bookID{
				newList = append(newList , b)
			}
		}
		member.BorrowedBooks = newList
		l.members[memberID] = member
		return nil

	}else{
		return errors.New("member not found")
	}
}
func (l *Library) ListAvailableBooks()[]models.Book{
	var availableBooks[]models.Book
	for _ , b := range l.books{
		if b.Status == "Available"{
			availableBooks = append(availableBooks , b)
		}
	}
	return availableBooks
}
func (l *Library) ListBorrowedBooks(memberID int)[]models.Book{
	member , ok :=l.members[memberID]
	if !ok{
		return nil
	}
	return member.BorrowedBooks

}