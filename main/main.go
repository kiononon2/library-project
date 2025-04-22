package main

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"libraryProject/entities"
	"log"
	"net/http"
	"strconv"
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
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			var books []entities.Book
			db.Preload("Author").Preload("Category").Preload("Reviews").Find(&books)
			json.NewEncoder(w).Encode(books)
		} else if idStr != "" {
			id, _ := strconv.Atoi(idStr)
			var book entities.Book
			if err := db.Preload("Author").Preload("Category").Preload("Reviews").First(&book, id).Error; err != nil {
				respondWithError(w, http.StatusNotFound, "Book not found")
				return
			}
			json.NewEncoder(w).Encode(book)
		}
	}
	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		if err := db.Delete(&entities.Book{}, id).Error; err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete book")
			return
		}
		w.WriteHeader(http.StatusNoContent)
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
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			var users []entities.User
			db.Preload("Loans").Preload("Reviews").Find(&users)
			json.NewEncoder(w).Encode(users)
		} else if idStr != "" {
			id, _ := strconv.Atoi(idStr)
			var user entities.User
			if err := db.Preload("Loans").Preload("Reviews").First(&user, id).Error; err != nil {
				respondWithError(w, http.StatusNotFound, "User not found")
				return
			}
			json.NewEncoder(w).Encode(user)
		}
	}
	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		if err := db.Delete(&entities.User{}, id).Error; err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete user")
			return
		}
		w.WriteHeader(http.StatusNoContent)
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
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			var authors []entities.Author
			db.Preload("Books").Find(&authors)
			json.NewEncoder(w).Encode(authors)
		} else if idStr != "" {
			id, _ := strconv.Atoi(idStr)
			var author entities.Author
			if err := db.Preload("Books").First(&author, id).Error; err != nil {
				respondWithError(w, http.StatusNotFound, "Author not found")
				return
			}
			json.NewEncoder(w).Encode(author)
		}
	}
	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		var count int64
		db.Model(&entities.Book{}).Where("author_id = ?", id).Count(&count)
		if count > 0 {
			respondWithError(w, http.StatusConflict, "Cannot delete author with linked books")
			return
		}

		if err := db.Delete(&entities.Author{}, id).Error; err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete author")
			return
		}
		w.WriteHeader(http.StatusNoContent)
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
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			var categories []entities.Category
			db.Preload("Books").Find(&categories)
			json.NewEncoder(w).Encode(categories)
		} else if idStr != "" {
			id, _ := strconv.Atoi(idStr)
			var categories []entities.Category
			if err := db.Preload("Books").First(&categories, id).Error; err != nil {
				respondWithError(w, http.StatusNotFound, "Category not found")
				return
			}
			json.NewEncoder(w).Encode(categories)
		}
	}
	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		if err := db.Delete(&entities.Category{}, id).Error; err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete category")
			return
		}
		w.WriteHeader(http.StatusNoContent)
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
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			var loans []entities.Loan
			db.Preload("User").Preload("Book").Find(&loans)
			json.NewEncoder(w).Encode(loans)
		} else if idStr != "" {
			id, _ := strconv.Atoi(idStr)
			var loan entities.Loan
			if err := db.Preload("User").Preload("Book").First(&loan, id).Error; err != nil {
				respondWithError(w, http.StatusNotFound, "Loan not found")
				return
			}
			json.NewEncoder(w).Encode(loan)
		}
	}
	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		if err := db.Delete(&entities.Loan{}, id).Error; err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete loan")
			return
		}
		w.WriteHeader(http.StatusNoContent)
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
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			var reviews []entities.Review
			db.Preload("User").Preload("Book").Find(&reviews)
			json.NewEncoder(w).Encode(reviews)
		} else if idStr != "" {
			id, _ := strconv.Atoi(idStr)
			var review entities.Review
			if err := db.Preload("User").Preload("Book").First(&review, id).Error; err != nil {
				respondWithError(w, http.StatusNotFound, "Review not found")
				return
			}
			json.NewEncoder(w).Encode(review)
		}
	}
	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		if err := db.Delete(&entities.Review{}, id).Error; err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete review")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
