// Package netpol provides primitives for command 'roxctl analyze netpol'
package netpol

import (
	"os"
	"path/filepath"

	npguard "github.com/np-guard/netpol-analyzer/pkg/netpol/connlist"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/roxctl/common/npg"
)

const defaultOutputFileNamePrefix = "connlist."

type netpolAnalyzer interface {
	ConnlistFromDirPath(dirPath string) ([]npguard.Peer2PeerConnection, error)
	ConnectionsListToString(conns []npguard.Peer2PeerConnection) (string, error)
	Errors() []npguard.ConnlistError
}

func (cmd *analyzeNetpolCommand) analyzeNetpols(analyzer netpolAnalyzer) error {
	conns, err := analyzer.ConnlistFromDirPath(cmd.inputFolderPath)
	if err != nil {
		return errors.Wrap(err, "error in connectivity analysis")
	}
	connsStr, err := analyzer.ConnectionsListToString(conns)
	if err != nil {
		return errors.Wrap(err, "error in formating connectivity list")
	}
	if err := cmd.ouputConnList(connsStr); err != nil {
		return err
	}
	var roxerr error
	for _, e := range analyzer.Errors() {
		if e.IsSevere() {
			cmd.env.Logger().ErrfLn("%s %s", e.Error(), e.Location())
			roxerr = npg.ErrErrors
		} else {
			cmd.env.Logger().WarnfLn("%s %s", e.Error(), e.Location())
			if cmd.treatWarningsAsErrors && roxerr == nil {
				roxerr = npg.ErrWarnings
			}
		}
	}
	return roxerr
}

func (cmd *analyzeNetpolCommand) ouputConnList(connsStr string) error {
	if cmd.outputToFile {
		dirpath, filename := filepath.Split(cmd.outputFilePath)
		if filename == "" {
			fileSuffix := cmd.outputFormat
			if cmd.outputFormat == "" {
				fileSuffix = "txt"
			}
			filename = defaultOutputFileNamePrefix + fileSuffix
		}

		if err := writeFile(filename, dirpath, connsStr); err != nil {
			return errors.Wrap(err, "error writing connlist output")
		}
	}

	cmd.printConnList(connsStr)
	return nil
}

func writeFile(filename string, destDir string, content string) error {
	outputPath := filepath.Join(destDir, filename)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return errors.Wrapf(err, "error creating directory for file %q", filename)
	}
	return errors.Wrap(os.WriteFile(outputPath, []byte(content), os.FileMode(0644)), "error writing file")
}

func (cmd *analyzeNetpolCommand) printConnList(connlist string) {
	cmd.env.Logger().PrintfLn(connlist)
}
