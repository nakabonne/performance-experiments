# performance-experiments

## Benchmarks

```
❯ go test -benchtime=10s -benchmem -bench=. -cpuprofile=./cpu.out -memprofile=./mem.out .
goos: darwin
goarch: amd64
pkg: github.com/nakabonne/slice-vs-linked-list
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkSliceAppend-8        	284814075	        42.02 ns/op	      42 B/op	       0 allocs/op
BenchmarkLinkedListAppend-8   	228556610	        47.60 ns/op	      24 B/op	       1 allocs/op
PASS
ok  	github.com/nakabonne/slice-vs-linked-list	45.968s


❯ go tool pprof ./mem.out
Type: alloc_space
Time: May 20, 2021 at 4:25pm (JST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 32.19GB, 100% of 32.19GB total
Dropped 14 nodes (cum <= 0.16GB)
      flat  flat%   sum%        cum   cum%
   24.83GB 77.12% 77.12%    24.83GB 77.12%  github.com/nakabonne/slice-vs-linked-list.BenchmarkSliceAppend
    7.36GB 22.87%   100%     7.36GB 22.87%  github.com/nakabonne/slice-vs-linked-list.(*LinkedList).Append (inline)
         0     0%   100%     7.36GB 22.87%  github.com/nakabonne/slice-vs-linked-list.BenchmarkLinkedListAppend
         0     0%   100%    32.19GB   100%  testing.(*B).launch
         0     0%   100%    32.19GB   100%  testing.(*B).runN
```

