package data

type RecipeIngredientSection struct {
	SectionName string             `json:"section"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

type RecipeIngredient struct {
	ID                      int64   `json:"Id"`
	IngredientName          string  `json:"ingredientName"`
	MeasurementQuantity     float64 `json:"ingredientAmount"`
	MeasurementAbbreviation string  `json:"measurementAbbreviation"`
}
