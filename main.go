package main

import (
	"log"

	"github.com/MakeNowJust/heredoc"
	"github.com/w-haibara/ensemble/unit"
)

func main() {
	units := []unit.Unit{
		{
			Name: "sample1",
			Body: heredoc.Doc(`
			[Unit]
			Description=Sample

			[Service]
			Type=exec
			ExecStart=/bin/echo AAA

			[Install]
			WantedBy=default.target
			`),
		},
	}
	if err := unit.Start(units); err != nil {
		log.Fatalln(err)
	}
}
