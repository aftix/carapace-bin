package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setCredentialsCmd = &cobra.Command{
	Use:   "set-credentials NAME [--client-certificate=path/to/certfile] [--client-key=path/to/keyfile] [--token=bearer_token] [--username=basic_user] [--password=basic_password] [--auth-provider=provider_name] [--auth-provider-arg=key=value] [--exec-command=exec_command] [--exec-api-version=exec_api_version] [--exec-arg=arg] [--exec-env=key=value]",
	Short: "Set a user entry in kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCredentialsCmd).Standalone()

	config_setCredentialsCmd.Flags().String("auth-provider", "", "Auth provider for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().StringSlice("auth-provider-arg", []string{}, "'key=value' arguments for the auth provider")
	config_setCredentialsCmd.Flags().String("client-certificate", "", "Path to client-certificate file for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("client-key", "", "Path to client-key file for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().Int("embed-certs", 0, "Embed client cert/key for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("exec-api-version", "", "API version of the exec credential plugin for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().StringSlice("exec-arg", []string{}, "New arguments for the exec credential plugin command for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("exec-command", "", "Command for the exec credential plugin for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().StringArray("exec-env", []string{}, "'key=value' environment values for the exec credential plugin")
	config_setCredentialsCmd.Flags().String("password", "", "password for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("token", "", "token for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("username", "", "username for the user entry in kubeconfig")
	config_setCredentialsCmd.Flag("embed-certs").NoOptDefVal = "true"
	configCmd.AddCommand(config_setCredentialsCmd)

	carapace.Gen(config_setCredentialsCmd).FlagCompletion(carapace.ActionMap{
		"client-certificate": carapace.ActionFiles(),
		"client-key":         carapace.ActionFiles(),
	})
}
