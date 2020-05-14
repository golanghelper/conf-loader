# conf-loader
Unmarshal values from env, json and yaml.

## Usage example


ENV variables:

```bash
MYAPP_INT_VAL=1
MYAPP_STRING_ENV=test1
MYAPP_OPTIONS_FIRST=1
MYAPP_OPTIONS_SECOND=test1
```

JSON file overwrites two values:

```json
{
  "int_val": 2,
  "options": {
    "second": "test2"
  }
}
```

Go app:

```go
package main
import (
	"fmt"
	"github.com/golanghelper/conf-loader"
	"log"
	"os"
)
// prepare test structures
type testOptionsConfig struct {
	First  string
	Second string
}
type testConfig struct {
	IntVal    int                `json:"int_val" envconfig:"int_val" yaml:"int_val"`
	StringVal string             `json:"string_json" envconfig:"string_env" yaml:"string_yaml"`
	Options   *testOptionsConfig `json:"options" envconfig:"options" yaml:"options"`
}
func main() {
	f, _ := os.Open("data.example.json")
	var v testConfig
	e := conf_loader.Unmarshal("myapp", f, &v)
	if e != nil {
		log.Fatalf("Load config error: %s", e.Error())
	}
	fmt.Print(v)
}
```

Output:

```
{2 test1 0xc042002620}
```
