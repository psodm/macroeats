package data

type RecipeInstruction struct {
	Step        int64  `json:"instructionStep"`
	Description string `json:"instructionDescription"`
}
