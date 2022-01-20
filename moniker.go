package moniker

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/bogem/id3v2/v2"
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
			continue
		}

		src := path.Join(trimNull(dir), file)

		tags, err := id3v2.Open(src, id3v2.Options{Parse: true})
		defer tags.Close()
		if err != nil {
			continue
		}

		filename := generateFilename(format, tags)
		dest := trimNull(path.Join(dir, filename))

		if err := os.Rename(src, dest); err != nil {
			fmt.Fprintf(os.Stderr, "error renaming file: %v (%s)\n", src, err)
		}
	}

	return nil
}

func generateFilename(format string, tags *id3v2.Tag) string {
	filename := format + ".mp3"

	formatters := map[string]string{
		"{album}":  trimNull(tags.Album()),
		"{artist}": trimNull(tags.Artist()),
		"{genre}":  trimNull(tags.Genre()),
		"{title}":  trimNull(tags.Title()),
		"{year}":   trimNull(tags.Year()),
	}

	for key, val := range formatters {
		filename = strings.Replace(filename, key, val, -1)
	}

	filename = strings.Replace(filename, "/", "_", -1)
	filename = strings.Replace(filename, "\\", "_", -1)

	return filename
}

func trimNull(str string) string {
	return strings.Trim(str, string(0))
}
