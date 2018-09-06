package moniker

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/bogem/id3v2"
)

// Run renames the MP3s in the given directory according to the given format.
func Run(dir, format string) error {
	if !exists(dir) {
		// TODO: nicer error messages
		return fmt.Errorf("directory (%s) does not exist", dir)
	}

	files, err := readFiles(dir)
	if err != nil {
		// TODO: nicer error messages
		return fmt.Errorf("unable to read files in directory (%s): %s", err)
	}

	if err := renameFiles(dir, format, files); err != nil {
		// TODO: nicer error messages
		return fmt.Errorf("error while renaming files: %s", err)
	}

	return nil
}

func exists(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func readFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0)

	for _, file := range files {
		names = append(names, file.Name())
	}

	return names, nil
}

func renameFiles(dir, format string, files []string) error {
	for _, file := range files {
		if path.Ext(file) != ".mp3" {
			continue
		}

		src := path.Join(dir, file)

		tags, err := id3v2.Open(src, id3v2.Options{Parse: true})
		if err != nil {
			return err
		}
		defer tags.Close()

		formatters := map[string]string{
			"{artist}": tags.Artist(),
			"{title}":  tags.Title(),
			"{album}":  tags.Album(),
			// TODO: add more formatters in the future...
		}

		filename := format + ".mp3"

		for key, val := range formatters {
			filename = strings.Replace(filename, key, val, -1)
		}

		dest := path.Join(dir, filename)

		if err := os.Rename(src, dest); err != nil {
			// TODO: nicer error messages
			return err
		}
	}

	return nil
}
