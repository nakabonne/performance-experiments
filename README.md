# performance-experiments

```
❯ go test -benchtime=4s -benchmem -bench=. -cpuprofile=./cpu.out -memprofile=./mem.out .
goos: darwin
goarch: amd64
pkg: github.com/nakabonne/performance-experiments/slice_linkedlist
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkSliceAppend-8        	94101172	        46.21 ns/op	      41 B/op	       0 allocs/op
BenchmarkLinkedListAppend-8   	63808350	        73.83 ns/op	      48 B/op	       1 allocs/op
PASS
ok  	github.com/nakabonne/performance-experiments/slice_linkedlist	9.985s

slice_linkedlist on ᚠ main [?] took 21s

❯ go tool pprof ./mem.out
Type: alloc_space
Time: May 20, 2021 at 1:06pm (JST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) web
(pprof) top
Showing nodes accounting for 6.52GB, 99.95% of 6.52GB total
Dropped 24 nodes (cum <= 0.03GB)
      flat  flat%   sum%        cum   cum%
    3.72GB 56.96% 56.96%     3.72GB 56.96%  github.com/nakabonne/performance-experiments/slice_linkedlist.(*Slice).Append
    2.80GB 42.99% 99.95%     2.80GB 42.99%  github.com/nakabonne/performance-experiments/slice_linkedlist.(*LikedList).Append
         0     0% 99.95%     2.80GB 42.99%  github.com/nakabonne/performance-experiments/slice_linkedlist.BenchmarkLinkedListAppend
         0     0% 99.95%     3.72GB 56.96%  github.com/nakabonne/performance-experiments/slice_linkedlist.BenchmarkSliceAppend
         0     0% 99.95%     6.52GB 99.95%  testing.(*B).launch
         0     0% 99.95%     6.52GB 99.95%  testing.(*B).runN

❯ go tool pprof ./cpu.out
Type: cpu
Time: May 20, 2021 at 1:06pm (JST)
Duration: 9.38s, Total samples = 14.85s (158.37%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 10250ms, 69.02% of 14850ms total
Dropped 74 nodes (cum <= 74.25ms)
Showing top 10 nodes out of 77
      flat  flat%   sum%        cum   cum%
    2230ms 15.02% 15.02%     4690ms 31.58%  runtime.scanobject
    2000ms 13.47% 28.48%     2000ms 13.47%  runtime.(*mspan).init (inline)
    1310ms  8.82% 37.31%     2010ms 13.54%  runtime.findObject
     980ms  6.60% 43.91%      980ms  6.60%  runtime.memmove
     870ms  5.86% 49.76%      870ms  5.86%  runtime.usleep
     760ms  5.12% 54.88%     4310ms 29.02%  github.com/nakabonne/performance-experiments/slice_linkedlist.(*LikedList).Append
     620ms  4.18% 59.06%     3360ms 22.63%  github.com/nakabonne/performance-experiments/slice_linkedlist.(*Slice).Append
     530ms  3.57% 62.63%      530ms  3.57%  runtime.madvise
     530ms  3.57% 66.20%      660ms  4.44%  runtime.spanOf
     420ms  2.83% 69.02%      420ms  2.83%  sync.(*Mutex).Unlock (partial-inline)
```
