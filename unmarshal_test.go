package envconfig

import (
	"fmt"
	"testing"
)

type testStruct struct {
	BoolField    bool    `env:"BOOL_FIELD"`
	Float32Field float32 `env:"FLOAT_FIELD"`
	Float64Field float64 `env:"FLOAT_FIELD"`
	IntField     int     `env:"INT_FIELD"`
	Int8Field    int     `env:"INT_FIELD"`
	Int16Field   int     `env:"INT_FIELD"`
	Int32Field   int     `env:"INT_FIELD"`
	Int64Field   int     `env:"INT_FIELD"`
	StringField  string  `env:"STRING_FIELD"`
	UintField    uint    `env:"INT_FIELD"`
	Uint8Field   uint    `env:"INT_FIELD"`
	Uint16Field  uint    `env:"INT_FIELD"`
	Uint32Field  uint    `env:"INT_FIELD"`
	Uint64Field  uint    `env:"INT_FIELD"`

	NotExistField string `env:"NOT_EXISTS"`
	RequiredField string `env:"STRING_FIELD" required:"true"`
	UntaggedField string
}

type testFailStruct struct {
	RequiredNotExistField string `env:"NOT_EXISTS" required:"true"`
}

var mockEnv = map[string]string{
	"BOOL_FIELD":   "yes",
	"FLOAT_FIELD":  "3.14",
	"INT_FIELD":    "42",
	"STRING_FIELD": "hello, world!",
}

func getMockEnv(name string) string {
	return mockEnv[name]
}

func TestUnmarshal(t *testing.T) {
	env = getMockEnv
	s := testStruct{}

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
	if len(s.NotExistField) > 0 {
		t.Fail()
	}
	if s.RequiredField != "hello, world!" {
		t.Fail()
	}
	if len(s.UntaggedField) > 0 {
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
