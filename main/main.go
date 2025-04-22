package main

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"libraryProject/entities"
	"log"
	"net/http"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(postgres.Open("user=postgres password=postgres dbname=gorillaLibraryProjectDB host=localhost port=5433 sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entities.User{}, &entities.Book{}, &entities.Author{}, &entities.Category{}, &entities.Loan{}, &entities.Review{})
}
func main() {
	initDB()
	defer func() {
		s, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}
		s.Close()
	}()

	//http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
	//	if r.Method == http.MethodPost {
	//		var book entities.Book
	//		json.NewDecoder(r.Body).Decode(&book)
	//		db.Create(&book)
	//		w.WriteHeader(http.StatusCreated)
	//		json.NewEncoder(w).Encode(book)
	//	}
	//
	//	if r.Method == http.MethodGet {
	//		var books []entities.Book
	//		db.Find(&books)
	//		json.NewEncoder(w).Encode(books)
	//	}
	//})
	http.HandleFunc("/book", bookHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/author", authorHandler)
	http.HandleFunc("/category", categoryHandler)
	http.HandleFunc("/loan", loanHandler)
	http.HandleFunc("/review", reviewHandler)

	server := &http.Server{
		Addr: ":3030",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var book entities.Book
		json.NewDecoder(r.Body).Decode(&book)
		db.Create(&book)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
	}
	if r.Method == http.MethodGet {
		var books []entities.Book
		db.Preload("Author").Preload("Category").Preload("Reviews").Find(&books)
		json.NewEncoder(w).Encode(books)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user entities.User
		json.NewDecoder(r.Body).Decode(&user)
		db.Create(&user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
	if r.Method == http.MethodGet {
		var users []entities.User
		db.Preload("Loans").Preload("Reviews").Find(&users)
		json.NewEncoder(w).Encode(users)
	}
}

func authorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var author entities.Author
		json.NewDecoder(r.Body).Decode(&author)
		db.Create(&author)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(author)
	}
	if r.Method == http.MethodGet {
		var authors []entities.Author
		db.Preload("Books").Find(&authors)
		json.NewEncoder(w).Encode(authors)
	}
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var category entities.Category
		json.NewDecoder(r.Body).Decode(&category)
		db.Create(&category)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(category)
	}
	if r.Method == http.MethodGet {
		var categories []entities.Category
		db.Preload("Books").Find(&categories)
		json.NewEncoder(w).Encode(categories)
	}
}

func loanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var loan entities.Loan
		json.NewDecoder(r.Body).Decode(&loan)
		db.Create(&loan)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(loan)
	}
	if r.Method == http.MethodGet {
		var loans []entities.Loan
		db.Preload("User").Preload("Book").Find(&loans)
		json.NewEncoder(w).Encode(loans)
	}
}

func reviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var review entities.Review
		json.NewDecoder(r.Body).Decode(&review)
		db.Create(&review)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(review)
	}
	if r.Method == http.MethodGet {
		var reviews []entities.Review
		db.Preload("User").Preload("Book").Find(&reviews)
		json.NewEncoder(w).Encode(reviews)
	}
}
