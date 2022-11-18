package models

type AddMessage struct {
	UserID   int     `json:"user_id"`
	Value    float64 `json:"value"`
	Comments string  `json:"comments"`
}

type TransferMessage struct {
	SrcUserID int     `json:"src_user_id"`
	DstUserID int     `json:"dst_user_id"`
	Value     float64 `json:"value"`
	Comments  string  `json:"comments"`
}

type ShortResponseMessage struct {
	Comments string `json:"comments"`
}

type BalanceResponseMessage struct {
	Balance  float64 `json:"balance"`
	Comments string  `json:"comments"`
}

type FinanceReportResponseMessage struct {
	FileURL string `json:"fileURL"`
}

type UserReportResponseMessage struct {
	AllTransactions []Transaction `json:"transactions"`
}

type AcceptServiceMessage struct {
	UserID    int `json:"user_id"`
	OrderID   int `json:"order_id"`
	ServiceID int `json:"service_id"`
}

type ReserveServiceMessage struct {
	UserID    int     `json:"user_id"`
	OrderID   int     `json:"order_id"`
	ServiceID int     `json:"service_id"`
	Value     float64 `json:"value"`
	Comments  string  `json:"comments"`
}

type CancelServiceMessage struct {
	UserID    int    `json:"user_id"`
	OrderID   int    `json:"order_id"`
	ServiceID int    `json:"service_id"`
	Comments  string `json:"comments"`
}
