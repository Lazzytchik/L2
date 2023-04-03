package pattern

import (
	"fmt"
	"time"
)

/*
	Паттерн Состояние (State)
	Поведенческий паттерн позволяющий менять поведение объекта в зависимости от его текущего состояния.
*/

// Context represents the traffic light object that can be in different states
type TrafficLight struct {
	state State
}

// State is the interface that all states must implement
type State interface {
	Handle(context *TrafficLight)
}

// RedState represents the traffic light in the red state
type RedState struct{}

func (s *RedState) Handle(context *TrafficLight) {
	fmt.Println("Traffic light is red.")
	time.Sleep(5 * time.Second)   // Simulate traffic light delay
	context.state = &GreenState{} // Transition to the next state
}

// GreenState represents the traffic light in the green state
type GreenState struct{}

func (s *GreenState) Handle(context *TrafficLight) {
	fmt.Println("Traffic light is green.")
	time.Sleep(5 * time.Second)    // Simulate traffic light delay
	context.state = &YellowState{} // Transition to the next state
}

// YellowState represents the traffic light in the yellow state
type YellowState struct{}

func (s *YellowState) Handle(context *TrafficLight) {
	fmt.Println("Traffic light is yellow.")
	time.Sleep(2 * time.Second) // Simulate traffic light delay
	context.state = &RedState{} // Transition to the next state
}

func main() {
	trafficLight := &TrafficLight{state: &RedState{}}

	// Loop infinitely, transitioning between traffic light states
	for {
		trafficLight.state.Handle(trafficLight)
	}
}
