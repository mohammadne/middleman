package servers

import (
	"fmt"

	"github.com/mohammadne/middleman/internal/configs"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/internal/network/server"
	"github.com/mohammadne/middleman/internal/storage"
	"github.com/mohammadne/middleman/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	use   = "servers"
	short = "run servers"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	envFlag := "set config environment, default is dev"
	cmd.Flags().StringP("env", "e", "", envFlag)

	return cmd
}

func main(cmd *cobra.Command, _ []string) {
	env := cmd.Flag("env").Value.String()
	configs := configs.Server(env)

	lg := logger.NewZap(configs.Logger)

	stopChannel := make(chan interface{})

	for _, port := range configs.Ports {
		config := network.ServerConfig{Host: configs.Host, Port: port}
		storagePath := fmt.Sprintf("%s/%s", configs.StorageDirectory, port)
		storage, err := storage.NewFileStorage(storagePath, lg)
		if err != nil {
			lg.Fatal("error creating storage", logger.Error(err))
		}

		server := server.New(&config, storage, lg)
		go server.Serve()
	}

	<-stopChannel
}
