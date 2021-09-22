package proxy

import "github.com/spf13/cobra"

const (
	use   = "proxy"
	short = "run proxy"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	envFlag := "set config environment, default is dev"
	cmd.Flags().StringP("env", "e", "", envFlag)

	return cmd
}

func main(cmd *cobra.Command, args []string) {}
