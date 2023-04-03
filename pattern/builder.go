package pattern

import "fmt"

/*
	Паттерн Строитель (Builder)
	Порождающий паттерн проектирования позволяющий создавать сложный объект постепенно.

	Реальные сценарии использования
		- Построение запросов к базе через ORM
		- Построение веб-форм
		- Построение объектов настоек (конфиги всякие)
*/

// Product represents the complex object being built
type Product struct {
	Part1 string
	Part2 int
	Part3 bool
}

// Builder defines the interface for building the product
type Builder interface {
	SetPart1(string) Builder
	SetPart2(int) Builder
	SetPart3(bool) Builder
	Build() *Product
}

// ConcreteBuilder implements the Builder interface and provides an implementation for building the product
type ConcreteBuilder struct {
	product *Product
}

func (b *ConcreteBuilder) SetPart1(part1 string) Builder {
	b.product.Part1 = part1
	return b
}

func (b *ConcreteBuilder) SetPart2(part2 int) Builder {
	b.product.Part2 = part2
	return b
}

func (b *ConcreteBuilder) SetPart3(part3 bool) Builder {
	b.product.Part3 = part3
	return b
}

func (b *ConcreteBuilder) Build() *Product {
	return b.product
}

// Director defines the interface for using the builder to construct the product
type Director struct {
	builder Builder
}

func (d *Director) Construct() *Product {
	return d.builder.SetPart1("part1").SetPart2(2).SetPart3(true).Build()
}

func main() {
	builder := &ConcreteBuilder{product: &Product{}}
	director := &Director{builder: builder}

	product := director.Construct()
	fmt.Printf("%+v\n", product)
}
