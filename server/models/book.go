package models

type Book struct {
	// gorm.Model
	ID            int          `json:"book_id"`
	Price         int          `json:"price"`
	Status        string       `json:"status"`
	TransferProof string       `json:"transfer_proof"`
	FilmId        int          `json:"film_id"`
	Film          FilmResponse `json:"film"`
	AccountNumber int          `json:"account_number"`
	OrderDate     string       `json:"order_date"`
	UserId        int          `json:"user_id"`
	User          UserResponse `json:"user"`
}

type BookResponse struct {
	// gorm.Model
	ID            int          `json:"book_id"`
	Price         int          `json:"price"`
	Status        string       `json:"status"`
	TransferProof string       `json:"transfer_proof"`
	FilmId        int          `json:"film_id"`
	Film          FilmResponse `json:"film"`
	AccountNumber int          `json:"account_number"`
	OrderDate     string       `json:"order_date"`
	UserId        int          `json:"user_id"`
	User          UserResponse `json:"user"`
}

func (BookResponse) TableName() string {
	return "books"
}
