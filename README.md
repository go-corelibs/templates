[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/templates)
[![codecov](https://codecov.io/gh/go-corelibs/templates/graph/badge.svg?token=gzPs8JV2Cx)](https://codecov.io/gh/go-corelibs/templates)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/templates)](https://goreportcard.com/report/github.com/go-corelibs/templates)

# templates - text and html template utilities

A collection of utilties for working with text and html templates.

# Installation

``` shell
> go get github.com/go-corelibs/templates@latest
```

# Examples

## AddParseTree

``` go
func htmlExample() {
    initial, _ := htmlTemplate.New("initial").Parse(`{{ printf "initial test" }}`)
    another, _ := htmlTemplate.New("another").Parse(`{{ printf "another test" }}`)
    err := templates.AddParseTree(another, initial) // err == nil
}

func textExample() {
    initial, _ := textTemplate.New("initial").Parse(`{{ printf "initial test" }}`)
    another, _ := textTemplate.New("another").Parse(`{{ printf "another test" }}`)
    err := templates.AddParseTree(another, initial) // err == nil
}
```

## Lookup

Lookup is primarily useful for theme rendering systems where a parent template
(`tt` in this example) is pre-loaded with all of the available source files in
the current theme. Some themes may support custom content types and the theme
render needs to be able to lookup an existing template matching a list of names
in a specific order of priority or preference. This example is looking for
`page.html.tmpl` first, then `page.html` and finally `page.tmpl`. If none of
these are present, the `found` return value will be `nil`.

``` go
func main() {
    list := []string{
        "page.html.tmpl",
        "page.html",
        "page.tmpl",
    }
    found := templates.Lookup(tt, list...)
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
