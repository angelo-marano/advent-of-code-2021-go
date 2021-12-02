package main

import (
	"fmt"

	"./day_one"
	"./day_two"
)

func main() {
	fmt.Println("--- Day One Solutions ---")

	day_one_answers, err := day_one.DayOne()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("There are %d measurements greater than the previous measurement. \n", day_one_answers.Ascending_items_count)
	fmt.Printf("There are %d windows larger than the previous window, \n", day_one_answers.Ascending_slices_count)

	fmt.Println("--- Day Two Solutions ---")
	day_two_answers, err := day_two.DayTwo()
	if err != nil {
		fmt.Print(err)
	}

	product := day_two_answers.Position.X * day_two_answers.Position.Y

	fmt.Printf("The sub ends up at horizional position %d and depth %d, with product of %d. \n", day_two_answers.Position.X, day_two_answers.Position.Y, product)

	product = day_two_answers.Position_With_Aim.X * day_two_answers.Position_With_Aim.Y

	fmt.Printf("Applying aim, The sub ends up at horizional position %d and depth %d, with product of %d. \n", day_two_answers.Position_With_Aim.X, day_two_answers.Position_With_Aim.Y, product)
}
