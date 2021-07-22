# gosync

```sh
go test -count 1000 .
ok  	_/Users/telepenin/src/go	0.200s
```

## benchmarks

- with sync.Mutex:

```sh
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkMain-8   	    2875	    488201 ns/op
PASS
ok  	_/Users/telepenin/src/go	1.577s
```

- with sync/atomic

```sh
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkMain-8   	   32653	     35917 ns/op
PASS
ok  	_/Users/telepenin/src/go	1.672s
```

- without async way

```sh
$go test -bench=.
goos: darwin
goarch: amd64
BenchmarkMain-8   	  535876	      2235 ns/op
PASS
ok  	_/Users/telepenin/src/go	1.351s
```
