package repo

import (
	"fmt"
	"todolist/spi"
)

type InMemoryRepo struct {
	currentId int
	todos map[string] *spi.Todo
}

func NewInMemoryRepo() *InMemoryRepo {
	r := new(InMemoryRepo)
	return r
}

func (r *InMemoryRepo) Init() error {
	r.todos = make(map[string] *spi.Todo)
	return nil
}

func (r *InMemoryRepo) Find(id string) *spi.Todo {
	return r.todos[id]
}

func (r *InMemoryRepo) FindAll() map[string] *spi.Todo {
	// FIXME: do not use for production code :-)
	return r.todos
}

func (r *InMemoryRepo) Create(t spi.Todo) string {
	r.currentId += 1
	id := fmt.Sprintf("%d", r.currentId)
	r.todos[id] = &t
	return id
}

func (r *InMemoryRepo) Destroy(id string) bool {
	isPresent := nil != r.todos[id]
	delete(r.todos, id)
	return isPresent
}

func (r *InMemoryRepo) Update(id string, todo spi.Todo) bool {
	isPresent := nil != r.todos[id]
	if isPresent {
		r.todos[id] = &todo
		return true
	} else {
		return false
	}
}