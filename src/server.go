package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/swarm-io/httpdumps/appconfig"
	"github.com/swarm-io/httpdumps/handlers/dumps"
	"github.com/swarm-io/httpdumps/handlers/health"
)

func main() {
	zerolog.SetGlobalLevel(appconfig.GetLogLevel())

	http.HandleFunc("/dumps", dumps.Dumps)

	// If used in some orchestration world like kubernetes
	http.HandleFunc("/health", health.Health)

	hostStr := ":" + appconfig.DumpsPort

	log.Info().Msg(fmt.Sprintf("INFO: Listening on '%s'", hostStr))
	http.ListenAndServe(hostStr, nil)
}
