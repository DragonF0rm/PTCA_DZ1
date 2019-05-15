package main

import (
	"fmt"
)

type State struct {
	Message string
	Next int
	Prev int
}

func (state *State)Print() {
	fmt.Println(state.Message)
}

var states = []State{
	{Message: "Страница1",
		Next: 1,
		Prev: -1},
	{Message: "Страница2",
		Next: 2,
		Prev: 0},
	{Message: "Страница3",
		Next: 3,
		Prev: 1},
	{Message: "Страница4",
		Next: 4,
		Prev: 2},
	{Message: "Страница5",
		Next: -1,
		Prev: 3},
}

var currentState *State

func main() {
	currentState = &states[0]
	currentState.Print()
	var input string

	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}

		switch input {
		case "1":
			if currentState != &states[0] {
				currentState = &states[0]
				currentState.Print()
			}
		case "2":
			if currentState != &states[1] {
				currentState = &states[1]
				currentState.Print()
			}
		case "3":
			if currentState != &states[2] {
				currentState = &states[2]
				currentState.Print()
			}
		case "4":
			if currentState != &states[3] {
				currentState = &states[3]
				currentState.Print()
			}
		case "5":
			if currentState != &states[4] {
				currentState = &states[4]
				currentState.Print()
			}
		case "next":
			if currentState.Next != -1 {
				currentState = &states[currentState.Next]
				currentState.Print()
			}
		case "prev":
			if currentState.Prev != -1 {
				currentState = &states[currentState.Prev]
				currentState.Print()
			}
		}
	}
}
