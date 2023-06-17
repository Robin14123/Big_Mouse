package main

import "math"

const pi = math.Pi

// Mouse related state and functions
func InitialState() MouseState {
	return MouseState{
		cell: Coordinates{
			x: 0,
			y: 0,
		},
		direction: NewDirectionCycle(),
	}
}

type MouseState struct {
	cell      Coordinates
	direction *Cycle
}

func move_mouse(current_state *MouseState, distance int) {
	if distance != 0 {
		switch current_state.direction.GetCurrent(0) {
		case North:
			current_state.cell.y += distance
		case South:
			current_state.cell.y -= distance
		case East:
			current_state.cell.x += distance
		case West:
			current_state.cell.x -= distance
		}
	}
}

func turn_mouse(current_state *MouseState, rotation Rotation) {
	switch rotation {
	case Left:
		current_state.direction.TrigoRotation()
	case Right:
		current_state.direction.AntiTrigoRotation()
	}
}

// Convenient tools to track the mouse current direction
type Direction string

const (
	North Direction = "n"
	South Direction = "s"
	East  Direction = "e"
	West  Direction = "w"
)

type Rotation float64

const (
	Right Rotation = 3 * pi / 2
	Left  Rotation = pi / 2
)

type Cycle struct {
	items []Direction
	index int
}

func NewDirectionCycle() *Cycle {
	return &Cycle{
		items: []Direction{North, East, South, West},
		index: 0,
	}
}

func (cy *Cycle) GetCurrent(offset int) Direction {
	return cy.items[(cy.index+offset+len(cy.items))%len(cy.items)]
}

func (cy *Cycle) AntiTrigoRotation() Direction {
	cy.index = (cy.index + 1 + len(cy.items)) % len(cy.items)
	return cy.GetCurrent(0)
}

func (cy *Cycle) TrigoRotation() Direction {
	cy.index = (cy.index - 1 + len(cy.items)) % len(cy.items)
	return cy.GetCurrent(0)
}
