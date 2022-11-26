package daemon

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var unixDomainSockPath = filepath.Join(os.TempDir(), "ensemble.sock")

func Serve() error {
	if _, err := os.Stat(unixDomainSockPath); err == nil {
		if err := os.Remove(unixDomainSockPath); err != nil {
			return err
		}
	}

	l, err := net.Listen("unix", unixDomainSockPath)
	if err != nil {
		log.Fatalln(err)
	}

	return http.Serve(l, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	task := strings.TrimPrefix(r.URL.String(), "/")
	log.Println("task:", task)

	switch task {
	case "task-01":
		log.Println("start task-01")
	default:
		log.Println("unkown task name:", task)
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}
