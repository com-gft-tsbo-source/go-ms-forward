package msforward

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/com-gft-tsbo-source/go-common/ms-framework/microservice"
)

// ###########################################################################
// ###########################################################################
// MsForward
// ###########################################################################
// ###########################################################################

// MsForward Encapsulates the ms-forward data
type MsForward struct {
	microservice.MicroService

	seededRand   *rand.Rand
	forwardMutex sync.Mutex
}

// ###########################################################################

// InitMsForwardFromArgs ...
func InitFromArgs(ms *MsForward, args []string, flagset *flag.FlagSet) *MsForward {
	var cfg Configuration

	if flagset == nil {
		flagset = flag.NewFlagSet("ms-forward", flag.PanicOnError)
	}

	InitConfigurationFromArgs(&cfg, args, flagset)
	microservice.Init(&ms.MicroService, &cfg.Configuration, nil)
	ms.seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	forwardHandler := ms.DefaultHandler()
	forwardHandler.Get = ms.httpGetForward
	ms.AddHandler("/forward", forwardHandler)
	return ms
}

// ---------------------------------------------------------------------------

var deviceMutex sync.Mutex

func (ms *MsForward) httpGetForward(w http.ResponseWriter, r *http.Request) (status int, contentLen int, msg string) {
	ms.forwardMutex.Lock()
	value := ms.seededRand.Intn(100)
	ms.forwardMutex.Unlock()
	status = http.StatusOK
	name := r.Header.Get("X-Cid")
	version := r.Header.Get("X-Version")
	environment := r.Header.Get("X-Environment")
	msg = fmt.Sprintf("'v%s' in '%s' Generated random number '%d' for client '%s@%s'.", ms.GetVersion(), environment, value, name, version)
	response := NewForwardResponse(msg, ms)
	response.Value = value
	ms.SetResponseHeaders("application/json; charset=utf-8", w, r)
	w.WriteHeader(status)
	contentLen = ms.Reply(w, response)
	return status, contentLen, msg
}
