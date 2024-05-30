package data

type RecipeIngredientSection struct {
	ID          int64
	SectionName string
}

type RecipeIngredient struct {
	ID            int64
	FoodID        int64
	SectionID     int64
	Quantity      float64
	MeasurementID int64
}
