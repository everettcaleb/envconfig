package envconfig

import (
	"fmt"
	"testing"
)

type customIntType int

type testStruct struct {
	BoolField      bool          `env:"BOOL_FIELD"`
	Float32Field   float32       `env:"FLOAT_FIELD"`
	Float64Field   float64       `env:"FLOAT_FIELD"`
	IntField       int           `env:"INT_FIELD"`
	Int8Field      int8          `env:"INT_FIELD"`
	Int16Field     int16         `env:"INT_FIELD"`
	Int32Field     int32         `env:"INT_FIELD"`
	Int64Field     int64         `env:"INT_FIELD"`
	StringField    string        `env:"STRING_FIELD"`
	UintField      uint          `env:"INT_FIELD"`
	Uint8Field     uint8         `env:"INT_FIELD"`
	Uint16Field    uint16        `env:"INT_FIELD"`
	Uint32Field    uint32        `env:"INT_FIELD"`
	Uint64Field    uint64        `env:"INT_FIELD"`
	CustomIntField customIntType `env:"INT_FIELD"`
	PointerField   *int          `env:"INT_FIELD"`

	NotExistField         string `env:"NOT_EXISTS"`
	RequiredField         string `env:"STRING_FIELD" required:"true"`
	UntaggedField         string
	taggedUnexportedField string `env:"NOT_EXISTS" required:"true"`

	StructField struct {
		SubBoolField       bool `env:"BOOL_FIELD" required:"true"`
		subUnexportedField bool `env:"NOT_EXISTS" required:"true"`
	}
	StringSliceField []string `env:"STRING_SLICE_FIELD"`
}

type testFailStruct struct {
	RequiredNotExistField string `env:"NOT_EXISTS" required:"true"`
}

var mockEnv = map[string]string{
	"BOOL_FIELD":         "yes",
	"FLOAT_FIELD":        "3.14",
	"INT_FIELD":          "42",
	"STRING_FIELD":       "hello, world!",
	"STRING_SLICE_FIELD": "hi:hello:hola",
}

func getMockEnv(name string) string {
	return mockEnv[name]
}

func TestUnmarshal(t *testing.T) {
	env = getMockEnv
	var i int
	s := testStruct{
		PointerField: &i,
	}

	err := Unmarshal(&s)
	if err != nil {
		t.Error(err)
	}
	if !s.BoolField {
		t.Fail()
	}
	if s.Float32Field != 3.14 {
		t.Fail()
	}
	if s.Float64Field != 3.14 {
		t.Fail()
	}
	if s.IntField != 42 {
		t.Fail()
	}
	if s.Int8Field != 42 {
		t.Fail()
	}
	if s.Int16Field != 42 {
		t.Fail()
	}
	if s.Int32Field != 42 {
		t.Fail()
	}
	if s.Int64Field != 42 {
		t.Fail()
	}
	if s.StringField != "hello, world!" {
		t.Fail()
	}
	if s.UintField != 42 {
		t.Fail()
	}
	if s.Uint8Field != 42 {
		t.Fail()
	}
	if s.Uint16Field != 42 {
		t.Fail()
	}
	if s.Uint32Field != 42 {
		t.Fail()
	}
	if s.Uint64Field != 42 {
		t.Fail()
	}
	if s.CustomIntField != 42 {
		t.Fail()
	}
	if *s.PointerField != 42 {
		t.Fail()
	}
	if len(s.NotExistField) > 0 {
		t.Fail()
	}
	if s.RequiredField != "hello, world!" {
		t.Fail()
	}
	if len(s.UntaggedField) > 0 {
		t.Fail()
	}
	if !s.StructField.SubBoolField {
		t.Fail()
	}
	if len(s.StringSliceField) != 3 {
		t.Fail()
	}

	s2 := testFailStruct{}
	err = Unmarshal(&s2)
	if err == nil {
		t.Fail()
	}
}

var exampleEnv = map[string]string{
	"PORT":          "80",
	"DB_CONNECTION": "postgres://user:pass@localhost:5432/postgres",
}

func getExampleEnv(name string) string {
	return exampleEnv[name]
}

func ExampleUnmarshal() {
	env = getExampleEnv

	type myStruct struct {
		Port         int    `env:"PORT"`
		DBConnString string `env:"DB_CONNECTION" required:"true"`
	}

	config := myStruct{
		Port: 3000,
	}

	err := Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Port:", config.Port)
	fmt.Println("DBConnString:", config.DBConnString)
	// Output:
	// Port: 80
	// DBConnString: postgres://user:pass@localhost:5432/postgres
}
