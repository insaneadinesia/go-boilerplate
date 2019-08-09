# Go(lang) Repository Pattern Boiler Plate
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/insaneadinesia/go-kredivo/blob/master/LICENSE)

Before you explore the whole codes, please understand the what is the repository pattern and how it works. Some of these articles may help you to understand the concept:
- [Trying Clean Architecture on Golang](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)
- [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Prerequisite
- This boiler plate use GIN as the framework. ([GIN](https://github.com/gin-gonic/gin))
- This boiler plate use GORM as the ORM. ([GORM](https://github.com/jinzhu/gorm))
- This boiler plate use Goose as database migration tool. Please intall it before. ([Goose](https://bitbucket.org/liamstask/goose/src/master))
- This boiler plate use Dep as the dependency management tool. Please install it before. ([Dep](https://github.com/golang/dep))
- For live reloading utility on local connection, you can use [Codegangsta - Gin](https://github.com/codegangsta/gin).

## Installation
1. Clone this repository
2. Run `dep ensure`
3. If you use [Codegangsta - Gin](https://github.com/codegangsta/gin), run `gin r main.go` or if not, you can run `go run main.go`

## License
See [LICENSE](https://github.com/insaneadinesia/go-boilerplate/blob/master/LICENSE).

## Localization
Thanks to [nicksnyder/go-i18n](nicksnyder/go-i18n) for the awesome package.

In this boiler plate, I just define two language (english and indonesian). You can override this according to what you need. The value of the language is store in `.yaml` file (in this case `en.yaml` and  `id.yaml`).
 
How to used :
```
...
import "${GOROOT}/go-boilerplate/lang"
...

func Test() {
	...
	val := lang.Translate("your_key_on_yaml_file", nil)

	// If you need to parse attribute
	attribute := map[string]string{"Attribute": "Value"}
	val := lang.Translate("your_key_on_yaml_file", attribute)
	...
}
...
```