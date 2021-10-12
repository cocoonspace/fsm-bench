# fsm-bench
FSM packages benchmark

## List of tested FSM packages :

* BenchmarkCocoonSpaceFSM : github.com/cocoonspace/fsm
* BenchmarkLooplabFSM : github.com/looplab/fsm

## Tests results

```
goos: darwin
goarch: amd64
pkg: github.com/cocoonspace/fsm-bench
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkCocoonSpaceFSM-12    	29371851	        40.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkLooplabFSM-12        	 2438946	       487.8 ns/op	     320 B/op	       4 allocs/op
```
