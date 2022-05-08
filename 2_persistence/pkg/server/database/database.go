package database

import "context"

// Database represents the operations that are done on a Database.
// This interface abstracts the underlying implementation.
type Database interface {
	Create(ctx context.Context, data []byte) error
}
