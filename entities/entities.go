package entities

import "time"

type Book struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	ISBN        string `gorm:"uniqueIndex" json:"isbn"`
	Year        int    `json:"year"`
	CoverURL    string `json:"cover_url"`
	Status      string `gorm:"default:'available'" json:"status"` // "available", "borrowed"

	AuthorID   uint `json:"author_id"`
	CategoryID uint `json:"category_id"`

	Author   Author   `json:"author"`
	Category Category `json:"category"`
	Reviews  []Review `json:"reviews"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `json:"-"`                          // не отдаём в JSON
	Role     string `gorm:"default:'user'" json:"role"` // "user", "admin", "librarian"

	Loans   []Loan   `json:"loans"`
	Reviews []Review `json:"reviews"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Category struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"uniqueIndex;not null" json:"name"`
	Books []Book `json:"books"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Author struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Bio      string `json:"bio"`
	PhotoURL string `json:"photo_url"`
	Books    []Book `json:"books"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Loan struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	UserID     uint       `json:"user_id"`
	BookID     uint       `json:"book_id"`
	BorrowedAt time.Time  `json:"borrowed_at"`
	DueDate    time.Time  `json:"due_date"`
	ReturnedAt *time.Time `json:"returned_at,omitempty"`

	User User `json:"user"`
	Book Book `json:"book"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Review struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	UserID  uint   `json:"user_id"`
	BookID  uint   `json:"book_id"`
	Rating  int    `gorm:"check:rating >= 1 AND rating <= 5" json:"rating"`
	Comment string `json:"comment"`

	User User `json:"user"`
	Book Book `json:"book"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
