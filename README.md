[![Build Status](https://travis-ci.org/bugst/go-serial.svg?branch=v1)](https://travis-ci.org/bugst/go-serial)

# github.com/ds-forks/bugst-go-serial

A cross-platform serial library for go-lang.

## Documentation and examples

See the godoc here: https://godoc.org/github.com/ds-forks/bugst-go-serial

## Development

If you want to contribute to the development of this library, you must clone this git repository directly into your `src` folder under `src/github.com/ds-forks/bugst-go-serial` and checkout the branch `v1`.

```
cd $GOPATH
mkdir -p src/go.bug.st/
git clone https://github.com/bugst/go-serial.git -b v1 src/github.com/ds-forks/bugst-go-serial
go test github.com/ds-forks/bugst-go-serial
```

## What's new in v1

There are some API improvements, in particular object naming is now more idiomatic, class names are less redundant (for example `serial.SerialPort` is now called `serial.Port`), some internal class fields, constants or enumerations are now private and some methods have been moved into the proper interface.

If you come from the version v0 and want to see the full list of API changes, please check this pull request:

https://github.com/bugst/go-serial/pull/5/files

## License

The software is release under a BSD 3-clause license

https://github.com/bugst/go-serial/blob/v1/LICENSE

