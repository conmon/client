package conmon

import (
	"github.com/conmon/client/container"
	"github.com/conmon/client/database"
	"github.com/conmon/client/system"

	"fmt"
)

// init initialized a driver for the particular database
func init() {
	driver := conmonDriver{}

	database.RegisterDriver("conmon-server", driver)
}

// conmonDriver struct implements the database.Driver interface
type conmonDriver struct {
}

func (c conmonDriver) Connect() error {
	return nil
}

// SendContainerStats formats the container statistics for sending
func (c conmonDriver) SendContainerStats(cstats container.Stats) error {
	fmt.Println(cstats.Uptime)
	return nil
}

func (c conmonDriver) SendSystemStats(cstats system.Stats) error {
	fmt.Println(cstats.Uptime)
	return nil
}

func (c conmonDriver) Close() error {
	return nil
}
