# bulk-marshal

> Benchmark different strategies for marshalling a slice of values.

## Strategies

Three different strategies are implemented and benchmarked for marshalling a
slice of values:

* **Serial:** The values are marshalled one by one.
* **Parallel**: A varying number of goroutine workers are spawned to marshall
  the values in parallel. It uses `chan`s, `sync.WaitGroup` and
  `context.Context` for synchronisation.
* **Pooled**: [pool][go-playground-pool] is used as a limited consumer goroutine
  pool. The pool size is varied similar to parallel strategy while benchmarking.

_Note: The number of workers range from 1x to 10x where x represents the
number of CPUs._

## Benchmarks

_**Run on Lenovo ThinkPad L450 Intel Core i5-5200U CPU @ 2.20GHz, 12GB 1600 MHz
DDR3 RAM using Go 1.9.4**_

Test data is loaded into struct values in the `init` method and used to
benchmark the different strategies using sub-benchmarks.

```
λ go test -v -run=XXX -bench=. ./strategies -benchtime=3s
goos: windows
goarch: amd64
pkg: github.com/sudo-suhas/bulk-marshal/strategies
Serial-4                 1000           3775684 ns/op           40576 B/op       1001 allocs/op
Parallel/WorkerCnt-1x-4                  2000           3194769 ns/op           41016 B/op       1007 allocs/op
Parallel/WorkerCnt-2x-4                  2000           3011639 ns/op           41029 B/op       1007 allocs/op
Parallel/WorkerCnt-3x-4                  2000           2865535 ns/op           41028 B/op       1007 allocs/op
Parallel/WorkerCnt-4x-4                  2000           2838516 ns/op           41047 B/op       1007 allocs/op
Parallel/WorkerCnt-5x-4                  2000           2849023 ns/op           41077 B/op       1008 allocs/op
Parallel/WorkerCnt-6x-4                  2000           2818502 ns/op           41065 B/op       1008 allocs/op
Parallel/WorkerCnt-7x-4                  2000           2830510 ns/op           41056 B/op       1007 allocs/op
Parallel/WorkerCnt-8x-4                  2000           2827008 ns/op           41125 B/op       1008 allocs/op
Parallel/WorkerCnt-9x-4                  2000           2837516 ns/op           41084 B/op       1008 allocs/op
Parallel/WorkerCnt-10x-4                 2000           2829010 ns/op           41114 B/op       1008 allocs/op
Pooled/WorkerCnt-1x-4                    1000           3740658 ns/op          340566 B/op       5129 allocs/op
Pooled/WorkerCnt-2x-4                    1000           3470466 ns/op          335072 B/op       5072 allocs/op
Pooled/WorkerCnt-3x-4                    1000           3444446 ns/op          332641 B/op       5046 allocs/op
Pooled/WorkerCnt-4x-4                    2000           3521501 ns/op          332321 B/op       5042 allocs/op
Pooled/WorkerCnt-5x-4                    2000           3446948 ns/op          332396 B/op       5042 allocs/op
Pooled/WorkerCnt-6x-4                    2000           3479972 ns/op          332627 B/op       5044 allocs/op
Pooled/WorkerCnt-7x-4                    1000           3541516 ns/op          332620 B/op       5043 allocs/op
Pooled/WorkerCnt-8x-4                    1000           3539514 ns/op          332662 B/op       5043 allocs/op
Pooled/WorkerCnt-9x-4                    2000           3483474 ns/op          332656 B/op       5042 allocs/op
Pooled/WorkerCnt-10x-4                   2000           3488978 ns/op          332736 B/op       5042 allocs/op
PASS
ok      github.com/sudo-suhas/bulk-marshal/strategies   121.666s
```

[![bulk marshal strategies][bulk-marshal-chart-image]][bulk-marshal-sheets]

## License

MIT © [Suhas Karanth][sudo-suhas]

[go-playground-pool]: https://github.com/go-playground/pool
[bulk-marshal-chart-image]: https://docs.google.com/spreadsheets/u/2/d/e/2PACX-1vTEXgwlGFETY60W_eREUjgOJXCeh4p4QG2OevZc9khJX8d_Q6eqeCEPW10Hffayut4iju0I1NjPYG_o/pubchart?oid=1852352816&amp;format=image
[bulk-marshal-sheets]: https://docs.google.com/spreadsheets/d/1ApBCF3vOIvoAMT-0U4BawYNJWgyLGZ8KhUjGGBDWS7g/edit?usp=sharing
[sudo-suhas]: https://github.com/sudo-suhas
