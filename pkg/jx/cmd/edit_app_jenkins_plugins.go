package cmd

import (
	"github.com/jenkins-x/jx/pkg/auth"
	"github.com/jenkins-x/jx/pkg/jenkins"
	"github.com/jenkins-x/jx/pkg/jx/cmd/opts"
	"github.com/jenkins-x/jx/pkg/jx/cmd/templates"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	editAppJenkinsPluginsLong = templates.LongDesc(`
		Edits a Jenkins App's plugins
`)

	editAppJenkinsPluginsExample = templates.Examples(`
		# Edits the plugins for a Jenkins App
		jx edit app jenkins plugins
	`)
)

// EditAppJenkinsPluginsOptions the options for the create spring command
type EditAppJenkinsPluginsOptions struct {
	EditOptions

	Name    string
	Enabled string

	IssuesAuthConfigSvc auth.ConfigService
}

// NewCmdEditAppJenkinsPlugins creates a command object for the "create" command
func NewCmdEditAppJenkinsPlugins(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &EditAppJenkinsPluginsOptions{
		EditOptions: EditOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "app jenkins plugins",
		Short:   "Edits the Jenkins Plugins for a Jenkins App",
		Long:    editAppJenkinsPluginsLong,
		Example: editAppJenkinsPluginsExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			CheckErr(err)
		},
	}
	cmd.Flags().StringVarP(&options.Enabled, optionEnabled, "e", "", "Enables or disables the addon")

	return cmd
}

// Run implements the command
func (o *EditAppJenkinsPluginsOptions) Run() error {
	data, err := jenkins.LoadUpdateCenterURL(jenkins.DefaultUpdateCenterURL)
	if err != nil {
		return errors.Wrapf(err, "failed to load URL %s", jenkins.DefaultUpdateCenterURL)
	}

	// TODO load from the GitOps values.yaml folder
	currentValues := []string{"jx-resources:1.0.0"}

	selection, err := data.PickPlugins(currentValues, o.In, o.Out, o.Err)
	if err != nil {
		return err
	}

	log.Infof("chosen selection:\n")
	for _, sel := range selection {
		log.Infof("    %s\n", util.ColorInfo(sel))
	}
	// TODO update the GitOps values.yaml folder
	return nil
}
