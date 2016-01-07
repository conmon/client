package api

import (
	"flag"
)

var apiPortFlag = flag.Int("api-port", 23563, "Port that the conmon client api should be exposed at")
