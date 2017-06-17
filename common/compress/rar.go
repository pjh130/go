/*
代码大部分来自github.com/mholt/archiver
*/

package archiver

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/nwaples/rardecode"
)

// Rar makes a .rar archive, but this is not implemented because
// RAR is a proprietary format. It is here only for symmetry with
// the other archive formats in this package.
func Rar(rarPath string, filePaths []string) error {
	return fmt.Errorf("make %s: RAR not implemented (proprietary format)", rarPath)
}

// Unrar extracts the RAR file at source and puts the contents
// into destination.
func UnRar(source, destination string) error {
	f, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("%s: failed to open archive: %v", source, err)
	}
	defer f.Close()

	rr, err := rardecode.NewReader(f, "")
	if err != nil {
		return fmt.Errorf("%s: failed to create reader: %v", source, err)
	}

	for {
		header, err := rr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if header.IsDir {
			err = os.MkdirAll(filepath.Join(destination, header.Name), 0755)
			if err != nil {
				return err
			}
			continue
		}

		// if files come before their containing folders, then we must
		// create their folders before writing the file
		err = os.MkdirAll(filepath.Dir(filepath.Join(destination, header.Name)), 0755)
		if err != nil {
			return err
		}

		err = writeNewFile(filepath.Join(destination, header.Name), rr, header.Mode())
		if err != nil {
			return err
		}
	}

	return nil
}
