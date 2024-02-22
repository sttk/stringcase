# [stringcase][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-image]][ci-url] [![MIT License][mit-img]][mit-url]

This library provides some functions that convert string cases between camelCase, COBOL-CASE, kebab-case, MACRO_CASE, PascalCase, snake_case and Train-Case.

Basically, these functions targets the upper and lower cases of only ASCII alphabets for capitalization, and all characters except ASCII alphabets and ASCII numbers are eliminated as word separators.

To limit characters using as separators, the functions named like *_with_sep are provided, and to keep specified characters, the functions named like *_with_keep are provided.

## Install

Write the following import declaration, then run `go mod tidy`.

```
import "github.com/sttk/stringcase"
```

## Usage

This functions contained in this library are excuted as follows:

```
func main() {
    input := "foo-bar-baz"
    camel := stringcase.CamelCase(input)
    fmt.Printf("%s\n", camel)
    // => "fooBarBaz"
}
```

## Supporting Go versions

This library supports Go 1.18 or later.

### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk/stringcase	0.311s	coverage: 100.0% of statements

Now using version go1.19.13
go version go1.19.13 darwin/amd64
ok  	github.com/sttk/stringcase	0.229s	coverage: 100.0% of statements

Now using version go1.20.14
go version go1.20.14 darwin/amd64
ok  	github.com/sttk/stringcase	0.260s	coverage: 100.0% of statements

Now using version go1.21.7
go version go1.21.7 darwin/amd64
ok  	github.com/sttk/stringcase	0.234s	coverage: 100.0% of statements

Now using version go1.22
go version go1.22.0 darwin/amd64
ok  	github.com/sttk/stringcase	0.231s	coverage: 100.0% of statements

Back to go1.22
Now using version go1.22
```

## License

Copyright (C) 2024 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk/stringcase
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk/stringcase.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk/stringcase
[ci-img]: https://github.com/sttk/stringcase/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/stringcase/actions
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT
