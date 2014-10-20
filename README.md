# Semantic Versioning for Go

![Build](https://img.shields.io/travis/hansrodtang/semver.svg?style=flat)  ![Coverage](https://img.shields.io/coveralls/hansrodtang/semver.svg?style=flat) ![Issues](https://img.shields.io/github/issues/hansrodtang/semver.svg?style=flat) ![Tip](https://img.shields.io/gratipay/hansrodtang.svg?style=flat)

A [Semantic Versioning](http://semver.org/) library for [Go](http://golang.org).

Covers version `2.0.0` of the semver specification.

Documentation on the syntax for the `Satifies()` method can be found  [here](https://www.npmjs.org/doc/misc/semver.html).


## Installation

```
  go get github.com/hansrodtang/semver
```

## Usage

```go
import github.com/hansrodtang/semver

v1, error := semver.New("1.5.0")
// do something with error
if v1.Satisfies("^1.0.0") {
  // do something
}
```

## Benchmarks

Test | Iterations | Time
------------------------|-----------|------------
BenchmarkParseSimple    | 5000000   | 356 ns/op
BenchmarkParseComplex   | 1000000   | 2200 ns/op
BenchmarkCompareSimple  | 500000000 | 3.85 ns/op
BenchmarkCompareComplex	| 100000000	| 17.3 ns/op

Run the benchmarks yourself with:

```
go test github.com/hansrodtang/semver -bench=.
```


## License

This software is licensed under the [MIT license](LICENSE.md).