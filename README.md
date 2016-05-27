Input Performance Experiment
=============================

I tried to determine the fastest way to read many numbers from stdin.

    $ yes 3215649832086464061 | go test -bench .
    testing: warning: no tests to run
    PASS
    BenchmarkScanf-2       	 2000000	       884 ns/op
    BenchmarkScan-2        	 2000000	       805 ns/op
    BenchmarkBufferedRead-2	 3000000	       444 ns/op
    BenchmarkBufferedScan-2	20000000	        77.6 ns/op
    ok  	github.com/mhutter/input-performance	8.572s
