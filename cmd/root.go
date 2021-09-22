package main

import (
	"github.com/mohammadne/middleman/cmd/clients"
	"github.com/mohammadne/middleman/cmd/proxies"
	"github.com/mohammadne/middleman/cmd/servers"
	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"

	use   = "middleman"
	short = "short middleman"
	long  = "long middleman"
)

func main() {
	cmd := &cobra.Command{Use: use, Short: short, Long: long}
	cmd.AddCommand(servers.Command(), proxies.Command(), clients.Command())

	if err := cmd.Execute(); err != nil {
		panic(map[string]string{"reason": errExecuteCMD, "error": err.Error()})
	}
}
