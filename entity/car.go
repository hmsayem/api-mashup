package entity

type Car struct {
	CarData `json:"Car"`
}

type CarData struct {
	ID    int    `json:"id"`
	Brand string `json:"car"`
	Model string `json:"car_model"`
	Year  string `json:"car_model_year"`
}
