// Package core
// Created by RTT.
// Author: teocci@yandex.com on 2021-Nov-09
package core

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"sync"
)

var sudoWarningOnce sync.Once

func commandWithSudo(cmd ...string) *exec.Cmd {
	if u, err := user.Current(); err == nil && u.Uid == "0" {
		return exec.Command(cmd[0], cmd[1:]...)
	}
	if !binaryExists("sudo") {
		sudoWarningOnce.Do(func() {
			log.Println(`Warning: "sudo" is not available, and mkcert is not running as root. The (un)install operation might fail. ⚠️`)
		})
		return exec.Command(cmd[0], cmd[1:]...)
	}
	return exec.Command("sudo", append([]string{"--prompt=Sudo password:", "--"}, cmd...)...)
}

func binaryExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
