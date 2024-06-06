package rdpTable

import (
	"sync"
)

var Table sync.Map

type Info struct {
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	FrameRate uint32 `json:"frame_rate"`

	SetStream func(stream bool) error
	// func([]bytes)
	Listener *sync.Map
}

func RdpRegister(info *Info) (unregister func()) {
	Table.Store(info.Name, info)
	return func() {
		Table.CompareAndDelete(info.Name, info)
	}
}

func ListenRegister(name string, onFrame func([]byte)) {

}
