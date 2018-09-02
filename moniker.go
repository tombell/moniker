package moniker

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

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

	for _, file := range files {
		if path.Ext(file) != ".mp3" {
			continue
		}

		if err := readTags(path.Join(dir, file)); err != nil {
			return err
		}
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

func readTags(file string) error {
	tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
	if err != nil {
		return err
	}
	defer tag.Close()

	// bpm := tag.CommonID("BPM")
	// bpmFrame := tag.GetTextFrame(bpm)

	// fmt.Println(tag.Artist())  // {artist}
	// fmt.Println(tag.Title())   // {title}
	// fmt.Println(tag.Album())   // {album}
	// fmt.Println(tag.Genre())   // {genre}
	// fmt.Println(tag.Year())    // {year}
	// fmt.Println(bpmFrame.Text) // {bpm}
	// fmt.Println("---")

	return nil
}

func rename(src, dst string) error {
	return os.Rename(src, dst)
}
