package basicmatter

import (
	"testing"
)

func TestConfig_ParseFile(t *testing.T) {
	config := New()

	err := config.parseFile("./test.conf")
	if err != nil {
		t.Fatalf("some error occurred: %s", err)
	}
}

type confObjectTest struct {
	Port string `config:"base:port"`
	Result bool `config:"base:result"`
	Int64 int64 `config:"base:int64.test"`
	Int8 int8 `config:"base:int8.test"`
	Int16 int16 `config:"base:int16.test"`
	SpecialV1 string `config:"special:special1"`
}

func TestConfig_Parse(t *testing.T) {
	config := New()
	obj := confObjectTest{Port: ":8081"}

	err := config.Unmarshal( &obj, "")
	if err != nil {
		t.Fatalf("some error occurred: %s", err)
	}

	t.Logf("obj: %v", obj)
}