//go:build windows

package proxy

import (
	"syscall"
)

// generateDefaultListenControl is nil on Windows
func generateDefaultListenControl(_ Config) func(network, address string, c syscall.RawConn) error {
	return nil
}
