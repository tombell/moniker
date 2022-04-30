package moniker

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/bogem/id3v2/v2"
)

func Run(dir, format string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("ioutil read dir failed: %w", err)
	}

	for _, file := range files {
		if path.Ext(file.Name()) != ".mp3" {
			continue
		}

		src := path.Join(trimNull(dir), file.Name())

		tags, err := id3v2.Open(src, id3v2.Options{Parse: true})
		if err != nil {
			continue
		}
		defer tags.Close()

		filename := format + ".mp3"

		formatters := map[string]string{
			"{album}":  trimNull(tags.Album()),
			"{artist}": trimNull(tags.Artist()),
			"{genre}":  trimNull(tags.Genre()),
			"{title}":  trimNull(tags.Title()),
			"{year}":   trimNull(tags.Year()),
		}

		for key, val := range formatters {
			filename = strings.ReplaceAll(filename, key, val)
		}

		filename = strings.ReplaceAll(filename, "/", "_")
		filename = strings.ReplaceAll(filename, "\\", "_")

		dest := trimNull(path.Join(dir, filename))

		if err := os.Rename(src, dest); err != nil {
			fmt.Fprintf(os.Stderr, "error renaming file: %v (%s)\n", src, err)
		}
	}

	return nil
}

func trimNull(str string) string {
	return strings.Trim(str, "\x00")
}
