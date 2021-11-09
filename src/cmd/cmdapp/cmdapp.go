// Package cmdapp
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-27
package cmdapp

const (
	Name  = "go-mkcert"
	Short = "Command mkcert is a simple zero-config tool to make development certificates."
	Long  = `Using certificates from real certificate authorities (CAs) for development can be dangerous or impossible (for hosts like example.test, localhost or 127.0.0.1), but self-signed certificates cause trust errors. Managing your own CA is the best solution, but usually involves arcane commands, specialized knowledge and manual steps. 

mkcert automatically creates and installs a local CA in the system root store, and generates locally-trusted certificates. mkcert does not automatically configure servers to use the certificates, though, that's up to you.`

	CRSName    = "csr"
	CRSShort   = ""
	CRSDesc    = "Generate a certificate based on the supplied CSR. Conflicts with all other flags and arguments except -install and -cert-file."
	CRSDefault = ""

	CFName    = "cert-file"
	CFShort   = ""
	CFDesc    = "Customize the output paths."
	CFDefault = ""

	KFName    = "key-file"
	KFShort   = ""
	KFDesc    = "Customize the output paths."
	KFDefault = ""

	PFName    = "p12-file"
	PFShort   = ""
	PFDesc    = "Customize the output paths."
	PFDefault = ""

	IName    = "install"
	IShort   = "i"
	IDesc    = "Install the local CA in the system trust store."
	IDefault = false

	UName    = "uninstall"
	UShort   = "u"
	UDesc    = "Uninstall the local CA (but do not delete it)."
	UDefault = false

	PName    = "pkcs12"
	PShort   = ""
	PDesc    = "Generate a \".p12\" PKCS #12 file, also know as a \".pfx\" file,\n\tcontaining certificate and key for legacy applications."
	PDefault = false

	EName    = "ecdsa"
	EShort   = ""
	EDesc    = "Generate a certificate with an ECDSA key."
	EDefault = false

	CName    = "client"
	CShort   = "c"
	CDesc    = "Generate a certificate for client authentication."
	CDefault = false

	CAName    = "ca-root"
	CAShort   = ""
	CADesc    = "Print the CA certificate and key storage location."
	CADefault = false
)

const (
	VersionTemplate = "%s %s.%s\n"
	Version         = "v1.0"
	Commit          = "0"
)
