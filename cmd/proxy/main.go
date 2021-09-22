package proxy

import (
	"github.com/mohammadne/middleman/internal/configs"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/internal/network/proxy"
	"github.com/mohammadne/middleman/internal/storage"
	"github.com/mohammadne/middleman/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	use   = "proxy"
	short = "run proxy server"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	envFlag := "set config environment, default is dev"
	cmd.Flags().StringP("env", "e", "", envFlag)

	return cmd
}

func main(cmd *cobra.Command, _ []string) {
	env := cmd.Flag("env").Value.String()
	configs := configs.Proxy(env)

	lg := logger.NewZap(configs.Logger)

	storage := storage.NewMemoryStorage(lg)

	serverConfigs := make([]network.ServerConfig, 0, len(configs.ServerPorts))
	for _, port := range configs.ServerPorts {
		config := network.ServerConfig{Host: configs.ServerHost, Port: port}
		serverConfigs = append(serverConfigs, config)
	}

	server := proxy.New(configs.Proxy, serverConfigs, storage, lg)
	server.Serve()
}
