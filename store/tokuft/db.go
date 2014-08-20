package tokuft

import (
	"github.com/siddontang/go-tokuft/tokuft"
	"github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/store/driver"
	"os"
)

type Store struct {
}

func (s Store) String() string {
	return DBName
}

func (s Store) Open(path string, c *config.Config) (driver.IDB, error) {
	env, err := tokuft.NewEnv()
	if err != nil {
		return &DB{}, err
	}

	if _, err := os.Stat(path); err != nil {
		err = os.MkdirAll(path, 0644)
		if err != nil {
			return &DB{}, err
		}
	}

	err = env.Open(path, tokuft.CREATE|tokuft.PRIVATE|tokuft.INIT_MPOOL|tokuft.INIT_LOCK|tokuft.INIT_TXN, 0644)
	if err != nil {
		return &DB{}, err
	}

	tx, err := env.BeginTx(nil, uint(0))
	if err != nil {
		return &DB{}, err
	}

	dbi, err := tx.OpenDB("ledis_tokuft.db", tokuft.CREATE, 0644)
	if err != nil {
		return &DB{}, err
	}

	if err := tx.Commit(); err != nil {
		return &DB{}, err
	}

	db := &DB{
		env:  env,
		db:   dbi,
		path: path,
	}

	return db, nil

}

func (s Store) Repair(path string, c *config.Config) error {
	println("tokuft not support repair")
	return nil
}

type DB struct {
	env  *tokuft.Env
	db   *tokuft.DB
	path string
	cfg  *config.Config
}

func (db *DB) Close() error {
	if err := db.db.Close(); err != nil {
		println("db.close", err.Error())
		return err
	}
	if err := db.env.Close(); err != nil {
		println("env.close", err.Error())
		return err
	}

	return nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	tx, err := db.env.BeginTx(nil, tokuft.TXN_READ_ONLY|tokuft.READ_UNCOMMITTED)
	if err != nil {
		return nil, err
	}

	defer tx.Commit()

	v, err := tx.Get(db.db, key)
	if err != nil {
		return nil, err
	} else if err == tokuft.NOTFOUND {
		return nil, nil
	}

	return v, nil
}

func (db *DB) Put(key []byte, value []byte) error {
	tx, err := db.env.BeginTx(nil, tokuft.READ_UNCOMMITTED)
	if err != nil {
		return err
	}

	defer tx.Abort()

	err = tx.Put(db.db, key, value)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (db *DB) BatchPut(writes []driver.Write) error {
	tx, err := db.env.BeginTx(nil, tokuft.READ_UNCOMMITTED)
	if err != nil {
		return err
	}

	defer tx.Abort()

	for _, w := range writes {
		if w.Value == nil {
			err = tx.Delete(db.db, w.Key)
		} else {
			err = tx.Put(db.db, w.Key, w.Value)
		}

		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (db *DB) Delete(key []byte) error {
	tx, err := db.env.BeginTx(nil, tokuft.READ_UNCOMMITTED)
	if err != nil {
		return err
	}

	defer tx.Abort()

	err = tx.Delete(db.db, key)
	if err != nil {
		return err
	} else if err == tokuft.NOTFOUND {
		return nil
	}

	return tx.Commit()
}

func (db *DB) iterator() *Iterator {
	// flags := uint(0)
	// if rdonly {
	// 	flags = (tokuft.TXN_READ_ONLY | tokuft.TXN_SNAPSHOT)
	// 	// flags = tokuft.TXN_READ_ONLY
	// }
	tx, err := db.env.BeginTx(nil, tokuft.TXN_READ_ONLY|tokuft.READ_UNCOMMITTED)
	if err != nil {
		return &Iterator{nil, nil, nil, nil, false, err, true}
	}

	c, err := tx.Cursor(db.db)
	if err != nil {
		tx.Abort()
		return &Iterator{nil, nil, nil, nil, false, err, true}
	}

	return &Iterator{tx, c, nil, nil, true, nil, true}
}

func (db *DB) NewIterator() driver.IIterator {
	return db.iterator()
}

func (db *DB) NewWriteBatch() driver.IWriteBatch {
	return driver.NewWriteBatch(db)
}

func (db *DB) Begin() (driver.Tx, error) {
	return newTx(db)
}

func init() {
	driver.Register(Store{})
}
