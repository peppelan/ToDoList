package repo

import (
	"fmt"
	"todolist/spi"
)

var currentId int

var todos spi.Todos

// Give us some seed data
func init() {
	Create(spi.Todo{Name: "Write presentation"})
	Create(spi.Todo{Name: "Host meetup"})
}

func Find(id int) spi.Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return spi.Todo{}
}

func FindAll() spi.Todos {
	return todos
}

func Create(t spi.Todo) spi.Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func Destroy(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
