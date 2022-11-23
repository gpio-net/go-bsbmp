module go-i2c-local

go 1.19

require (
	github.com/d2r2/go-logger v0.0.0-20210606094344-60e9d1233e22
	github.com/traulfs/tsb v0.0.0-20221121171133-f9ad3bbef23c
)

require (
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07 // indirect
	golang.org/x/sys v0.1.0 // indirect
)

replace github.com/d2r2/go-logger => ../go-logger-local
