package manager

import (
	"github.com/conmon/client/container"
	"github.com/conmon/client/database"
	"github.com/conmon/client/system"
)

type machineStruct struct {
	containers []container.Stats
	stop       chan bool
	db         database.Driver
}

func (m *machineStruct) monitor() {
	for {
		select {
		case <-m.stop:
			m.cleanup()
			return
		default:
			m.gatherdata()
		}
	}
}

func (m *machineStruct) cleanup() {

}

func (m *machineStruct) gatherdata() {
	m.db.SendContainerStats(container.Stats{Uptime: 1})
	m.db.SendSystemStats(system.Stats{Uptime: 1})
}
