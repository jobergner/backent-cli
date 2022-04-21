// env provides utility methods to ensure the
// generating process does not start until the user's
// filesystem fulfills the requirements
package env

import (
	"os"
)

// EnsureDir ensures that a directory exists at `path`
// does nothing if dir already exists.
func EnsureDir(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {

		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}

	} else {
		return err
	}

	return nil
}
