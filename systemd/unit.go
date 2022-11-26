package systemd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var systemdUnitFilesDirPath = ""

func init() {
	path, err := os.UserConfigDir()
	if err != nil {
		panic(fmt.Sprint("os.UserConfigDir() failed:", err.Error()))
	}

	SetSystemdUnitFileDirPath(filepath.Join(path, "systemd", "user"))

	log.Println("systemd unit file dir path:", systemdUnitFilesDirPath)
}

func SetSystemdUnitFileDirPath(path string) {
	systemdUnitFilesDirPath = path
}

type Unit struct {
	Name string
	Body string
}

func Start(units []Unit) error {
	for _, unit := range units {
		if err := createUnitFile([]byte(unit.Body), unit.FilePath()); err != nil {
			return err
		}
	}

	if err := run(systemctlReloadSaemonCommand()); err != nil {
		return err
	}

	env := Env{
		"PATH": os.Getenv("PATH"),
	}
	if err := run(systemctlSetEnviromentCommand(env)); err != nil {
		return err
	}

	for _, unit := range units {
		if err := run(systemctlRestartUnitCommand(unit.FileName())); err != nil {
			return err
		}
	}

	return nil
}

func (unit Unit) FileName() string {
	return "ensemble-" + unit.Name + ".service"
}

func (unit Unit) FilePath() string {
	return filepath.Join(systemdUnitFilesDirPath, unit.FileName())
}

func createUnitFile(body []byte, path string) error {
	log.Println("create unit file:", path)

	if _, err := os.Stat(path); err == nil {
		if err := os.Remove(path); err != nil {
			return err
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err := f.Write(body); err != nil {
		return err
	}

	return nil
}

func run(cmd *exec.Cmd) error {
	log.Println("exec command:", strings.Join(cmd.Args, " "))
	b, err := cmd.CombinedOutput()
	if len(b) > 0 {
		log.Println("=== command output ===\n", string(b))
		log.Println("======================")
	}
	if err != nil {
		return err
	}
	return nil
}

func systemctlReloadSaemonCommand() *exec.Cmd {
	return exec.Command("systemctl", "--user", "daemon-reload")
}

func systemctlRestartUnitCommand(unit string) *exec.Cmd {
	return exec.Command("systemctl", "--user", "restart", unit)
}

func systemctlSetEnviromentCommand(env Env) *exec.Cmd {
	return exec.Command(
		"systemctl",
		append(
			[]string{"--user", "set-environment"},
			env.Strings()...,
		)...,
	)
}
