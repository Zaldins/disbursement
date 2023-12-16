package dto

type (
	DisbursementRequest struct {
		WalletID uint    `json:"wallet_id"`
		Amount   float64 `json:"amount"`
	}
)