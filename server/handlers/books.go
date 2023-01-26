package handlers

import (
	"context"
	"crypto/tls"
	"encoding/json"
	bookdto "final-task/dto/book"
	dto "final-task/dto/result"
	"final-task/models"
	"final-task/repositories"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"

	// package for midtrans

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"

	// package for cloudinary
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type Bookhandler struct {
	BookRepository repositories.BookRepository
}

func HandlerBook(BookRepository repositories.BookRepository) *Bookhandler {
	return &Bookhandler{BookRepository}
}

var c = coreapi.Client{
	ServerKey: os.Getenv("SERVER_KEY"),
	ClientKey: os.Getenv("CLIENT_KEY"),
}

func (h *Bookhandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book, err := h.BookRepository.FindBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	if err != nil {
		fmt.Println("errornya adalah", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: book}
	json.NewEncoder(w).Encode(response)
}

func (h *Bookhandler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// user, err := h.UserRepository.GetUser(id)

	book, err := h.BookRepository.FindBook(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK,
		Data: book,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Bookhandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContext := r.Context().Value("dataFile")

	// block code for saving image to cloudinary
	// assign nama file ke variable filename
	filepath := dataContext.(string)

	// NOTE: face error caused by key value in postman using whitespace after it , DONT DO THAT !!

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dewetour"})

	if err != nil {
		fmt.Println("error when uploading images", err.Error())
	}

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// This block of code create a new transaction id with unique number and assing it to model to query to db
	var request bookdto.CreateTransactionRequest
	var TransIdIsMatch = false
	var TransactionId int

	for !TransIdIsMatch {
		TransactionId = int(time.Now().Unix())
		transaction, _ := h.BookRepository.FindBook(TransactionId)
		if transaction.ID == 0 {
			TransIdIsMatch = true
		}
	}
	// filename := dataContext.(string)
	dataPrice, _ := strconv.Atoi(r.FormValue("price"))
	dataFilmId, _ := strconv.Atoi(r.FormValue("film_id"))

	request = bookdto.CreateTransactionRequest{
		ID:            TransactionId,
		Price:         dataPrice,
		Status:        "pending",
		TransferProof: resp.SecureURL,
		FilmId:        dataFilmId,
		// AccountNumber: r.FormValue(""),
		OrderDate: r.FormValue("order_date"),
		UserId:    userId,
	}

	// validate request against struct form created
	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book := models.Book{
		ID:            TransactionId,
		Price:         dataPrice,
		Status:        request.Status,
		TransferProof: request.TransferProof,
		FilmId:        request.FilmId,
		AccountNumber: request.AccountNumber,
		OrderDate:     request.OrderDate,
		UserId:        userId,
	}

	newTransaction, err := h.BookRepository.CreateBook(book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

	}
	// dataTransactions, err := h.BookRepository.FindBook(int(newTransaction.ID))
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(err.Error())
	// 	return
	// }

	newData, err := h.BookRepository.FindBook(newTransaction.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{OrderID: strconv.Itoa(int(newTransaction.ID)), GrossAmt: int64(newTransaction.Price)},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: newData.User.Full_name,
			Email: newData.User.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)

}

func SendEmail(status string, transaksi models.Book) {

	var CONFIG_SMTP_HOST = "smtp.gmail.com"
	var CONFIG_SMTP_PORT = 587
	var CONFIG_SENDER_NAME = "Net <arifthalhah@gmail.com>"
	var CONFIG_AUTH_EMAIL = os.Getenv("SYSTEM_EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("SYSTEM_PASSWORD")

	var BookName = transaksi.User.Full_name
	var price = strconv.Itoa(transaksi.Film.Price)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", transaksi.User.Email)
	mailer.SetHeader("Subject", "Status Transaksi")
	mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
    <html lang="en">
      <head>
      <meta charset="UTF-8" />
      <meta http-equiv="X-UA-Compatible" content="IE=edge" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Document</title>
      <style>
        h1 {
        color: brown;
        }
      </style>
      </head>
      <body>
      <h2>Product payment :</h2>
      <ul style="list-style-type:none;">
        <li>Name : %s</li>
        <li>Total Harga: Rp.%s</li>
        <li>Status : <b>%s</b></li>
        <li>Iklan : <b>%s</b></li>
      </ul>
      </body>
    </html>`, BookName, price, status, "Gausah Beli Lagi"))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *Bookhandler) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// pp.Println(notificationPayload)

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)
	parsed, _ := strconv.Atoi(orderId)

	transaction, _ := h.BookRepository.FindBook(parsed)

	// pp.Println(transaction)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal

			h.BookRepository.UpdateBook("pending", transaction.ID)
			// SendEmail("pending", transaction)
		} else if fraudStatus == "accept" {

			h.BookRepository.UpdateBook("success", transaction.ID)
			// SendEmail("success", transaction)
		}
	} else if transactionStatus == "settlement" {

		h.BookRepository.UpdateBook("success", transaction.ID)
		SendEmail("success", transaction)
	} else if transactionStatus == "deny" {
		// and later can become success
		h.BookRepository.UpdateBook("failed", transaction.ID)

	} else if transactionStatus == "cancel" || transactionStatus == "expire" {

		h.BookRepository.UpdateBook("failed", transaction.ID)
		SendEmail("failed", transaction)
	} else if transactionStatus == "pending" {

		h.BookRepository.UpdateBook("pending", transaction.ID)
		// SendEmail("pending", transaction)
	}

	w.WriteHeader(http.StatusOK)

}

// func (h *Transactionhandler) UpdateTrans(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	dataContext := r.Context().Value("dataFile")

// 	// assign nama file ke variable filename
// 	filename := dataContext.(string)

// 	// NOTE: face error caused by key value in postman using whitespace after it , DONT DO THAT !!

// 	// get data country convrt ke int
// 	dataQty, _ := strconv.Atoi(r.FormValue("counter_qty"))
// 	dataTotal, _ := strconv.Atoi(r.FormValue("total"))
// 	dataTrip, _ := strconv.Atoi(r.FormValue("trip_id"))
// 	// dataPrice, _ := strconv.Atoi(r.FormValue("price"))
// 	// dataQuota, _ := strconv.Atoi(r.FormValue("quota"))

// 	request := transactiondto.CreateTransactionRequest{
// 		CounterQty: dataQty,
// 		Total:      dataTotal,
// 		Status:     r.FormValue("status"),
// 		Attachment: filename,
// 		TripId:     dataTrip,
// 	}

// 	// fmt.Println(request)
// 	// return

// 	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 	// 	w.WriteHeader(http.StatusBadRequest)
// 	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 	// 	json.NewEncoder(w).Encode(response)
// 	// 	return
// 	// }

// 	// validate request against struct form created

// 	// fmt.Println("baris ke 203 trip handler ", request)
// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// countryId := strconv.Atoi()

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	// fmt.Println(id)
// 	// return

// 	// trip, _ = h.TripRepository.FindSingleTrip(id)

// 	// fmt.Println(request.Country)
// 	// return

// 	trans := models.Transaction{}

// 	// check all field for emptieness

// 	if request.Attachment != "" {
// 		trans.Attachment = request.Attachment
// 	}
// 	if request.Status != "" {
// 		trans.Status = request.Status
// 	}

// 	if request.CounterQty != 0 {
// 		trans.CounterQty = request.CounterQty
// 	}

// 	if request.Total != 0 {
// 		trans.Total = request.Total
// 	}

// 	if request.TripId != 0 {
// 		trans.TripId = request.TripId
// 	}

// 	// fmt.Println(request)
// 	// return

// 	data, err := h.TransactionRepository.UpdateTrans()

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err.Error())

// 	}
// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
// 	json.NewEncoder(w).Encode(response)

// }
