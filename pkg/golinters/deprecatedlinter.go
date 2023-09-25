package golinters

import (
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	linter "github.com/wwbweibo/deprecatedlinter"
	"golang.org/x/tools/go/analysis"
)

const name = "deprecatedlinter"

func NewDeprecatedLinter() *goanalysis.Linter {
	a := linter.Analyzer
	return goanalysis.NewLinter(
		name,
		"helps to check the deprecated code based on version",
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeWholeProgram)
}
