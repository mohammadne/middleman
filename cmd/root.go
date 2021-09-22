package main

import (
	"github.com/mohammadne/middleman/cmd-new/client"
	"github.com/mohammadne/middleman/cmd-new/proxy"
	"github.com/mohammadne/middleman/cmd-new/server"
	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"

	use   = "library"
	short = "short library"
	long  = "long library"
)

func main() {
	cmd := &cobra.Command{Use: use, Short: short, Long: long}
	cmd.AddCommand(server.Command(), proxy.Command(), client.Command())

	if err := cmd.Execute(); err != nil {
		panic(map[string]string{"reason": errExecuteCMD, "error": err.Error()})
	}
}
