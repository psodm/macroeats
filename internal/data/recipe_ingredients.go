package data

type RecipeIngredientSection struct {
	SectionName string             `json:"section"`
	Ingredients []RecipeIngredient `json:"ingredient"`
}

type RecipeIngredient struct {
	IngredientName         string  `json:"ingredientName"`
	MeasurementQuantity    float64 `json:"ingredientAmount"`
	MeasurementDescription string  `json:"measurementDescription"`
}
