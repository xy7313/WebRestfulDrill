package main

import "time"

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	// Url      string    `json:"url"`
	// Count bool      `json:"count"`
}

type Todos []Todo
