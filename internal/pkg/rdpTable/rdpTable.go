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

func Register(info *Info) {
	Table.Store(info.Name, info)
}
