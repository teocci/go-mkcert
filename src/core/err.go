// Package core
// Created by RTT.
// Author: teocci@yandex.com on 2021-Nov-09
package core

import "log"

func fatalIfErr(err error, msg string) {
	if err != nil {
		log.Fatalf("ERROR: %s: %s", msg, err)
	}
}

func fatalIfCmdErr(err error, cmd string, out []byte) {
	if err != nil {
		log.Fatalf("ERROR: failed to execute \"%s\": %s\n\n%s\n", cmd, err, out)
	}
}
