# [stringcase][repo-url] [![Release][release-img]][release-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

This library provides some functions that convert string cases between Ada_Case, camelCase,
COBOL-CASE, kebab-case, MACRO_CASE, PascalCase, snake_case, Title Case, and Train-Case.
In addition, the functions `Capitalize`, `Lowerize`, and `Upperize` are provided to convert
string cases with a custom joiner character.

Essentially, these functions only target ASCII uppercase and lowercase letters for capitalization.
All characters other than ASCII uppercase and lowercase letters and ASCII numbers are removed as
word separators.

If you want to use some symbols as separators, specify those symbols in the `Separators` field of
`Options` struct and use the `〜CaseWithOptions` function for the desired case.
If you want to retain certain symbols and use everything else as separators, specify those symbols
in `Keep` field of `Options` struct and use the `〜CaseWithOptions` function for the desired case.

Additionally, you can specify whether to place word boundaries before and/or after non-alphabetic
characters with conversion options.
This can be set using the `SeparateBeforeNonAlphabets` and `SeparateAfterNonAlphabets` fields in
the `Options` struct.

The `〜Case` functions that do not take `Options` as an argument only place word boundaries after
non-alphabetic characters.
In other words, they behave as if
`SeparateBeforeNonAlphabets = false` and `SeparateAfterNonAlphabets = true`.


## Install

To get the latest version of this package, run the following command:

```bash
go get github.com/sttk/stringcase
```

To get a specific version of this package, run the following command:

```bash
go get github.com/sttk/stringcase@v1.0.0
```


## Usage

The functions contained in this library are executed as follows:

```go
func main() {
    input := "fooBar123Bar"
    snake := stringcase.SnakeCase(input)
    fmt.Printf("%s\n", snake)
    // => "foo_bar123_baz"
}
```

If you want the conversion to behave differently, use `〜CaseWithOptions`.

```go
func main() {
    opts := stringcase.Options{
        SeparateBeforeNonAlphabets: true,
        SeparateAfterNonAlphabets:  true,
    }
    input := "fooBar123Bar"
    snake := stringcase.SnakeCaseWithOptions(input, opts)
    fmt.Printf("%s\n", snake)
    // => "foo_bar_123_baz"
}
```

You can also use the functions `Capitalize`, `Lowerize`, and `Upperize` to convert
strings into capitalized, lowercased, or uppercased words joined by a custom joiner character:

```go
func main() {
    opts := stringcase.Options{
        SeparateBeforeNonAlphabets: true,
        SeparateAfterNonAlphabets:  true,
    }
    input := "fooBar123Bar"
    output := stringcase.Capitalize(input, '.', opts)
    fmt.Printf("%s\n", output)
    // => "Foo.Bar.123.Baz"
}
```

## Supporting Go versions

This library supports Go 1.18 or later.

### Actual test results for each Go version:

```sh
% go-fav -tags=github.sttk.errs.notify 1.26.2 1.25.9 1.24.13 1.23.12 1.20.14 \
         -ldflags="-linkmode=external" 1.22.12 1.21.13 1.19.13 1.18.10
go version go1.26.2 darwin/amd64
ok  	github.com/sttk/stringcase	0.585s	coverage: 100.0% of statements

go version go1.25.9 darwin/amd64
ok  	github.com/sttk/stringcase	0.540s	coverage: 100.0% of statements

go version go1.24.13 darwin/amd64
ok  	github.com/sttk/stringcase	0.535s	coverage: 100.0% of statements

go version go1.23.12 darwin/amd64
ok  	github.com/sttk/stringcase	0.349s	coverage: 100.0% of statements

go version go1.20.14 darwin/amd64
ok  	github.com/sttk/stringcase	0.345s	coverage: 100.0% of statements

go version go1.22.12 darwin/amd64
ok  	github.com/sttk/stringcase	0.534s	coverage: 100.0% of statements

go version go1.21.13 darwin/amd64
ok  	github.com/sttk/stringcase	0.534s	coverage: 100.0% of statements

go version go1.19.13 darwin/amd64
# github.com/sttk/stringcase.test
ld: warning: '***/go.o' has malformed LC_DYSYMTAB, expected 62 undefined symbols to start at index 8649, found 72 undefined symbols starting at index 64
ok  	github.com/sttk/stringcase	0.509s	coverage: 100.0% of statements

go version go1.18.10 darwin/amd64
# github.com/sttk/stringcase.test
ld: warning: '***/go.o' has malformed LC_DYSYMTAB, expected 59 undefined symbols to start at index 8513, found 69 undefined symbols starting at index 65
ok  	github.com/sttk/stringcase	0.513s	coverage: 100.0% of statements
```

## License

Copyright (C) 2024-2026 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk/stringcase
[release-img]: https://img.shields.io/badge/release-1.0.0-0f9999.svg
[release-url]: https://github.com/sttk/stringcase/releases
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk/stringcase.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk/stringcase
[ci-img]: https://github.com/sttk/stringcase/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/stringcase/actions?query=branch%3Amain
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT
