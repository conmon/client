package manager

import (
	"github.com/conmon/client/database"
)

//TODO: Entire Manager

var machine *machineStruct

// Start function starts the manager
func Start(db database.Driver) error {
	machine := machineStruct{
		db: db,
	}

	go machine.monitor()

	return nil
}

func Stop() error {
	return nil
}
