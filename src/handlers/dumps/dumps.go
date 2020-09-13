package dumps

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/rs/zerolog/log"
)

func Dumps(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg(fmt.Sprintf("INFO: Dumping Request: %s %s", r.Proto, r.Method))
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("ERROR: Failed to dump request: %s", err))
	}
	fmt.Println(string(dump))
	log.Info().Msg(string(dump))
	fmt.Fprintf(w, "OK")
}
