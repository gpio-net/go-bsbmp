module examples/example1.go

go 1.19

require (
	github.com/d2r2/go-bsbmp v0.0.0-20190515110334-3b4b3aea8375
	github.com/d2r2/go-i2c v0.0.0-20191123181816-73a8a799d6bc
	github.com/d2r2/go-logger v0.0.0-20210606094344-60e9d1233e22
)

require (
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07 // indirect
	github.com/traulfs/tsb v0.0.0-20221121171133-f9ad3bbef23c // indirect
	golang.org/x/sys v0.1.0 // indirect
)

replace github.com/d2r2/go-bsbmp => /.

replace github.com/d2r2/go-i2c => /go-i2c-local

replace github.com/d2r2/go-logger => /go-logger-local
