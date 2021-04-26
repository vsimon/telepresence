// +build !windows

package client

import (
	"context"
	"net"
	"os"

	"github.com/telepresenceio/telepresence/v2/pkg/proc"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
)

const (
	// ConnectorSocketName is the path used when communicating to the connector process
	ConnectorSocketName = "/tmp/telepresence-connector.socket"

	// DaemonSocketName is the path used when communicating to the daemon process
	DaemonSocketName = "/var/run/telepresence-daemon.socket"
)

// DialSocket dials the given unix socket and returns the resulting connection
func DialSocket(c context.Context, socketName string) (*grpc.ClientConn, error) {
	return grpc.DialContext(c, "unix:"+socketName,
		grpc.WithInsecure(),
		grpc.WithNoProxy(),
		grpc.WithBlock(),
	)
}

// ListenSocket returns a listener for the given named pipe and returns the resulting connection
func ListenSocket(_ context.Context, socketName string) (net.Listener, error) {
	if proc.IsAdmin() {
		origUmask := unix.Umask(0)
		defer unix.Umask(origUmask)
	}
	return net.Listen("unix", socketName)
}

// SocketExists returns true if a socket is found at the given path
func SocketExists(path string) bool {
	s, err := os.Stat(path)
	return err == nil && s.Mode()&os.ModeSocket != 0
}
