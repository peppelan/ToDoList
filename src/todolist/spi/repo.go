package spi

// Defines the interface that a repository of Todos should implement
type Repo interface {

   // Initializes the repository.
   // This method, once the initialization is finished, will return 'nil' if the initialization is successful,
   // or an explanatory error otherwise.
   Init() error

   // Fetches a given to-do; returns nil when not found
   Find(id string) *Todo

   // Fetches all to-do's
   FindAll() Todos

   // Creates a to-do.
   // Returns the ID of the created object
   Create(t Todo) string

   // Removes the to-do with the given ID from the repo.
   // Returns true when the object was found and removed, false when the object was not found.
   Destroy(id string) bool

   // Updates a to-do.
   // Returns true when the object was found and updated, false when the object was not found.
   Update(t Todo) bool
}
