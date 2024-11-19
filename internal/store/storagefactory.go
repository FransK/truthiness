package store

// Use the StoreFactory to create the right type of storage
type StoreFactory interface {
	NewStore() Storage
}
