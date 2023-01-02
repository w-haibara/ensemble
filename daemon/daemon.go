package daemon

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/w-haibara/ensemble/docker"
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

	client, err := docker.NewClient()
	if err != nil {
		log.Println("container-run failed:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	ctx := context.Background()
	request := docker.NewRequestFromURLValues(r.URL.Query())

	switch task {
	case "container-run":
		if err := client.ContainerRun(ctx, request); err != nil {
			log.Println("container-run failed:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "container-start":
		if err := client.ContainerStart(ctx, request); err != nil {
			log.Println("container-start failed:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "container-stop":
		if err := client.ContainerStop(ctx, request); err != nil {
			log.Println("container-stop failed:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "container-remove":
		if err := client.ContainerRm(ctx, request); err != nil {
			log.Println("container-remove failed:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		log.Println("unkown task name:", task)
		w.WriteHeader(http.StatusBadRequest)
	}
}
