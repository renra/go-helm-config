# Go Helm Config

A project that provides access to both common and env-specific config variables overridden by envvars. It is originally written for projects managed by helm which is reflected in the assumed folder structure. Further work is planned to make this completely generic.

## Usage

```go
import (
  "fmt"
  "github.com/renra/go-helm-config/config"
)

func main() {
  conf := config.Load("charts/go-helm-logger/env", "testing")
}
```

This assumes you have a `values.yml` file under `charts/go-helm-logger/env` and that inside that file is a key called `env_vars`. If this file is not found, `config` panics.

Then it looks for `testing/values.yml` again with `env_vars` inside. If this file is not found, `config` will not care. Any keys here take precedence over the keys in the common `values.yml`.

Last but not least `config` reads all environment variables and stores them too, overriding all previously set keys. Since environment variables usually have upcased names, the names are downcased so `HERO_NAME` becomes `hero_name`.

The resulting type has two useful methods: `Get(string) interface{}` which returns the config value as is and `GetString(string) string` which uses a string type assertion on the config value.

See the examples folder for actual code examples. You can run it by `docker-compose up`.


## Known limitations

* made mainly for strings, lacks support for arbitrary types, if you put a number into your config, wrap it in apostrophes
* assumes rigid folder and file structure
