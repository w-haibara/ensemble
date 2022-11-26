package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/w-haibara/ensemble/daemon"
	"github.com/w-haibara/ensemble/systemd"
)

func main() {
	switch os.Args[1] {
	case "deploy":
		deploy()
	case "notify":
		notify(os.Args[2])
	case "serve":
		serve()
	}
}

func deploy() {
	units := []systemd.Unit{
		{
			Name: "sample1",
			Body: heredoc.Doc(`
			[Unit]
			Description=Sample

			[Service]
			Type=exec
			ExecStart=/bin/bash -c "ensemble notify task-01"

			[Install]
			WantedBy=default.target
			`),
		},
	}
	if err := systemd.Start(units); err != nil {
		log.Fatalln(err)
	}
}

func notify(name string) {
	logfile := "/tmp/ensemble.log"
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprint("open", logfile, "failed:", err.Error()))
	}
	defer f.Close()

	log.SetOutput(f)

	c := daemon.NewClient()
	if err := c.Notify(name); err != nil {
		log.Fatalln(err)
	}
}

func serve() {
	if err := daemon.Serve(); err != nil {
		log.Fatalln("an error occurred while http serving:", err)
	}
}
