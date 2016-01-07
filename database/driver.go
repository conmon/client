package database

import (
	"github.com/conmon/client/container"
	"github.com/conmon/client/system"
)

// Driver interface contains the methods that can be used to send data to particular
// servers
type Driver interface {
	Connect() error
	SendContainerStats(container.Stats) error
	SendSystemStats(system.Stats) error
	Close() error
}
