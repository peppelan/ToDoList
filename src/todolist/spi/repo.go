package spi

// Defines the interface that a repository of Todos should implement
type Repo interface {
   Find(id string) Todo
   FindAll() Todos
   Create(t Todo) Todo

   // Removes the to-do with the given ID from the repo.
   // returns true when the object was found and removed, false when the object was not found,
   Destroy(id string) bool
}
