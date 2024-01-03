package main

import "fmt"

type Movement struct {
	x, y, z int
}

func (m *Movement) moveUp() {
	m.y++
}

func (m *Movement) moveRight() {
	m.x++
}

func (m *Movement) moveDown() {
	m.y--
}

func (m *Movement) moveLeft() {
	m.x--
}

func (m *Movement) moveForward() {
	m.z++
}

func (m *Movement) moveBackward() {
	m.z--
}

func (m *Movement) showCoordinates() {
	fmt.Printf("X: %d, Y: %d, Z: %d\n", m.x, m.y, m.z)
}
