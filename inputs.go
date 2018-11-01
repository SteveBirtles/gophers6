package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	myX     float64 = -50
	myY     float64 = 10
	myZ     float64 = 0
	pitch   float64 = 0
	bearing float64 = 0
)

func processInputs() {

	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}

	// <-- add code here

}
