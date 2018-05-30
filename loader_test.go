package conf_loader

import (
	"os"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {

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

	envPrefix := "myapp"

	// set env variables
	os.Setenv("MYAPP_INT_VAL", "1")
	os.Setenv("MYAPP_STRING_ENV", "test1")
	os.Setenv("MYAPP_OPTIONS_FIRST", "1")
	os.Setenv("MYAPP_OPTIONS_SECOND", "test1")

	// open yaml file

	// open json file

	tests := []struct {
		name      string
		expected  *testConfig
		envPrefix string
		f         *os.File
	}{
		{
			name:      "env_without_file",
			envPrefix: envPrefix,
			expected: &testConfig{
				IntVal:    1,
				StringVal: "test1",
				Options: &testOptionsConfig{
					First:  "1",
					Second: "test1",
				},
			},
		},
		{
			name:      "env_with_jsonfile",
			envPrefix: envPrefix,
			f:         jsonFile(),
			expected: &testConfig{
				IntVal:    2,
				StringVal: "test1",
				Options: &testOptionsConfig{
					First:  "1",
					Second: "test2",
				},
			},
		},
		{
			name:      "env_with_yamlfile",
			envPrefix: envPrefix,
			f:         yamlFile(),
			expected: &testConfig{
				IntVal:    1,
				StringVal: "test3",
				Options: &testOptionsConfig{
					First:  "3",
					Second: "test1",
				},
			},
		},
		{
			name:      "jsonfile_without_env",
			envPrefix: "unknown",
			f:         jsonFile(),
			expected: &testConfig{
				IntVal: 2,
				Options: &testOptionsConfig{
					Second: "test2",
				},
			},
		},
		{
			name:      "yamlfile_without_env",
			envPrefix: "unknown",
			f:         yamlFile(),
			expected: &testConfig{
				StringVal: "test3",
				Options: &testOptionsConfig{
					First: "3",
				},
			},
		},
	}
	var e error
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := &testConfig{}
			e = Unmarshal(test.envPrefix, test.f, v)
			if e != nil {
				t.Errorf("Failed with error: %s", e.Error())
			}
			if test.f != nil {
				defer test.f.Close()
			}

			if !reflect.DeepEqual(test.expected, v) {
				t.Errorf("Got %s, exptected %s. Unproper pointer values.", v, test.expected)
			}
		})
	}
}

func jsonFile() *os.File {
	f, _ := os.Open("data.example.json")
	return f
}

func yamlFile() *os.File {
	f, _ := os.Open("data.example.yml")
	return f
}
