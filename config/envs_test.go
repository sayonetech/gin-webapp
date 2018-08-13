package config

import "testing"
import "os"

func TestReadenv(t *testing.T) {
	os.Setenv("test", "random")
	value := string(Read()["test"].(string))
	if value != "random" {
		t.Errorf("read value is different than set value %s", value)
	}

}
