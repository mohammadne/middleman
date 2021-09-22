package main

import (
	"github.com/mohammadne/middleman/cmd/clients"
	"github.com/mohammadne/middleman/cmd/proxy"
	"github.com/mohammadne/middleman/cmd/servers"
	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"

	use   = "middleman"
	short = "middleman short description"
	long  = "middleman long description"
)

func main() {
	cmd := &cobra.Command{Use: use, Short: short, Long: long}
	cmd.AddCommand(servers.Command(), proxy.Command(), clients.Command())

	if err := cmd.Execute(); err != nil {
		panic(map[string]string{"reason": errExecuteCMD, "error": err.Error()})
	}
}
