package main

import (
	"go-basics/variables"
)

func main() {
	// godotenv.Load()
	// fmt.Println(os.Getenv("GO_ENV"))
	// fmt.Println(calculator.Offset)
	// fmt.Println(calculator.Sum(1, 2))
	// fmt.Println("Multiply: ",calculator.Multiply(1, 2))
	variables.Init()
}