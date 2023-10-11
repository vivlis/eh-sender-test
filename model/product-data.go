package model

// ProductData - data structure
type ProductData struct {
	MarketPlace     string   `json:"marketplace"`
	Market          string   `json:"market "`
	EventCreated    string   `json:"eventCreated"`
	ProductID       string   `json:"productId"`
	ValidationCodes []string `json:"validationCodes"`
}
