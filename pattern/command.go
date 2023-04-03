package pattern

import "fmt"

/*
	Паттерн Команда (Command)
	Поведенческий паттерн позволяющий представлять действия как объекты, передавать в запросы, логировать и отменять.
*/

// Command is the interface that all commands must implement
type Command interface {
	Execute()
}

// AddCommand is a concrete command that implements the Command interface
type AddCommand struct {
	calculator *Calculator
	value      int
}

func (c *AddCommand) Execute() {
	c.calculator.Add(c.value)
}

// SubtractCommand is another concrete command that implements the Command interface
type SubtractCommand struct {
	calculator *Calculator
	value      int
}

func (c *SubtractCommand) Execute() {
	c.calculator.Subtract(c.value)
}

// Calculator is the object that the commands act upon
type Calculator struct {
	total int
}

func (c *Calculator) Add(value int) {
	c.total += value
}

func (c *Calculator) Subtract(value int) {
	c.total -= value
}

// Invoker is the object that executes the commands
type Invoker struct {
	commands []Command
}

func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func main() {
	calculator := &Calculator{}

	addCommand := &AddCommand{calculator: calculator, value: 5}
	subtractCommand := &SubtractCommand{calculator: calculator, value: 2}

	invoker := &Invoker{}
	invoker.StoreCommand(addCommand)
	invoker.StoreCommand(subtractCommand)

	invoker.ExecuteCommands()

	fmt.Printf("Total: %d\n", calculator.total)
}
