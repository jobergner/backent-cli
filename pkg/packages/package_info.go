// packages provides basic information of all packages that get generated.
package packages

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/factory/client"
	"github.com/jobergner/backent-cli/pkg/factory/message"
	"github.com/jobergner/backent-cli/pkg/factory/server"
	"github.com/jobergner/backent-cli/pkg/factory/state"
)

type Factory interface {
	Write() string
}

type PackageInfo struct {
	Name string
	// SourcePath is the relative path to the example of this package
	SourcePath string
	// StaticCodeIdentifier is the map key to the related StaticCode
	StaticCodeIdentifier string
	// DynamicCodeFactory writes all code which cannot be statically sourced fom examples
	DynamicCodeFactory Factory
}

// Packages lists package information. Order is relevant,
// as all dependencies need to exist before the depending package
// is generated
func Packages(ast *ast.AST) []PackageInfo {
	return []PackageInfo{
		{
			SourcePath:           "./examples/state",
			Name:                 "state",
			StaticCodeIdentifier: "importedCode_state",
			DynamicCodeFactory:   state.NewFactory(ast),
		},
		{
			SourcePath:           "./examples/connect",
			Name:                 "connect",
			StaticCodeIdentifier: "importedCode_connect",
			DynamicCodeFactory:   nil,
		},
		{
			SourcePath:           "./examples/message",
			Name:                 "message",
			StaticCodeIdentifier: "importedCode_message",
			DynamicCodeFactory:   message.NewFactory(ast),
		},
		{
			SourcePath:           "./examples/logging",
			Name:                 "logging",
			StaticCodeIdentifier: "importedCode_logging",
			DynamicCodeFactory:   nil,
		},
		{
			SourcePath:           "./examples/server",
			Name:                 "server",
			StaticCodeIdentifier: "importedCode_server",
			DynamicCodeFactory:   server.NewFactory(ast),
		},
		{
			SourcePath:           "./examples/client",
			Name:                 "client",
			StaticCodeIdentifier: "importedCode_client",
			DynamicCodeFactory:   client.NewFactory(ast),
		},
	}
}
