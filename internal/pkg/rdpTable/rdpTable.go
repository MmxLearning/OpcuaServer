package rdpTable

import (
	"sync"
)

var Table sync.Map

type Info struct {
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	FrameRate uint32 `json:"frame_rate"`

	SetStream func(stream bool) error `json:"-"`
	// func([]bytes)
	Listener *sync.Map `json:"-"`
}

func RdpRegister(info *Info) (unregister func()) {
	Table.Store(info.Name, info)
	return func() {
		Table.CompareAndDelete(info.Name, info)
	}
}

func LoadRdp(name string) (*Info, bool) {
	val, ok := Table.Load(name)
	if !ok {
		return nil, false
	}
	return val.(*Info), true
}

func ListenRegister(name, listener string, onFrame func([]byte)) (unregister func(), ok bool) {
	info, ok := LoadRdp(name)
	if !ok {
		return nil, false
	}
	info.Listener.Store(listener, onFrame)
	return func() {
		info.Listener.CompareAndDelete(listener, onFrame)
	}, true
}

func RdpList() []*Info {
	var result = make([]*Info, 0)
	Table.Range(func(_, value any) bool {
		result = append(result, value.(*Info))
		return true
	})
	return result
}
