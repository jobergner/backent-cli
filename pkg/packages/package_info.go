// packages provides basic information of all packages that get generated.
package packages

import (
	"fmt"
	"path/filepath"

	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/factory/client"
	"github.com/jobergner/backent-cli/pkg/factory/message"
	"github.com/jobergner/backent-cli/pkg/factory/server"
	"github.com/jobergner/backent-cli/pkg/factory/state"
	"github.com/jobergner/backent-cli/pkg/factory/webclient"
)

type Language int

const (
	LangGo Language = iota + 1
	LangTS
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
	// FileName is the name of the single file (besides marshallers) of the generated package
	FileName string
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
			FileName:             "state.go",
		},
		{
			SourcePath:           "./examples/connect",
			Name:                 "connect",
			StaticCodeIdentifier: "importedCode_connect",
			DynamicCodeFactory:   nil,
			FileName:             "connect.go",
		},
		{
			SourcePath:           "./examples/message",
			Name:                 "message",
			StaticCodeIdentifier: "importedCode_message",
			DynamicCodeFactory:   message.NewFactory(ast),
			FileName:             "message.go",
		},
		{
			SourcePath:           "./examples/logging",
			Name:                 "logging",
			StaticCodeIdentifier: "importedCode_logging",
			DynamicCodeFactory:   nil,
			FileName:             "logging.go",
		},
		{
			SourcePath:           "./examples/server",
			Name:                 "server",
			StaticCodeIdentifier: "importedCode_server",
			DynamicCodeFactory:   server.NewFactory(ast),
			FileName:             "server.go",
		},
		{
			// NOTE: does not havy any code to source from
			SourcePath: "./examples/webclient",
			Name:       "webclient",
			// NOTE: tell process to not use any static code
			StaticCodeIdentifier: "",
			DynamicCodeFactory:   webclient.NewFactory(ast),
			FileName:             "index.ts",
		},
		{
			SourcePath:           "./examples/client",
			Name:                 "client",
			StaticCodeIdentifier: "importedCode_client",
			DynamicCodeFactory:   client.NewFactory(ast),
			FileName:             "client.go",
		},
	}
}

func (p PackageInfo) Lang() Language {
	ext := filepath.Ext(p.FileName)

	switch ext {
	case ".go":
		return LangGo
	case ".ts":
		return LangTS
	}

	panic(fmt.Errorf("unknown extension %s", ext))
}

func (p PackageInfo) Paths(baseDirPath string) (string, string) {
	dirPath := filepath.Join(baseDirPath, p.Name)
	filePath := filepath.Join(dirPath, p.FileName)
	return dirPath, filePath
}
