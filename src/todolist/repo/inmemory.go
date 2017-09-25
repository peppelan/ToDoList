package repo

import (
	"fmt"
	"todolist/spi"
)

type InMemoryRepo struct {
	currentId int
	todos spi.Todos
}

// Give us some seed data
func NewInMemoryRepo() *InMemoryRepo {
	r := new(InMemoryRepo)

	r.Create(spi.Todo{Name: "Write presentation"})
	r.Create(spi.Todo{Name: "Host meetup"})

	return r
}

func (r *InMemoryRepo) Find(id string) spi.Todo {
	for _, t := range r.todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return spi.Todo{}
}

func (r *InMemoryRepo) FindAll() spi.Todos {
	return r.todos
}

func (r *InMemoryRepo) Create(t spi.Todo) spi.Todo {
	r.currentId += 1
	t.Id = fmt.Sprintf("%d", r.currentId)
	r.todos = append(r.todos, t)
	return t
}

func (r *InMemoryRepo) Destroy(id string) bool {
	for i, t := range r.todos {
		if t.Id == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return true
		}
	}
	return false
}
