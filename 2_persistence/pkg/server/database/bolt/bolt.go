package bolt

import (
	"context"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

// Bolt is the Bolt database.
// It satisfies the Database interface.
type Bolt struct {
	db *bolt.DB
}

const (
	dbName     = "test.db"
	bucketName = "users"
)

// New returns a new Bolt implementation.
func New(ctx context.Context, directory string) (*Bolt, error) {
	db, err := bolt.Open(fmt.Sprintf("%s/%s", directory, dbName), 0600, nil)
	if err != nil {
		return nil, err
	}

	// Ensure that the bucket exists, if not create it.
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Bolt{
		db: db,
	}, nil
}

// Close closes the database.
// Make sure to close the database once used.
func (b *Bolt) Close(ctx context.Context) {
	b.db.Close()
}

// Create implements the Database interface.
func (b *Bolt) Create(ctx context.Context, data []byte) error {

	fmt.Println(string(data))

	return nil
}
