package spi

// Defines the interface that a repository of Todos should implement
type Repo interface {
   Find(id int) Todo
   FindAll() Todos
   Create(t Todo) Todo
   Destroy(id int) error
}
