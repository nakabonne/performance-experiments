# performance-experiments
The performance of slice looks a bit better than linked-list.

## Benchmarks

```
❯ go test -benchtime=4s -benchmem -bench=. -cpuprofile=./cpu.out -memprofile=./mem.out .
goos: darwin
goarch: amd64
pkg: github.com/nakabonne/slice-vs-linked-list
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkSliceAppend-8        	119230053	        40.10 ns/op	      41 B/op	       0 allocs/op
BenchmarkLinkedListAppend-8   	59566522	        68.79 ns/op	      48 B/op	       1 allocs/op
PASS
ok  	github.com/nakabonne/slice-vs-linked-list	13.598s
```

```
❯ go tool pprof ./mem.out
Type: alloc_space
Time: May 20, 2021 at 3:49pm (JST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 10.97GB, 100% of 10.98GB total
Dropped 12 nodes (cum <= 0.05GB)
      flat  flat%   sum%        cum   cum%
    8.31GB 75.69% 75.69%     8.31GB 75.69%  github.com/nakabonne/slice-vs-linked-list.(*Slice).Append
    2.67GB 24.29%   100%     2.67GB 24.29%  github.com/nakabonne/slice-vs-linked-list.(*LinkedList).Append
         0     0%   100%     2.67GB 24.29%  github.com/nakabonne/slice-vs-linked-list.BenchmarkLinkedListAppend
         0     0%   100%     8.31GB 75.69%  github.com/nakabonne/slice-vs-linked-list.BenchmarkSliceAppend
         0     0%   100%    10.97GB   100%  testing.(*B).launch
         0     0%   100%    10.97GB   100%  testing.(*B).runN
```

```
❯ go tool pprof ./cpu.out
Type: cpu
Time: May 20, 2021 at 3:48pm (JST)
Duration: 13.01s, Total samples = 23.27s (178.85%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 17530ms, 75.33% of 23270ms total
Dropped 90 nodes (cum <= 116.35ms)
Showing top 10 nodes out of 72
      flat  flat%   sum%        cum   cum%
    4030ms 17.32% 17.32%     8020ms 34.46%  runtime.scanobject
    2520ms 10.83% 28.15%     2520ms 10.83%  runtime.(*mspan).init
    2350ms 10.10% 38.25%     3710ms 15.94%  runtime.findObject
    2200ms  9.45% 47.70%     2200ms  9.45%  runtime.usleep
    1540ms  6.62% 54.32%     1540ms  6.62%  runtime.memmove
    1520ms  6.53% 60.85%     7120ms 30.60%  github.com/nakabonne/slice-vs-linked-list.(*Slice).Append
    1110ms  4.77% 65.62%     1340ms  5.76%  runtime.spanOf
    1040ms  4.47% 70.09%     1040ms  4.47%  sync.(*Mutex).Unlock
     660ms  2.84% 72.93%     1680ms  7.22%  runtime.wbBufFlush1
     560ms  2.41% 75.33%      560ms  2.41%  runtime.(*mspan).refillAllocCache
```

