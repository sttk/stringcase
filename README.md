# [stringcase][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

This library provides functions that convert strings into the following cases:

1. camelCase
2. COBOL-CASE
3. kebab-case
4. MACRO_CASE
5. PascalCase
6. snake_case
7. Train-Case

Essentially, these functions only target ASCII uppercase and lowercase letters for capitalization.
All characters other than ASCII uppercase and lowercase letters and ASCII numbers are removed as word separators.

If you want to use some symbols as separators, specify those symbols in the `Separators` field of `Options` struct and use the `〜CaseWithOptions` function for the desired case.
If you want to retain certain symbols and use everything else as separators, specify those symbols in `Keep` field of `Options` struct and use the `〜CaseWithOptions` function for the desired case.

When converting cases, there are several possible ways to handle numbers and symbols that remain in the converted string. This library supports the following four types of behavior:

1. Insert a word separator before a string of consecutive numbers or symbols<br>
(Specify `SeparateBeforeNonAlphabets = true, SeparateAfterNonAlphabets = false` of `Options`)
2. Insert a word separator after a string of consecutive numbers or symbols<br>
(Specify `SeparateBeforeNonAlphabets = false, SeparateAfterNonAlphabets = true` of `Options`)
3. Insert word separators both before and after a string of consecutive numbers or symbols<br>
(Specify `SeparateBeforeNonAlphabets = true, SeparateAfterNonAlphabets = true` of `Options`)
4. Do not insert before and after numbers or symbols as word separators<br>
(Specify `SeparateBeforeNonAlphabets = false, SeparateAfterNonAlphabets = false` of `Options`)

`〜Case` functions that do not take `Options` as an argument will behave as described in type 2. above.

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
ok  	github.com/sttk/stringcase	0.748s	coverage: 100.0% of statements

Now using version go1.19.13
go version go1.19.13 darwin/amd64
ok  	github.com/sttk/stringcase	0.366s	coverage: 100.0% of statements

Now using version go1.20.14
go version go1.20.14 darwin/amd64
ok  	github.com/sttk/stringcase	0.377s	coverage: 100.0% of statements

Now using version go1.21.13
go version go1.21.13 darwin/amd64
ok  	github.com/sttk/stringcase	0.372s	coverage: 100.0% of statements

Now using version go1.22.10
go version go1.22.10 darwin/amd64
ok  	github.com/sttk/stringcase	0.395s	coverage: 100.0% of statements

Now using version go1.23.6
go version go1.23.6 darwin/amd64
ok  	github.com/sttk/stringcase	0.379s	coverage: 100.0% of statements

Back to go1.23.6
Now using version go1.23.6
```

## License

Copyright (C) 2024-2025 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk/stringcase
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk/stringcase.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk/stringcase
[ci-img]: https://github.com/sttk/stringcase/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/stringcase/actions
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT
