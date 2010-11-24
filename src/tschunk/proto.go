package proto

import "os"

type Transaction struct{}

type Protocol interface {
	DbLookup(key string) (val string, err os.Error)
	DbUpdate(key string, oldVal string, newVal string) (ok bool, err os.Error)
	TxStart(id string) (tx Transaction, err os.Error)
	TxUpdate(id string, key string, oldVal string, newVal string) (err os.Error)
	TxCommit(id string, commit bool)
}
