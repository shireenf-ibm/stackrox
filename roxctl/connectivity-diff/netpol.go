package connectivitydiff

import (
	"os"

	npguard "github.com/np-guard/netpol-analyzer/pkg/netpol/diff"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stackrox/rox/pkg/errox"
	"github.com/stackrox/rox/roxctl/common/environment"
	"github.com/stackrox/rox/roxctl/common/npg"
)

type diffNetpolCommand struct {
	// Properties that are bound to cobra flags.
	stopOnFirstError      bool
	treatWarningsAsErrors bool
	inputFolderPath1      string
	inputFolderPath2      string
	outputFilePath        string
	removeOutputPath      bool
	outputToFile          bool
	outputFormat          string

	// injected or constructed values
	env environment.Environment
}

// Command defines the connectivity-diff command tree
func Command(cliEnvironment environment.Environment) *cobra.Command {
	diffNetpolCmd := &diffNetpolCommand{env: cliEnvironment}
	c := &cobra.Command{
		Use:   "connectivity-diff --dir1=<folder-path1> --dir2=<folder-path2>",
		Short: "(Technology Preview) Report connectivity diff based on two directories of network policies and workload resources YAML manifests.",
		Long: `Based on given two folders containing Kubernetes workloads and network policy YAMLs, will report all differences in allowed connections between the directories.
		Will write to stdout if no output flags are provided.

** This is a Technology Preview feature **
Technology Preview features are not supported with Red Hat production service level agreements (SLAs) and might not be functionally complete.
Red Hat does not recommend using them in production.
These features provide early access to upcoming product features, enabling customers to test functionality and provide feedback during the development process.
For more information about the support scope of Red Hat Technology Preview features, see https://access.redhat.com/support/offerings/techpreview/`,

		Args: cobra.ExactArgs(0),
		RunE: func(c *cobra.Command, args []string) error {
			diffNetpolCmd.env.Logger().WarnfLn("This is a Technology Preview feature. Red Hat does not recommend using Technology Preview features in production.")
			analyzer, err := diffNetpolCmd.construct(args)
			if err != nil {
				return err
			}
			if err := diffNetpolCmd.validate(); err != nil {
				return err
			}
			return diffNetpolCmd.analyzeConnectivityDiff(analyzer)
		},
	}

	c.Flags().StringVarP(&diffNetpolCmd.inputFolderPath1, "dir1", "", "", "first resources dir path")
	c.Flags().StringVarP(&diffNetpolCmd.inputFolderPath2, "dir2", "", "", "second resources dir path to be compared with the first dir path")
	c.Flags().BoolVar(&diffNetpolCmd.treatWarningsAsErrors, "strict", false, "treat warnings as errors")
	c.Flags().BoolVar(&diffNetpolCmd.stopOnFirstError, "fail", false, "fail on the first encountered error")
	c.Flags().BoolVar(&diffNetpolCmd.removeOutputPath, "remove", false, "remove the output path if it already exists")
	c.Flags().BoolVar(&diffNetpolCmd.outputToFile, "save-to-file", false, "whether to save connections diff output into default file")
	c.Flags().StringVarP(&diffNetpolCmd.outputFilePath, "output-file", "f", "", "save connections diff output into specific file")
	c.Flags().StringVarP(&diffNetpolCmd.outputFormat, "output-format", "o", defaultOutputFormat, "configure the connections diff in specific format, supported formats: txt|md|csv")
	return c
}

func (cmd *diffNetpolCommand) construct(args []string) (diffAnalyzer, error) {
	var opts []npguard.DiffAnalyzerOption
	if cmd.env != nil && cmd.env.Logger() != nil {
		opts = append(opts, npguard.WithLogger(npg.NewLogger(cmd.env.Logger())))
	}
	if cmd.stopOnFirstError {
		opts = append(opts, npguard.WithStopOnError())
	}
	if cmd.outputFormat != "" {
		opts = append(opts, npguard.WithOutputFormat(cmd.outputFormat))
	}
	if cmd.outputFilePath != "" {
		cmd.outputToFile = true
	}
	return npguard.NewDiffAnalyzer(opts...), nil
}

func (cmd *diffNetpolCommand) validate() error {
	if cmd.inputFolderPath1 == "" {
		return errors.New("directory path dir1 is required")
	}
	if cmd.inputFolderPath2 == "" {
		return errors.New("directory path dir2 is required")
	}
	if err := cmd.setupPath(cmd.outputFilePath); err != nil {
		return errors.Wrap(err, "failed to set up file path")
	}
	return nil
}

func (cmd *diffNetpolCommand) setupPath(path string) error {
	if _, err := os.Stat(path); err == nil && !cmd.removeOutputPath {
		return errox.AlreadyExists.Newf("path %s already exists. Use --remove to overwrite or select a different path.", path)
	} else if !os.IsNotExist(err) {
		return errors.Wrapf(err, "failed to check if path %s exists", path)
	}
	return nil
}
