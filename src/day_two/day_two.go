package day_two

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"../advent_utils"
)

type Position struct {
	X int
	Y int
}

type DayTwoAnswers struct {
	Position          Position
	Position_With_Aim Position
}

type Direction string

const (
	Forward Direction = "forward"
	Down    Direction = "down"
	Up      Direction = "up"
)

type Instruction struct {
	Direction Direction
	count     int
}

func DayTwo() (DayTwoAnswers, error) {

	file, err := os.Open("./input/day_two.txt")
	var position Position
	var position_with_aim Position
	var answers DayTwoAnswers
	if err != nil {
		return answers, err
	}

	defer file.Close()

	lines, err := advent_utils.ReadAllLines(file)
	if err != nil {
		return answers, err
	}

	aim := 0
	for _, line := range lines {
		instruction, err := parse_instruction(line)
		if err != nil {
			return answers, err
		}
		apply_instruction(instruction, &position)
		apply_instruction_with_aim(instruction, &position_with_aim, &aim)
	}

	answers.Position = position
	answers.Position_With_Aim = position_with_aim

	return answers, nil
}

func apply_instruction(instruction Instruction, p *Position) {
	switch instruction.Direction {
	case Down:
		p.Y += instruction.count
	case Up:
		p.Y -= instruction.count
	case Forward:
		p.X += instruction.count
	}
}

func apply_instruction_with_aim(instruction Instruction, p *Position, aim *int) {
	switch instruction.Direction {
	case Down:
		*aim += instruction.count
	case Up:
		*aim -= instruction.count
	case Forward:
		p.X += instruction.count
		p.Y += instruction.count * *aim
	}
}

func parse_instruction(str string) (Instruction, error) {
	split := strings.Split(str, " ")
	direction := split[0]

	var d Direction
	var instruction Instruction

	switch direction {
	case "forward":
		d = Forward
	case "up":
		d = Up
	case "down":
		d = Down
	default:
		return instruction, errors.New("direction must be one of 'forward', 'up', 'down'")
	}

	count, err := strconv.Atoi(split[1])
	if err != nil {
		return instruction, err
	}

	instruction.count = count
	instruction.Direction = d
	return instruction, nil

}
