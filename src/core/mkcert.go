// Package core
// Created by RTT.
// Author: teocci@yandex.com on 2021-Nov-09
package core

import (
	"crypto"
	"crypto/x509"
)

type mkcert struct {
	installMode, uninstallMode bool
	pkcs12, ecdsa, client      bool
	keyFile, certFile, p12File string
	csrPath                    string

	CAROOT string
	caCert *x509.Certificate
	caKey  crypto.PrivateKey

	ignoreCheckFailure bool
}
