package main

import "time"

// Domain object for my application
type Todo struct {
	// Short name of the to-do
	Name		string		`json:"name"`

	// Full description of the to-do
	Description	string		`json:"description"`

	// Flag for indicating it has been done
	Completed	bool     	`json:"completed"`

	// Deadline for the to-do
	Due			time.Time	`json:"due"`
}

type Todos []Todo