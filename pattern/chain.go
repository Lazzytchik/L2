package pattern

import "fmt"

// Handler is the interface that all handlers must implement
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request int)
}

// ConcreteHandlerA is a concrete handler that implements the Handler interface
type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *ConcreteHandlerA) Handle(request int) {
	if request >= 0 && request < 10 {
		fmt.Printf("ConcreteHandlerA handled the request: %d\n", request)
	} else if h.next != nil {
		h.next.Handle(request)
	} else {
		fmt.Println("No handler can handle the request.")
	}
}

// ConcreteHandlerB is another concrete handler that implements the Handler interface
type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *ConcreteHandlerB) Handle(request int) {
	if request >= 10 && request < 20 {
		fmt.Printf("ConcreteHandlerB handled the request: %d\n", request)
	} else if h.next != nil {
		h.next.Handle(request)
	} else {
		fmt.Println("No handler can handle the request.")
	}
}

// ConcreteHandlerC is yet another concrete handler that implements the Handler interface
type ConcreteHandlerC struct {
	next Handler
}

func (h *ConcreteHandlerC) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *ConcreteHandlerC) Handle(request int) {
	if request >= 20 && request < 30 {
		fmt.Printf("ConcreteHandlerC handled the request: %d\n", request)
	} else if h.next != nil {
		h.next.Handle(request)
	} else {
		fmt.Println("No handler can handle the request.")
	}
}

func main() {
	concreteHandlerA := &ConcreteHandlerA{}
	concreteHandlerB := &ConcreteHandlerB{}
	concreteHandlerC := &ConcreteHandlerC{}

	// Create a chain of handlers
	concreteHandlerA.SetNext(concreteHandlerB).SetNext(concreteHandlerC)

	// Handle requests with the first handler in the chain
	concreteHandlerA.Handle(5)
	concreteHandlerA.Handle(15)
	concreteHandlerA.Handle(25)
}
