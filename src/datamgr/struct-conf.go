// Package data
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-30
package data

import (
	"crypto"
	"crypto/x509"
)

type InfoConf struct {
	InstallMode, UninstallMode bool
	Pkcs12, Ecdsa, Client      bool
	KeyFile, CertFile, P12File string
	CsrPath                    string

	CAROOT string
	CaCert *x509.Certificate
	CaKey  crypto.PrivateKey

	Args []string

	IgnoreCheckFailure bool
}
