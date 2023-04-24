package marshallers

import (
	"os"
	"path/filepath"
)

var marshallerImport = []byte(`package marshallerimport 
import (   
	_ "github.com/mailru/easyjson/gen"
	_ "github.com/mailru/easyjson/jlexer"
	_ "github.com/mailru/easyjson/jwriter"
)`)

func WriteImportFile(path string) error {
	filePath := filepath.Join(path, "marshaller_import.go")

	if err := os.WriteFile(filePath, marshallerImport, 0644); err != nil {
		return err
	}

	return nil
}
