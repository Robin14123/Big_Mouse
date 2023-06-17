package main

import (
	"fmt"
)

type WallMap struct {
	width      int
	height     int
	horizontal [][]bool
	vertical   [][]bool
}

func InitiateWallMap(width int, height int) *WallMap {
	// Uses the real size of the maze, so add +1 to map the walls
	// The map represent the edges with True if a wall exists
	horizontal_map := make([][]bool, height+1)
	vertical_map := make([][]bool, height)

	horizontal_map[height] = make([]bool, width)
	r := 0
	for r < height {
		horizontal_map[r] = make([]bool, width)
		vertical_map[r] = make([]bool, width+1)
		c := 0
		for c < width {

			// We fill the first & last row of horizontal map with "true"
			if r == 0 {
				horizontal_map[r][c] = true
			} else if r == height-1 {
				horizontal_map[r+1][c] = true
			} else {
				horizontal_map[r][c] = false
			}

			// We fill every first & last element of each row (eg first & last column) with "true"
			if c == 0 {
				vertical_map[r][c] = true
			} else if c == width-1 {
				vertical_map[r][c+1] = true
			} else {
				vertical_map[r][c] = false
			}
			c += 1
		}
		r += 1
	}
	return &WallMap{
		height:     height,
		width:      width,
		horizontal: horizontal_map,
		vertical:   vertical_map,
	}
}

func Show2DArray(array [][]bool) {
	i := 0
	for i < len(array) {
		fmt.Println(array[i])
		i += 1
	}
}

func basic_exploration(mouse_state *MouseState) {
	maze_width, maze_height := mazeWidth(), mazeHeight()
	flood_strat := NewFloodStrategy(maze_width, maze_height)
	previous_cell := Coordinates{x: 0, y: 0}
	current_cell := Coordinates{x: 0, y: 0}
	for !wasReset() {
		detect_walls(mouse_state, flood_strat.maze_walls)
		move_to_next_cell(mouse_state, previous_cell, flood_strat)
		previous_cell = current_cell
		current_cell = mouse_state.cell
	}
	ackReset()
}

func detect_walls(mouse_state *MouseState, walls_map *WallMap) []Direction {
	detected_walls_direction := []Direction{}
	x, y := mouse_state.cell.x, mouse_state.cell.y
	if wallLeft() {
		direction := mouse_state.direction.GetCurrent(-1)
		setWall(x, y, string(direction))
		detected_walls_direction = append(detected_walls_direction, direction)
	}
	if wallRight() {
		direction := mouse_state.direction.GetCurrent(1)
		setWall(x, y, string(direction))
		detected_walls_direction = append(detected_walls_direction, direction)
	}
	if wallFront() {
		direction := mouse_state.direction.GetCurrent(0)
		setWall(x, y, string(direction))
		detected_walls_direction = append(detected_walls_direction, direction)
	}
	update_wall_mapping(x, y, detected_walls_direction, walls_map)
	return detected_walls_direction
}

func update_wall_mapping(mouse_x int, mouse_y int, directions []Direction, walls_map *WallMap) {
	for _, direction := range directions {
		switch direction {
		case North:
			walls_map.horizontal[mouse_y+1][mouse_x] = true
		case South:
			walls_map.horizontal[mouse_y][mouse_x] = true
		case East:
			walls_map.vertical[mouse_y][mouse_x+1] = true
		case West:
			walls_map.vertical[mouse_y][mouse_x] = true
		}
	}
}

func move_to_next_cell(mouse_state *MouseState, previous_cell Coordinates, strat Strategy) {
	// Get current cell value and next direction
	_, next_cell := strat.process_one_step(mouse_state, previous_cell)

	orient_mouse_towards_next_cell(mouse_state, next_cell)
	// We need to be oriented according to the shortest path

	moveForward()
	mouse_state.cell = next_cell
}

func orient_mouse_towards_next_cell(mouse_state *MouseState, next_cell Coordinates) {
	var next_direction Direction
	if mouse_state.cell.x < next_cell.x {
		next_direction = East
	} else if mouse_state.cell.x > next_cell.x {
		next_direction = West
	} else if mouse_state.cell.y < next_cell.y {
		next_direction = North
	} else if mouse_state.cell.y > next_cell.y {
		next_direction = South
	}

	// create methods for MouseState to turn & update the mouse direction at the same time
	switch next_direction {
	case mouse_state.direction.GetCurrent(1):
		turnRight()
		mouse_state.direction.AntiTrigoRotation()
	case mouse_state.direction.GetCurrent(-1):
		turnLeft()
		mouse_state.direction.TrigoRotation()
	case mouse_state.direction.GetCurrent(-2):
		turnLeft()
		turnLeft()
		mouse_state.direction.TrigoRotation()
		mouse_state.direction.TrigoRotation()
	}

}
