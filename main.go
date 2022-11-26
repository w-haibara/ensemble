package main

import (
	"log"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/w-haibara/ensemble/daemon"
	"github.com/w-haibara/ensemble/systemd"
)

func main() {
	switch os.Args[1] {
	case "deploy":
		units := []systemd.Unit{
			{
				Name: "sample1",
				Body: heredoc.Doc(`
				[Unit]
				Description=Sample
	
				[Service]
				Type=exec
				ExecStart=/bin/bash -c "ensemble notify task-name"
	
				[Install]
				WantedBy=default.target
				`),
			},
		}
		if err := systemd.Start(units); err != nil {
			log.Fatalln(err)
		}
	case "notify":
		c := daemon.NewClient()
		if err := c.Notify(os.Args[2]); err != nil {
			log.Fatalln(err)
		}
	case "serve":
		if err := daemon.Serve(); err != nil {
			log.Fatalln("an error occurred while http serving:", err)
		}
	}
}
