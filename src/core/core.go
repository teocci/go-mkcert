// Package core
// Created by RTT.
// Author: teocci@yandex.com on 2021-Nov-09
package core

import (
	"log"
	"os"

	"github.com/teocci/go-mkcert/src/datamgr"
)

const (
	rootName    = "rootCA.pem"
	rootKeyName = "rootCA-key.pem"
)

func Start(d data.InfoConf) error {
	var m = &mkcert{
		installMode:   d.InstallMode,
		uninstallMode: d.UninstallMode,
		csrPath:       d.CsrPath,
		pkcs12:        d.Pkcs12,
		ecdsa:         d.Ecdsa,
		client:        d.Client,
		certFile:      d.CertFile,
		keyFile:       d.KeyFile,
		p12File:       d.P12File,
	}

	m.CAROOT = envCAROOT()
	if m.CAROOT == "" {
		log.Fatalln("ERROR: failed to find the default CA location, set one as the CAROOT env var")
	}
	fatalIfErr(os.MkdirAll(m.CAROOT, 0755), "failed to create the CAROOT")
	m.loadCA()

	if m.installMode {
		m.install()
		return nil
	} else if m.uninstallMode {
		m.uninstall()
		return nil
	} else {
		var warning bool
		if storeEnabled("system") && !m.checkPlatform() {
			warning = true
			log.Println("Note: the local CA is not installed in the system trust store.")
		}
		if storeEnabled("nss") && hasNSS && CertutilInstallHelp != "" && !m.checkNSS() {
			warning = true
			log.Printf("Note: the local CA is not installed in the %s trust store.", NSSBrowsers)
		}
		if storeEnabled("java") && hasJava && !m.checkJava() {
			warning = true
			log.Println("Note: the local CA is not installed in the Java trust store.")
		}
		if warning {
			log.Println("Run \"mkcert -install\" for certificates to be trusted automatically.")
		}
	}

	if m.csrPath != "" {
		m.makeCertFromCSR()
		return nil
	}

	m.makeCert(d.Args)

	return nil
}
