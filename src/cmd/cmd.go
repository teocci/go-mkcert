// Package cmd
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-27
package cmd

import (
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/url"
	"regexp"

	"golang.org/x/net/idna"

	"github.com/spf13/cobra"
	"github.com/teocci/go-mkcert/src/cmd/cmdapp"
	"github.com/teocci/go-mkcert/src/config"
	"github.com/teocci/go-mkcert/src/core"
	"github.com/teocci/go-mkcert/src/datamgr"
	"github.com/teocci/go-mkcert/src/logger"
)

var (
	app = &cobra.Command{
		Use:           cmdapp.Name,
		Short:         cmdapp.Short,
		Long:          cmdapp.Long,
		Args:          validate,
		RunE:          runE,
		SilenceErrors: false,
		SilenceUsage:  true,
	}

	crs      string
	certFile string
	keyFile  string
	p12File  string

	install   bool
	uninstall bool
	pkcs12    bool
	ecdsa     bool
	client    bool
	caRoot    bool
)

// Add supported cli commands/flags
func init() {
	cobra.OnInitialize(initConfig)

	app.Flags().StringVarP(&crs, cmdapp.CRSName, cmdapp.CRSShort, cmdapp.CRSDefault, cmdapp.CRSDesc)
	app.Flags().StringVarP(&certFile, cmdapp.CFName, cmdapp.CFShort, cmdapp.CFDefault, cmdapp.CFDesc)
	app.Flags().StringVarP(&keyFile, cmdapp.KFName, cmdapp.KFShort, cmdapp.KFDefault, cmdapp.KFDesc)
	app.Flags().StringVarP(&p12File, cmdapp.PFName, cmdapp.PFShort, cmdapp.PFDefault, cmdapp.PFDesc)

	app.Flags().BoolVarP(&install, cmdapp.IName, cmdapp.IShort, cmdapp.IDefault, cmdapp.IDesc)
	app.Flags().BoolVarP(&uninstall, cmdapp.UName, cmdapp.UShort, cmdapp.UDefault, cmdapp.UDesc)
	app.Flags().BoolVarP(&pkcs12, cmdapp.PName, cmdapp.PShort, cmdapp.PDefault, cmdapp.PDesc)
	app.Flags().BoolVarP(&ecdsa, cmdapp.EName, cmdapp.EShort, cmdapp.EDefault, cmdapp.EDesc)
	app.Flags().BoolVarP(&client, cmdapp.CName, cmdapp.CShort, cmdapp.CDefault, cmdapp.CDesc)
	app.Flags().BoolVarP(&caRoot, cmdapp.CAName, cmdapp.CAShort, cmdapp.CADefault, cmdapp.CADesc)

	//_ = app.MarkFlagRequired(cmdapp.CRSName)

	config.AddFlags(app)
}

// Load config
func initConfig() {
	if err := config.LoadConfigFile(); err != nil {
		log.Fatal(err)
	}

	config.LoadLogConfig()
}

func validate(cmd *cobra.Command, args []string) error {
	fmt.Printf("%#v\n", cmd)
	argsLength := len(args)

	if config.Version {
		fmt.Printf(cmdapp.VersionTemplate, cmdapp.Name, cmdapp.Version, cmdapp.Commit)

		return nil
	}

	if !config.Verbose {
		cmd.HelpFunc()(cmd, args)

		return fmt.Errorf("")
	}

	if cmd.Use == "install" && argsLength == 0 {
		return nil
	}

	if cmd.Use == "uninstall" && argsLength == 0 {
		return nil
	}

	if argsLength == 0 {
		cmd.HelpFunc()(cmd, args)
		return fmt.Errorf("no arguments")
	}

	hostnameRegexp := regexp.MustCompile(`(?i)^(\*\.)?[0-9a-z_-]([0-9a-z._-]*[0-9a-z_-])?$`)
	for i, name := range args {
		if ip := net.ParseIP(name); ip != nil {
			continue
		}
		if email, err := mail.ParseAddress(name); err == nil && email.Address == name {
			continue
		}
		if uriName, err := url.Parse(name); err == nil && uriName.Scheme != "" && uriName.Host != "" {
			continue
		}
		punycode, err := idna.ToASCII(name)
		if err != nil {
			return fmt.Errorf("ERROR: %q is not a valid hostname, IP, URL or email: %s", name, err)
		}
		args[i] = punycode
		if !hostnameRegexp.MatchString(punycode) {
			return fmt.Errorf("ERROR: %q is not a valid hostname, IP, URL or email", name)
		}
	}

	return nil
}

func runE(ccmd *cobra.Command, args []string) error {
	var err error
	config.Log, err = logger.New(config.LogConfig)
	if err != nil {
		return ErrCanNotLoadLogger(err)
	}

	initData := data.InfoConf{
		//Host:      crs,
		//Port:      install,
		//ConnID:    pkcs12,
		//ModuleTag: certFile,
		//DroneID:   ecdsa,
		//FlightID:  client,
	}
	// Make channel for errors
	errs := make(chan error)
	go func() {
		errs <- core.Start(initData)
	}()

	// Break if any of them return an error (blocks exit)
	if err := <-errs; err != nil {
		config.Log.Fatal(err)
	}

	return err
}

func Execute() {
	err := app.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
