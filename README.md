# Retrool

Simple tool to repeat some actions after some time for GoLang

[![GoDoc](https://godoc.org/github.com/makssof/retrool?status.svg)](https://godoc.org/github.com/makssof/retrool)
[![Release](https://img.shields.io/github/v/release/makssof/retrool.svg)](https://github.com/makssof/retrool/releases/)
[![License](https://img.shields.io/github/license/makssof/retrool.svg)](https://github.com/makssof/retrool/blob/master/LICENSE)

## Content

- [Installation](#installation)
- [Usage](#usage)
- [Contribution](#contribution)

## Installation

To install the package just run:

```bash
go get -u github.com/makssof/retrool
```

## Usage

```go
import "github.com/makssof/retrool"

func main() {
    rand.Seed(time.Now().UnixNano())

    // tryOptions := retrool.DefaultTryOptions
    tryOptions := &retrool.TryOptions{
        StartInterval:        1,
        Addition:             1,
        AdditionCoefficient:  0.5,
        MaxTries:             10, 
        FailureDecisionMaker: retrool.DefaultDecisionMaker,
    }

    var n int

    retrool.Try(tryOptions, func (i int) bool {
        n = rand.Intn(101)

        return n < 10
    })
}
```

## Contribution

The tool is open-sourced under the [MIT](LICENSE) license.

If you will find some error, want to add something or ask a question - feel free to create an issue and/or make a pull request.

Any contribution is welcome.