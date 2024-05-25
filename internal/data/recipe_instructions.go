package data

type RecipeInstruction struct {
	Step        int64  `json:"step"`
	Description string `json:"instructionDescription"`
}
