package database

import (
	"flag"
	"fmt"
)

var (
	// Database flags
	TypeFlag    = flag.String("db-type", "conmon-server", "backend to use")
	AddrFlag    = flag.String("db-addr", "localhost:23564", "Server address")
	UserFlag    = flag.String("db-user", "conmon", "User to connect to database")
	PassFlag    = flag.String("db-pass", "conmon", "Password to connect to database")
	NameFlag    = flag.String("db-name", "conmon", "Database name")
	SecureFlag  = flag.Bool("db-secr", true, "Database security enabled")
	TimeoutFlag = flag.Int("db-timeout", 30, "Database connection timeout (seconds)")
	RetryFlag   = flag.Int("db-retry", 3, "Amount of times the database connection would fail before ConMon shuts down")
)

// Contain the list of database drivers in a map
var drivers = map[string](Driver){}

// RegisterDriver registers a driver which implements database.Driver.
// It adds it to the list of available drivers
func RegisterDriver(name string, driver Driver) {
	drivers[name] = driver
}

// New returns a Driver interface that can be used to send data to appropate servers
// If a driver is not found it returns an error
func New() (Driver, error) {
	name := *TypeFlag
	if name == "" {
		return nil, fmt.Errorf("driver name cannot be empty")
	}

	driver, ok := drivers[name]
	if !ok {
		return nil, fmt.Errorf("unknown backend storage driver %s", name)
	}

	return driver, nil
}
