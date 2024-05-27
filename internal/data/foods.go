package data

type Food struct {
	ID                      int64   `json:"Id"`
	FoodName                string  `json:"foodName"`
	ServingQuantity         float64 `json:"servingSize"`
	ServingUnitAbbreviation string  `json:"servingUnitAbbreviation"`
	Energy                  float64 `json:"energy"`
	Calories                float64 `json:"calories"`
	Protein                 float64 `json:"protein"`
	Carbohydrates           float64 `json:"carbohydrates"`
	Fat                     float64 `json:"fat"`
}
