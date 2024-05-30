package data

type RecipeIngredientSection struct {
	ID          int64
	SectionName string
	Ingredients []RecipeIngredient
}

type RecipeIngredient struct {
	ID            int64
	FoodID        int64
	Quantity      float64
	MeasurementID int64
}
