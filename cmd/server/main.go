package server

import (
	"github.com/mohammadne/middleman/internal/configs"
	"github.com/mohammadne/middleman/internal/network/server"
	"github.com/mohammadne/middleman/internal/storage"
	"github.com/mohammadne/middleman/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	envFlag := "set config environment, default is dev"
	cmd.Flags().StringP("env", "e", "", envFlag)

	return cmd
}

func main(cmd *cobra.Command, _ []string) {
	env := cmd.Flag("env").Value.String()
	config := configs.Server(env)

	lg := logger.NewZap(config.Logger)

	stopChannel := make(chan interface{})

	for _, serverCfg := range config.Servers {
		storage, err := storage.NewFileStorage("", lg)
		if err != nil {
			lg.Fatal("error creating storage", logger.Error(err))
		}

		server := server.New(&serverCfg, storage)
		go server.Serve()
	}

	<-stopChannel
}
