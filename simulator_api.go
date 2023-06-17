package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mazeHeight() int {
	fmt.Println("mazeHeight")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	conv_response, err := strconv.Atoi(response)
	if err != nil {
		panic(err)
	}
	return conv_response
}

func mazeWidth() int {
	fmt.Println("mazeWidth")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	conv_response, err := strconv.Atoi(response)
	if err != nil {
		panic(err)
	}
	return conv_response
}

func moveForward() (string, error) {
	fmt.Println("moveForward")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	if response == "crash" {
		return response, fmt.Errorf("Car crashed")
	}
	return response, nil
}

func turnRight() string {
	fmt.Println("turnRight")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	return response
}

func turnLeft() string {
	fmt.Println("turnLeft")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	return response
}

func wallFront() bool {
	fmt.Println("wallFront")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	return response == "true"
}

func wallLeft() bool {
	fmt.Println("wallLeft")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	return response == "true"
}

func wallRight() bool {
	fmt.Println("wallRight")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	return response == "true"
}

func setWall(x int, y int, direction string) {
	fmt.Printf("setWall %d %d %s\n", x, y, direction)
}

func wasReset() bool {
	fmt.Println("wasReset")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	return response == "true"
}

func ackReset() string {
	fmt.Println("ackReset")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	return response
}

func setColor(x int, y int, color string) {
	fmt.Printf("setColor %d %d %s\n", x, y, color)
}

func setText(x int, y int, text string) {
	fmt.Printf("setText %d %d %s\n", x, y, text)
}
