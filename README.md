# envconfig
Package envconfig provides a function called Unmarshal that allows you to dump
environment variable values into a structure using struct tags

TODO: Add badges here

## Installing
You can install the package using the `go get` command:

    go get -u github.com/everettcaleb/envconfig

## Usage Example
For an environment where `PORT`, `DB_CONNECTION`, and `REDIS_HOST` are defined and we want to
require that a database connection string and Redis host are provided:

    package main

    import (
        "fmt"

        "github.com/everettcaleb/envconfig"
    )

    type appSettings struct {
        Port         int    `env:"PORT"`
        DBConnection string `env:"DB_CONNECTION" required:"true"`
        RedisPort    int    `env:"REDIS_PORT"`
        RedisHost    string `env:"REDIS_HOST" required:"true"`
    }

    func main() {
        // Initialize with defaults (if zero value is not desired)
        settings := appSettings{
            Port: 80,
            RedisPort: 6379
        }

        // Pull in environment variables
        err := envconfig.Unmarshal(&settings)
        if err != nil {
            fmt.Println(err)
            panic("failed to load configuration from environment, see above error")
        }

        // Print them out
        fmt.Println("Port:", settings.Port)
        fmt.Println("DB Connection:", settings.DBConnection)
        fmt.Println("Redis Port:", settings.RedisPort)
        fmt.Println("Redis Host:", settings.RedisHost)
    }

This example will print out:

    Port: 80
    DB Connection: ...
    Redis Port: 6379
    Redis Host: ...

## Functions
Below is documentation for exported functions:

### `func Unmarshal(i interface{}) error`
`Unmarshal` dumps environment variable values into a `struct` (pass a `*struct`)
It will automatically look up environment variables that are named in a struct tag.
For example, if you tag a field in your struct with ```env:"MY_VAR"``` it will
be filled with the value of that environment variable. If you would like an error
to be returned if the field is not set, use the tag ```required:"true"```. For boolean,
fields, any of the following values are valid (any case): `"true"`, `"false"`, `"yes"`, `"no"`,
`"t"`, `"f"`, `"y"`, `"n"`, `"1"`, `"0"`. Empty/unset environment variable values will not overwrite
the struct fields that are already set.

## TODO
Currently, `Unmarshal` doesn't support arrays, pointers, slices, or structs. I want to add support for
this later (slices will be comma separated probably).

## Contributing
Feel free to contribute by forking and creating a pull request. If you find any issues please
post them here and I'll resolve them when I get a chance.

## License
MIT License

Copyright &copy; 2019 Caleb Everett

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.