package conf_loader

import (
	"encoding/json"
	"github.com/go-yaml/yaml"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"os"
)

// Unmarshal parses data from environment variables and a file (if defined)
// and stores the result in the value pointed to by v
func Unmarshal(envPrefix string, f *os.File, v interface{}) (e error) {

	// find env variables and fill the interface with values
	e = envconfig.Process(envPrefix, v)

	// file validation
	if f != nil && e == nil {
		fContent, ferr := ioutil.ReadAll(f)
		if ferr != nil {
			return
		}

		// check the file type
		if json.Valid(fContent) {
			e = json.Unmarshal(fContent, v)
		} else {
			// TODO - find a better way to process a specific file type
			e = yaml.Unmarshal(fContent, v)
		}
	}

	return
}
