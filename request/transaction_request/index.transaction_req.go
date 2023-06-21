package transaction_request

type TranferReq struct {
	RecipientAccountNumber string `json:"recipient_account_number" binding:"required"`
	Pin                    int `json:"pin" binding:"required"`
	Amount                 int `json:"amount" binding:"required"`
}
