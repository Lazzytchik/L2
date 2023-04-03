package pattern

import "fmt"

/*
	Паттерн Фабричный метод (Factory method)
	Порождающий паттерн проектирования который предоставляет интерфейс для создания объектов изнутри.

	Сценарии использования:
		- Для динамического создания объектов
		- Для тестирования
  		- Для плагин-архитектуры
*/

type Shape interface {
	Draw() string
}

// Define a Rectangle struct that implements the Shape interface
type Rectangle struct{}

func (r *Rectangle) Draw() string {
	return "Drawing a rectangle"
}

// Define a Circle struct that implements the Shape interface
type Circle struct{}

func (c *Circle) Draw() string {
	return "Drawing a circle"
}

// ShapeFactory interface
type ShapeFactory interface {
	CreateShape() Shape
}

// RectangleFactory struct that implements the ShapeFactory interface
type RectangleFactory struct{}

func (rf *RectangleFactory) CreateShape() Shape {
	return &Rectangle{}
}

// CircleFactory struct that implements the ShapeFactory interface
type CircleFactory struct{}

func (cf *CircleFactory) CreateShape() Shape {
	return &Circle{}
}

func main() {
	// Create a RectangleFactory
	rectangleFactory := &RectangleFactory{}

	// Use the RectangleFactory to create a Rectangle
	rectangle := rectangleFactory.CreateShape()

	// Draw the Rectangle
	fmt.Println(rectangle.Draw())

	// Create a CircleFactory
	circleFactory := &CircleFactory{}

	// Use the CircleFactory to create a Circle
	circle := circleFactory.CreateShape()

	// Draw the Circle
	fmt.Println(circle.Draw())
}
