package clients

import (
	"github.com/mohammadne/middleman/internal/configs"
	"github.com/mohammadne/middleman/internal/network/client"
	"github.com/mohammadne/middleman/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	use   = "clients"
	short = "run clients generator to mock servers with data"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	envFlag := "set config environment, default is dev"
	cmd.Flags().StringP("env", "e", "", envFlag)

	previewFlag := "if set to true, it will only preview changes and doesn't execute them"
	cmd.Flags().BoolP("preview", "p", false, previewFlag)

	return cmd
}

func main(cmd *cobra.Command, args []string) {
	env := cmd.Flag("env").Value.String()
	configs := configs.Client(env)

	lg := logger.NewZap(configs.Logger)

	client := client.New(configs.Client, nil, lg)
	client.Run()
}
