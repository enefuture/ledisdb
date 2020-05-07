package store

import (
	"github.com/enefuture/ledisdb/store/driver"
)

type Slice interface {
	driver.ISlice
}
