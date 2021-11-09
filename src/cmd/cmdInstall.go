// Package cmd
// Created by RTT.
// Author: teocci@yandex.com on 2021-Nov-09
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/teocci/go-mkcert/src/cmd/cmdinstall"
)

func init() {
	app.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   cmdinstall.Name,
	Short: cmdinstall.Short,
	Long:  cmdinstall.Long,
	Args:  validate,
	RunE:  installE,
}

func installE(cmd *cobra.Command, args []string) error {
	fmt.Printf("%#v\n", cmd)

	return nil
}
