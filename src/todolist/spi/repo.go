package spi

// Defines the interface that a repository of Todos should implement
type Repo interface {
   Find(id string) Todo
   FindAll() Todos
   Create(t Todo) Todo
   Destroy(id string) error
}
