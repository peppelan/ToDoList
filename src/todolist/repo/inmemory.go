package repo

import (
	"fmt"
	"todolist/spi"
)

type InMemoryRepo struct {
	currentId int
	todos spi.Todos
}

func NewInMemoryRepo() *InMemoryRepo {
	r := new(InMemoryRepo)
	return r
}

func (r *InMemoryRepo) Init() error {
	return nil
}

func (r *InMemoryRepo) Find(id string) *spi.Todo {
	for _, t := range r.todos {
		if t.Id == id {
			return &t
		}
	}
	// return nil if not found
	return nil
}

func (r *InMemoryRepo) FindAll() spi.Todos {
	return r.todos
}

func (r *InMemoryRepo) Create(t spi.Todo) string {
	r.currentId += 1
	t.Id = fmt.Sprintf("%d", r.currentId)
	r.todos = append(r.todos, t)
	return t.Id
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

func (r *InMemoryRepo) Update(todo spi.Todo) bool {
	for i, t := range r.todos {
		if t.Id == todo.Id {
			r.todos[i] = todo
			return true
		}
	}
	return false
}