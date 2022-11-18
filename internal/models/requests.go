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
