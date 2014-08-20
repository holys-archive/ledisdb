package tokuft

import (
	"github.com/siddontang/go-tokuft/tokuft"
	"github.com/siddontang/ledisdb/store/driver"
)

type Tx struct {
	db *tokuft.DB
	tx *tokuft.Tx
}

func newTx(db *DB) (*Tx, error) {
	tx, err := db.env.BeginTx(nil, uint(0))
	if err != nil {
		return nil, err
	}

	return &Tx{db.db, tx}, nil
}

func (t *Tx) Get(key []byte) ([]byte, error) {
	return t.tx.Get(t.db, key)
}

func (t *Tx) Put(key []byte, value []byte) error {
	return t.tx.Put(t.db, key, value)
}

func (t *Tx) Delete(key []byte) error {
	return t.tx.Delete(t.db, key)
}

func (t *Tx) newIterator() *Iterator {
	c, err := t.tx.Cursor(t.db)
	if err != nil {
		return &Iterator{nil, nil, nil, nil, false, err, false}
	}

	return &Iterator{t.tx, c, nil, nil, true, nil, false}
}

func (t *Tx) NewIterator() driver.IIterator {
	return t.newIterator()
}

func (t *Tx) NewWriteBatch() driver.IWriteBatch {
	return driver.NewWriteBatch(t)

}

func (t *Tx) BatchPut(writes []driver.Write) error {
	var err error
	for _, w := range writes {
		if w.Value == nil {
			err = t.tx.Delete(t.db, w.Key)
		} else {
			err = t.tx.Put(t.db, w.Key, w.Value)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Tx) Commit() error {
	return t.tx.Commit()

}

func (t *Tx) Rollback() error {
	t.tx.Abort()
	return nil
}
