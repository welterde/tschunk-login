package proto

import "os"

type Protocol interface {
	DbLookup(key string) (val string, err os.Error)
}
