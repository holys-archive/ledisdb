package tokuft

import (
	"github.com/siddontang/go-tokuft/tokuft"
)

type Iterator struct {
	tx              *tokuft.Tx
	c               *tokuft.Cursor
	key             []byte
	value           []byte
	valid           bool
	err             error
	closeAutoCommit bool
}

func (it *Iterator) Close() error {
	if err := it.c.Close(); err != nil {
		it.tx.Abort()
		return err
	}

	if !it.closeAutoCommit {
		return it.err
	}

	if it.err != nil {
		it.tx.Abort()
		return it.err
	}

	return it.tx.Commit()
}

func (it *Iterator) First() {
	it.key, it.value, it.err = it.c.Get(nil, nil, tokuft.FIRST)
	it.setState()
}

func (it *Iterator) Last() {
	it.key, it.value, it.err = it.c.Get(nil, nil, tokuft.LAST)
	it.setState()
}

func (it *Iterator) Next() {
	it.key, it.value, it.err = it.c.Get(nil, nil, tokuft.NEXT)
	it.setState()
}

func (it *Iterator) Prev() {
	it.key, it.value, it.err = it.c.Get(nil, nil, tokuft.PREV)
	it.setState()
}

func (it *Iterator) Seek(key []byte) {
	it.key, it.value, it.err = it.c.Get(key, nil, tokuft.SET_RANGE)
	it.setState()
}

func (it *Iterator) setState() {
	if it.err != nil {
		if it.err == tokuft.NOTFOUND {
			it.err = nil
		}
		it.valid = false
	} else {
		it.valid = !(it.key == nil && it.value == nil)
	}
}

func (it *Iterator) Valid() bool {
	return it.valid
}

func (it *Iterator) Error() error {
	return it.err
}

func (it *Iterator) Key() []byte {
	return it.key
}

func (it *Iterator) Value() []byte {
	return it.value
}
