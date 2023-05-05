//go:build unix

package proxy

import (
	"github.com/AdguardTeam/golibs/errors"
	"github.com/AdguardTeam/golibs/log"
	"syscall"
)

// generateDefaultListenControl is not nil on Unix
func generateDefaultListenControl(pc Config) func(network, address string, c syscall.RawConn) error {
	if !pc.ReuseAddr {
		return nil
	}
	return func(network, address string, c syscall.RawConn) error {
		var errSysCall error
		errControl := c.Control(func(fd uintptr) {
			// Enable SO_REUSEADDR
			errSysCall = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
			if errSysCall != nil {
				log.Error("Could not set SO_REUSEADDR socket option: %s", errSysCall)
			}
		})
		return errors.WithDeferred(errSysCall, errControl)
	}
}
