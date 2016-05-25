package commands

import (
	"net/http"

	"github.com/pivotal-cf/cm-cli/actions"
	"github.com/pivotal-cf/cm-cli/config"
	. "github.com/pivotal-cf/cm-cli/errors"
	"github.com/pivotal-cf/cm-cli/repositories"
)

type SetCommand struct {
	SecretIdentifier string `short:"n" required:"yes" long:"name" description:"Selects the secret being set"`
	SecretContent    string `short:"s" long:"secret" description:"Sets a value for a secret name"`
	Generate         bool   `short:"g" long:"generate" description:"System will generate random credential. Cannot be used in combination with --secret."`
}

func (cmd SetCommand) Execute([]string) error {
	if !cmd.Generate && cmd.SecretContent == "" {
		return NewSetOptionMissingError()
	}

	secretRepository := repositories.NewSecretRepository(http.DefaultClient)

	action := actions.NewSet(secretRepository, config.ReadConfig())

	secret, err := action.SetSecret(cmd.SecretIdentifier, cmd.SecretContent)
	if err != nil {
		return err
	}

	secret.PrintSecret()

	return nil
}
