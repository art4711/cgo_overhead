# Overhead of cgo function calls #

In some other project I found out that cgo has a significant overhead
in calling functions in C. This is a simple benchmark to demonstrate
that.

## Why function pointers. ##

An observant reader will note that I'm calling my functions weirdly.
This is to defeat the compiler. At the first attempt I did this the
overhead of Go function calls was almost zero and that's because the
compiler helpfully inlined everything which defeated the whole purpose
of the benchmark. This is the best I can do, it's probably not
perfectly accurate (the call to `bar` is probably inlined inside the
func), but it shows us the magnitude(s) of the problem.

## results ##

    BenchmarkCFuncall-4 	10000000	       184 ns/op
    BenchmarkGoFuncall-4	1000000000	         2.00 ns/op

No comment, the results speak for themselves.
