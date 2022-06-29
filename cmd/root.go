package cmd

import (
	"github.com/kpurdon/apir/pkg/discoverer"
	"github.com/kpurdon/apir/pkg/requester"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type accounterCLI struct {
	api *requester.Client
}

func AccounterCmd() (*cobra.Command, error) {
	cli := &accounterCLI{}

	cmd := &cobra.Command{
		Use:     "accounter",
		Short:   "a tool to manage various aspects and behaviors relating to accounts",
		PreRunE: cli.setupConfig,
	}

	cli.setupFlags(cmd)
	cli.addCommands(cmd)

	return cmd, nil
}

func (c *accounterCLI) setupFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP("debug", "d", false, "enable debugging")
}

func (c *accounterCLI) setupConfig(cmd *cobra.Command, args []string) error {
	apiClient := requester.NewClient("admin",
		requester.WithDatadog(),
		requester.WithRetry(),
	)
	err := apiClient.AddAPI("accounts_api", discoverer.NewDirect(viper.GetString("targets.accounts_api")))
	if err != nil {
		return err
	}
	// and the services
	// ...
	c.api = apiClient

	return nil
}

func (c *accounterCLI) addCommands(cmd *cobra.Command) {
	subCmds := []func(cfg *accounterCLI) *cobra.Command{create} // get, update, delete}
	for _, subCmd := range subCmds {
		cmd.AddCommand(subCmd(c))
	}
}
