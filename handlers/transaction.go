package handlers

import (
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var c = coreapi.Client{
	ServerKey: os.Getenv("SERVER_KEY"),
	ClientKey: os.Getenv("CLIENT_KEY"),
}
type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository

}
func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction{
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) GetByLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	order, err := h.TransactionRepository.GetByLogin(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: order}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTransaction)FindTransaction(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	transactions,err := h.TransactionRepository.FindTransaction()
	if err != nil{
	w.WriteHeader(http.StatusInternalServerError)
	response := dto.ErrorResult{Code: http.StatusBadRequest,Message :err.Error()}
	json.NewEncoder(w).Encode(response)
	return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code :http.StatusOK,Data:transactions}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTransaction) GetTransactionById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	id,_ := strconv.Atoi(mux.Vars(r)["id"])
	transaction,err := h.TransactionRepository.GetTransactionById(id)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code:http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response:= dto.SuccessResult { Code: http.StatusOK,Data: transaction}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content_type","application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	time := time.Now()
	miliTime := time.Unix()

	transaction := models.Transaction{
		ID: int(miliTime),
		UserID: userId,
		Total: 35000,
		Status: "pending",
		Limit: 0,
		StatusUser: "No Active",
	}
	newTransaction,err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	dataTransactions,err := h.TransactionRepository.GetTransactionById(newTransaction.ID)
	if err != nil {
		 w.WriteHeader(http.StatusInternalServerError)
		 json.NewEncoder(w).Encode(err.Error())
		return
		}
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(dataTransactions.ID),
			GrossAmt: int64(dataTransactions.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataTransactions.User.FullName,
			Email: dataTransactions.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)
}

// Notification method ...
func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.UpdateTransaction("pending", orderId)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			h.TransactionRepository.UpdateTransaction("success", orderId)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		h.TransactionRepository.UpdateTransaction("success", orderId)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.UpdateTransaction("pending", orderId)
	}

	w.WriteHeader(http.StatusOK)
}


	
