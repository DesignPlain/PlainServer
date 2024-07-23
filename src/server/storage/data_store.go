package storage

import (
	"DesignSphere_Server/src/utils"
	"os"

	"github.com/cockroachdb/pebble"
)

type Store struct {
	// Handler to a pebble database connection
	db_Handler *pebble.DB

	// Database directory path
	db_Directory string
}

// Initialize the database store for the given DB directory `db_Dir`
func Init(db_Dir string) (*Store, error) {
	utils.Log(utils.DEBUG, "Initializing DB directory.")
	if _, err := os.Stat(db_Dir); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(db_Dir, os.ModePerm); err != nil {
			utils.Log(utils.ERROR, UnableToAccessDBDirectory.Error()+", error: "+err.Error())
			return nil, UnableToAccessDBDirectory
		}
	}

	db, err := pebble.Open(db_Dir, &pebble.Options{})
	if err != nil {
		utils.Log(utils.ERROR, DBInitFailed.Error()+", error: "+err.Error())
		return nil, DBInitFailed
	}

	return &Store{db_Handler: db, db_Directory: db_Dir}, nil
}

func (store *Store) Close() error {
	if store.db_Handler != nil {
		store.db_Handler.Close()
		if err := store.db_Handler.Close(); err != nil {
			utils.Log(utils.ERROR, DBCloseFailed.Error()+", error: "+err.Error())
			return DBCloseFailed
		}

		store.db_Handler = nil
		return nil
	}

	return DBNotInitialized
}

func (store *Store) Set(key []byte, value []byte) error {
	if store.db_Handler != nil {
		//utils.Log(utils.DEBUG, string(value))
		err := store.db_Handler.Set(key, value, pebble.Sync)
		if err != nil {
			utils.Log(utils.ERROR, DBWriteFailed.Error()+", error: "+err.Error())
			return DBWriteFailed
		}

		store.db_Handler.Flush()

		return nil
	}
	return DBNotInitialized
}

func (store *Store) Get(key []byte) ([]byte, error) {
	if store.db_Handler != nil {
		value, closer, err := store.db_Handler.Get([]byte(key))

		//utils.Log(utils.DEBUG, string(value))
		if err != nil {
			if err == pebble.ErrNotFound {
				return nil, DBKeyNotFound
			}

			utils.Log(utils.ERROR, DBWriteFailed.Error()+", error: "+err.Error())
			return nil, DBReadFailed
		}

		// Not calling close can cause memory leak.
		if err := closer.Close(); err != nil {
			utils.Log(utils.ERROR, DBWriteFailed.Error()+", error: "+err.Error())
			return nil, err
		}
		return value, nil
	}
	return nil, DBNotInitialized
}

/*
Storage Error Type enum
*/
type StorageError int64

const (
	DBInitFailed StorageError = iota
	DBCloseFailed
	DBNotInitialized
	UnableToAccessDBDirectory
	DBWriteFailed
	DBReadFailed
	DBKeyNotFound
	DBReadCloserFailed
)

func (err StorageError) Error() string {
	switch err {
	case DBInitFailed:
		return "Initializing database directory failed."
	case DBCloseFailed:
		return "Closing database connection failed."
	case DBNotInitialized:
		return "Database is not initialized."
	case UnableToAccessDBDirectory:
		return "Unable to access DB directory."
	case DBWriteFailed:
		return "Writing to the database failed."
	case DBReadFailed:
		return "Reading from the database failed."
	case DBKeyNotFound:
		return "Requested key not found in the database."
	case DBReadCloserFailed:
		return "Closing the database reader failed."
	}

	return "Unknown Storage Error"
}
