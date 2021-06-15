package cmd

import (
	"github.com/jenkins-x-plugins/jx-updatebot/pkg/cmd/argo"
	"github.com/jenkins-x-plugins/jx-updatebot/pkg/cmd/environment"
	"github.com/jenkins-x-plugins/jx-updatebot/pkg/cmd/pipeline"
	"github.com/jenkins-x-plugins/jx-updatebot/pkg/cmd/pr"
	"github.com/jenkins-x-plugins/jx-updatebot/pkg/cmd/sync"
	"github.com/jenkins-x-plugins/jx-updatebot/pkg/cmd/version"
	"github.com/jenkins-x-plugins/jx-updatebot/pkg/rootcmd"
	"github.com/jenkins-x/jx-helpers/v3/pkg/cobras"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/spf13/cobra"
)

// Main creates the new command
func Main() *cobra.Command {
	cmd := &cobra.Command{
		Use:   rootcmd.TopLevelCommand,
		Short: "commands for creating Pull Requests on repositories when versions change",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}
	cmd.AddCommand(cobras.SplitCommand(argo.NewCmdArgoPromote()))
	cmd.AddCommand(cobras.SplitCommand(environment.NewCmdUpgradeEnvironment()))
	cmd.AddCommand(cobras.SplitCommand(pipeline.NewCmdUpgradePipeline()))
	cmd.AddCommand(cobras.SplitCommand(pr.NewCmdPullRequest()))
	cmd.AddCommand(cobras.SplitCommand(sync.NewCmdEnvironmentSync()))
	cmd.AddCommand(cobras.SplitCommand(version.NewCmdVersion()))
	return cmd
}
