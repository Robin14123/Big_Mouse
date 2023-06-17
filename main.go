package main

func main() {
	setColor(0, 0, "G")
	mouse_state := InitialState()
	basic_exploration(&mouse_state)
}
