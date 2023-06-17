package main

import (
	"fmt"
	"math"
)

type Coordinates struct {
	x int
	y int
}

func (cell *Coordinates) Equal(other *Coordinates) bool {
	return cell.x == other.x && cell.y == other.y
}

type CellDistance struct {
	cell  Coordinates
	value int
}

type FloodMap struct {
	width    int
	height   int
	array_2d [][]int
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinInt(values ...int) int {
	curr_min := math.MaxInt32
	for _, val := range values {
		if val < curr_min {
			curr_min = val
		}
	}
	return curr_min
}

type Strategy interface {
	process_one_step(mouse_state *MouseState, previous_cell Coordinates) (int, Coordinates)
}

type FloodStrategy struct {
	flood_map  *FloodMap
	maze_walls *WallMap
}

func InitiateFloodMap(width int, height int) *FloodMap {
	// Create a flood map, eg an array on which the goal is at the center (4 cells square)
	// and we compute the distance to this goal for each cell of the map
	goal_x1, goal_y1 := width/2, height/2
	goal_x2, goal_y2 := goal_x1+1, goal_y1+1
	distance_array := make([][]int, height)
	r := 0
	for r < height {
		distance_array[r] = make([]int, width)
		c := 0
		for c < width {
			x_distance := MinInt(AbsInt(goal_x1-c), AbsInt(goal_x2-c))
			y_distance := MinInt(AbsInt(goal_y1-r), AbsInt(goal_y2-r))
			distance_array[r][c] = x_distance + y_distance
			c += 1
		}
		r += 1
	}
	return &FloodMap{
		height:   height,
		width:    width,
		array_2d: distance_array,
	}
}

func NewFloodStrategy(width int, height int) *FloodStrategy {
	return &FloodStrategy{
		flood_map:  InitiateFloodMap(width, height),
		maze_walls: InitiateWallMap(width, height),
	}
}

func (flood_strat *FloodStrategy) process_one_step(mouse_state *MouseState, previous_cell Coordinates) (int, Coordinates) {

	// Define a new 2d array used to know for which neighbor we need to compute the distance
	already_visited := make([][]bool, flood_strat.flood_map.height)
	for y := range already_visited {
		already_visited[y] = make([]bool, flood_strat.flood_map.width)
		for x := range already_visited[y] {
			already_visited[y][x] = false
		}
	}

	// Get current cell value and next direction
	return flood_strat.get_flood_cell_value(mouse_state.cell, previous_cell, already_visited)

}

func (flood_strat *FloodStrategy) get_flood_cell_value(current_cell Coordinates, last_cell Coordinates, already_visited [][]bool) (int, Coordinates) {
	cell_x, cell_y := current_cell.x, current_cell.y
	// First, we set this cell as visited
	already_visited[cell_y][cell_x] = true
	goal_x1, goal_y1 := flood_strat.flood_map.width/2, flood_strat.flood_map.height/2
	goal_x2, goal_y2 := goal_x1-1, goal_y1-1
	// Final condition: being in the Goal square
	if (cell_x == goal_x1 || cell_x == goal_x2) && (cell_y == goal_y1 || cell_y == goal_y2) {
		flood_strat.flood_map.array_2d[cell_y][cell_x] = 0
		return 0, Coordinates{x: cell_x, y: cell_y}
	}
	// Else, look for accessible neighbors by checking surroundings walls
	accessible_neighbors := []Coordinates{}
	if flood_strat.maze_walls.horizontal[cell_y][cell_x] == false {
		// South
		accessible_neighbors = append(accessible_neighbors, Coordinates{x: cell_x, y: cell_y - 1})
	}
	if flood_strat.maze_walls.horizontal[cell_y+1][cell_x] == false {
		// North
		accessible_neighbors = append(accessible_neighbors, Coordinates{x: cell_x, y: cell_y + 1})
	}
	if flood_strat.maze_walls.vertical[cell_y][cell_x+1] == false {
		// East
		accessible_neighbors = append(accessible_neighbors, Coordinates{x: cell_x + 1, y: cell_y})
	}
	if flood_strat.maze_walls.vertical[cell_y][cell_x] == false {
		// West
		accessible_neighbors = append(accessible_neighbors, Coordinates{x: cell_x - 1, y: cell_y})
	}

	// Finally, find the neighbor with the lowest value (shortest path)
	closest_neighbors := []CellDistance{}
	var next_cell_dist CellDistance
	for _, neighbor := range accessible_neighbors {
		var neighbor_value int

		if !already_visited[neighbor.y][neighbor.x] {
			// If we haven't updated the neighbor yet, we compute its value
			already_visited[neighbor.y][neighbor.x] = true
			neighbor_value, _ = flood_strat.get_flood_cell_value(neighbor, current_cell, already_visited)
		} else {
			// Else don't update neighbors, we just get the value
			neighbor_value = flood_strat.flood_map.array_2d[neighbor.y][neighbor.x]
		}

		// Now, we add the neighbor to the closest collection if needed
		next_cell_dist = CellDistance{
			cell:  neighbor,
			value: neighbor_value,
		}
		if len(closest_neighbors) == 0 {
			// No neighbors visited yet, the first one is a candidate
			closest_neighbors = []CellDistance{next_cell_dist}
		} else if neighbor_value < closest_neighbors[0].value {
			// We empty the collection because this is the only closest now
			closest_neighbors = []CellDistance{next_cell_dist}
		} else if neighbor_value == closest_neighbors[0].value {
			// There are mutliple neighbors with the same value, we will make the decision after
			if next_cell_dist.cell.Equal(&last_cell) {
				closest_neighbors = append(closest_neighbors, next_cell_dist)
			} else {
				closest_neighbors = append([]CellDistance{next_cell_dist}, closest_neighbors...)
			}
		}
	}

	next_cell_dist = closest_neighbors[0]
	setText(cell_x, cell_y, fmt.Sprint(next_cell_dist.value+1))
	flood_strat.flood_map.array_2d[cell_y][cell_x] = next_cell_dist.value + 1
	return next_cell_dist.value + 1, next_cell_dist.cell
}
