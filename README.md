# go-steun
Handy misc tools for go development
![example branch parameter]()
[![Build Status](https://github.com/cthulhu/go-steun/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/cthulhu/go-steun/actions)  [![Goreport](https://goreportcard.com/badge/github.com/cthulhu/go-steun)](https://goreportcard.com/badge/github.com/cthulhu/go-steun) [![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/cthulhu/go-steun/master/LICENSE)

## Content

* Fixture - Easy text fixtures management
* Pointers - Some helpers to generate pointers from const values
* Raise - Wrapper that panics on errors
* Structs - Set/Get helpers to write/read random fields from structs by their string names
* TimeId - a set of helpers to work with Date/Time as a timeID in YYDDMM format
* URI - a set of helpers duplicating functionality of JS's encodeuricomponent/decodeuricomponent and others
* Retry - simple retry wrapper, supports Before and After retry

## Installation

    go get github.com/cthulhu/go-steun

  OR alternatively you can query individual packages

    go get github.com/cthulhu/go-steun/{package}

  OR just import it in your source file and let go-dep solve it for you

## Contribution

Any possible contributions are welcome. See [Contributing](CONTRIBUTING.md) to get started.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/cthulhu/go-steun/issues).



## License

  See [LICENSE](LICENSE) File
