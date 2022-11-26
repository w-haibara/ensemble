package daemon

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
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

	h := func(_ http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
	}

	return http.Serve(l, http.HandlerFunc(h))
}
