package ui

import (
	"github.com/koho/frpmgr/pkg/config"
	"github.com/lxn/walk"
)

// quickAddBinder is the view model of quick-add dialog
type quickAddBinder struct {
	RemotePort   string
	LocalAddr    string
	LocalPort    string
	LocalPortMin string
	LocalPortMax string
	Dir          string
}

// QuickAdd is the interface that must be implemented to build a quick-add dialog
type QuickAdd interface {
	// Run a new simple dialog to input few parameters
	Run(owner walk.Form) (int, error)
	// GetProxies returns the proxies generated by quick-add dialog
	GetProxies() []*config.Proxy
}