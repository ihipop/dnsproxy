package proxy

import (
	"net"
)

// ListenConfig returns the default [net.ListenConfig]
func ListenConfig(c Config) (lc *net.ListenConfig) {
	return &net.ListenConfig{
		Control: generateDefaultListenControl(c),
	}
}
