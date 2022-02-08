# RPN-Calculator

A simple RPN calculator with basic functions written in golang.

## Installation

If you have golang installed and have added `$GOBIN` to path, you can simply install using `make install`.

You can also install one of the published binaries for the operating system you want to use.

## Using the calculator

Once you have the calculator installed, you are able to use it via reverse polish notation:

```bash
$ calc 6 6 x
> 36

$ calc 4 2 + 4 + 8 - 2 / 3
> 3
```