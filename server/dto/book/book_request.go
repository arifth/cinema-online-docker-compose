package dto

type CreateTransactionRequest struct {
	Price         int    `json:"price"`
	Status        string `json:"status"`
	TransferProof string `json:"transfer_proof"`
	FilmId        int    `json:"film_id"`
	AccountNumber int    `json:"account_number"`
	OrderDate     string `json:"order_date"`
	UserId        int    `json:"user_id"`
}
