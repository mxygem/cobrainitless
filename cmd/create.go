package cmd

import (
	"fmt"
	"log"

	"github.com/mxygem/cobrainitless/internal/accounts"
	"github.com/mxygem/cobrainitless/internal/accounts/users"
	"github.com/spf13/cobra"
)

type createCLI struct {
	accounts *accounts.Client
}

func create(cfg *accounterCLI) *cobra.Command {
	cli := &createCLI{accounts: accounts.NewClient(cfg.api)}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a new user",
		RunE:  cli.create,
	}

	cli.setupFlags(cmd)

	return cmd
}

func (c *createCLI) setupFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("type", "t", "user", "the resource type to create. options: [user]")
	cmd.Flags().StringP("email", "e", "", "specify an email to use")
}

func (c *createCLI) create(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	resourceType := flags.ShorthandLookup("t").Value.String()

	switch resourceType {
	case "user":
		email := flags.ShorthandLookup("e").Value.String()
		log.Printf("create command sending email to users.Create: %q\n", email)
		return users.Create(c.accounts, email)
	default:
		return fmt.Errorf("unknown resource type of %q received", resourceType)
	}
}
