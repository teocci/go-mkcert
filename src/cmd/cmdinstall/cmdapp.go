// Package cmdinstall
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-27
package cmdinstall

const (
	Name  = "install"
	Short = "Install the local CA in the system trust store."
	Long  = `mkcert automatically creates and installs a local CA in the system root store, and generates locally-trusted certificates.`

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
