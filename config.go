package jsonconf

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

func Load(path string, v interface{}) (err error) {
	// Nothing to do if there's no path.
	if len(path) == 0 {
		panic("No path given.")
	}

	// Expand user directory if needed.
	if path[0] == '~' {
		usr, err := user.Current()
		if err != nil {
			return err
		}
		path = filepath.Join(usr.HomeDir, path[1:])
	}

	// Get the absolute path.
	if path, err = filepath.Abs(path); err != nil {
		return err
	}

	// Load the file contents.
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Unmarshal the data.
	return json.Unmarshal(buf, &v)
}
