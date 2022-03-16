package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("-------------------")
	fmt.Println("BMI Calculator")
	fmt.Println("-------------------")

	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Please enter your hight (m): ")
	heightInput, _ := reader.ReadString('\n')

	weightInput = weightInput[0 : len(weightInput)-2]
	heightInput = heightInput[0 : len(heightInput)-2]

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)
	fmt.Printf(" \nYour BMI: %.2f  \n", bmi)

	if bmi < 18.5 {
		fmt.Print("Your body status: Underweight")
	} else if bmi < 24.9 {
		fmt.Print("Your body status: Healthy Weight")
	} else if bmi < 29.9 {
		fmt.Print("Your body status: Overweight")
	} else {
		fmt.Print("Your body status: Obesity")
	}
	fmt.Println("\n -------------------")
}
