package main

type Options interface {
	print()
}

type OptionRange struct {
	min int
	max int
}

type Commit struct {
	changeType string
	emoji      string
	message    string
}
