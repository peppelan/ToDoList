package repo

import "todolist/spi"

// Creates the appropriate repository depending on the application configuration
func NewRepo() spi.Repo {
	r := NewInMemoryRepo()
	r.Init()
	return r
}