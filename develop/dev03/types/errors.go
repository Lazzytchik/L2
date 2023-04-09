package dev03

type NonUniqueError struct {
}

func (e NonUniqueError) Error() string {
	return "same element"
}
