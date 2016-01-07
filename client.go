package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/conmon/client/database"
	"github.com/conmon/client/manager"
	"github.com/conmon/client/version"
	"github.com/golang/glog"
)

var (
	// Client running functions
	useApiFlag  = flag.Bool("use-api", false, "Use the ConMon client API or not")
	versionFlag = flag.Bool("version", false, "print version of ConMon client")
	verboseFlag = flag.Bool("verbose", false, "Enable Verbose logging")
)

func main() {
	// Store errors
	var err error

	// TODO: Write comment for this
	defer glog.Flush()
	flag.Parse()

	//If version flag is specified show the version and then quit
	if *versionFlag {
		fmt.Printf("ConMon client version %s (r%s)\n", version.Version, version.Revision)
		os.Exit(0)
	}

	// Create a new Database driver
	db, err := database.New()
	if err != nil {
		glog.Fatal("cannot create database: ", err)
	}

	// Connect to database
	if err = db.Connect(); err != nil {
		glog.Fatalf("cannot connect to %s: %s\n", *database.TypeFlag, err)
	} else {
		glog.Infof("connected to %s database\n")
	}

	if err = manager.Start(db); err != nil {
		glog.Fatalf("Unable to start the manager: %s\n", err)
	}

	handleSignals()
}

func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		sig := <-c
		if err := manager.Stop(); err != nil {
			glog.Errorln("manager failed to stop: ", err)
		}
		glog.Infof("stopping manager, %s given", sig)
		os.Exit(0)
	}()
}
