# performance-experiments

```
❯ go test -benchtime=4s -benchmem -bench=. -cpuprofile=./cpu.out -memprofile=./mem.out .
goos: darwin
goarch: amd64
pkg: github.com/nakabonne/performance-experiments
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkSliceAppend-8        	100000000	        52.35 ns/op	      49 B/op	       0 allocs/op
BenchmarkLinkedListAppend-8   	69016024	        69.61 ns/op	      48 B/op	       1 allocs/op
PASS
ok  	github.com/nakabonne/performance-experiments	15.983s

❯ go tool pprof ./mem.out
Type: alloc_space
Time: May 20, 2021 at 3:36pm (JST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 10.28GB, 100% of 10.28GB total
Dropped 15 nodes (cum <= 0.05GB)
      flat  flat%   sum%        cum   cum%
    5.65GB 54.91% 54.91%     5.65GB 54.91%  github.com/nakabonne/performance-experiments.(*LikedList).Append
    4.63GB 45.06%   100%     4.63GB 45.06%  github.com/nakabonne/performance-experiments.(*Slice).Append
         0     0%   100%     5.65GB 54.91%  github.com/nakabonne/performance-experiments.BenchmarkLinkedListAppend
         0     0%   100%     4.63GB 45.06%  github.com/nakabonne/performance-experiments.BenchmarkSliceAppend
         0     0%   100%    10.28GB   100%  testing.(*B).launch
         0     0%   100%    10.28GB   100%  testing.(*B).runN
(pprof) exit

❯ go tool pprof ./cpu.out
Type: cpu
Time: May 20, 2021 at 3:35pm (JST)
Duration: 14.58s, Total samples = 24.01s (164.64%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 14960ms, 62.31% of 24010ms total
Dropped 93 nodes (cum <= 120.05ms)
Showing top 10 nodes out of 88
      flat  flat%   sum%        cum   cum%
    3860ms 16.08% 16.08%     8620ms 35.90%  runtime.scanobject
    2890ms 12.04% 28.11%     2890ms 12.04%  runtime.(*mspan).init
    2080ms  8.66% 36.78%     3100ms 12.91%  runtime.findObject
    1690ms  7.04% 43.82%     1690ms  7.04%  runtime.usleep
    1190ms  4.96% 48.77%     1190ms  4.96%  runtime.memmove
     710ms  2.96% 51.73%      710ms  2.96%  sync.(*Mutex).Unlock
     690ms  2.87% 54.60%      970ms  4.04%  runtime.spanOf
     630ms  2.62% 57.23%      660ms  2.75%  runtime.pageIndexOf (inline)
     620ms  2.58% 59.81%     4040ms 16.83%  github.com/nakabonne/performance-experiments.(*Slice).Append
     600ms  2.50% 62.31%      710ms  2.96%  runtime.heapBitsSetType
```
