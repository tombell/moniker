package moniker

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/bogem/id3v2"
)

// Run renames the MP3s in the given directory according to the given format
// based on ID3 tags.
func Run(dir, format string) error {
	if !exists(dir) {
		return fmt.Errorf("error directory does not exist: %s", dir)
	}

	files, err := readFiles(dir)
	if err != nil {
		return fmt.Errorf("error reading files in directory: %s", err)
	}

	if err := renameFiles(dir, format, files); err != nil {
		return fmt.Errorf("error renaming files: %s", err)
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
			// Skipping non-MP3 files...
			continue
		}

		src := path.Join(dir, file)

		tags, err := id3v2.Open(src, id3v2.Options{Parse: true})
		if err != nil {
			// Skipping when failing to read ID3 tags...
			continue
		}
		defer tags.Close()

		filename := format + ".mp3"

		formatters := map[string]string{
			"{artist}": trimNull(tags.Artist()),
			"{title}":  trimNull(tags.Title()),
			"{album}":  trimNull(tags.Album()),
			"{genre}":  trimNull(tags.Genre()),
		}

		for key, val := range formatters {
			filename = strings.Replace(filename, key, val, -1)
		}

		filename = strings.Replace(filename, "/", "_", -1)
		filename = strings.Replace(filename, "\\", "_", -1)

		dest := path.Join(dir, filename)

		src = trimNull(src)
		dest = trimNull(dest)

		if err := os.Rename(src, dest); err != nil {
			fmt.Printf("failed to rename file: %v (%v)\n", src, err)
		}
	}

	return nil
}

func trimNull(str string) string {
	return strings.Trim(str, string(0))
}
