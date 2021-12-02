package day_one

import (
	"os"
	"strconv"

	"../advent_utils"
)

type DayOneAnswers struct {
	Ascending_items_count  int
	Ascending_slices_count int
}

func DayOne() (DayOneAnswers, error) {
	file, err := os.Open("./input/day_one.txt")
	var a DayOneAnswers
	if err != nil {
		return a, err
	}

	defer file.Close()
	lines, err := advent_utils.ReadAllLines(file)
	if err != nil {
		return a, err
	}

	nums, err := convert_to_int_array(lines)

	if err != nil {
		return a, err
	}

	ascending_items, err := count_ascending_items(nums)
	if err != nil {
		return a, err
	}
	ascending_slices, err := count_ascending_slices(nums, 3)
	if err != nil {
		return a, err
	}

	a.Ascending_items_count = ascending_items
	a.Ascending_slices_count = ascending_slices

	return a, nil

}

func convert_to_int_array(lines []string) ([]int, error) {
	var nums []int
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nums, err
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func count_ascending_items(lines []int) (int, error) {
	count := 0
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			count++
		}
	}

	return count, nil
}

func count_ascending_slices(lines []int, slice_size int) (int, error) {
	count := 0
	previous_sum := sum(lines[0:slice_size])
	for i := 1; i <= len(lines)-slice_size; i++ {
		new_sum := sum(lines[i : i+slice_size])
		if new_sum > previous_sum {
			count++
		}
		previous_sum = new_sum
	}

	return count, nil
}

func sum(s []int) int {
	sum := 0
	for _, num := range s {
		sum += num
	}

	return sum

}
