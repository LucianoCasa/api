package order

type Order struct {
	ID           int     `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerName string  `json:"customer_name"`
	Total        float64 `json:"total"`
	CreatedAt    string  `json:"created_at"`
}
