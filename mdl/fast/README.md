# fast rendering

demo and attempt to render those components much more fastly.


## benchmarks

preliminary results

```sh
$ go test -bench=.
BenchmarkRenderWithTemplate-4   	   30000	     56485 ns/op
BenchmarkRenderWithFast-4       	 2000000	       968 ns/op
PASS
ok  	command-line-arguments	4.989s
```
