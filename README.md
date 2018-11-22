Affine Benchmark

This is a trivial benchmark that came into existence because I
was thinking about the API for doing this, and wanted to know about
performance of several options.

Results:

    Old version (obsolete-ish):

    AffineTypeSwitch-8     10000  141219 ns/op       0 B/op      0 allocs/op
    AffineInterface-8       2000  785927 ns/op  240001 B/op  30000 allocs/op
    AffineTypeSwitchPre-8  10000  154993 ns/op       0 B/op      0 allocs/op
    AffineInterfacePre-8    5000  317816 ns/op       0 B/op      0 allocs/op
    AffineConcrete-8      200000   10532 ns/op       0 B/op      0 allocs/op

In the old version, I didn't realize at the time that the concrete
implementation was benefitting in part from being inlined. I also
concluded that the benchmark of the interface allocation was not
very interesting, so I dropped that, and added new tests:

    BenchmarkAffineConcreteInline-8           300000              5709 ns/op
    BenchmarkAffineConcrete-8                  10000            113512 ns/op
    BenchmarkAffineTypeSwitchHand-8            10000            125817 ns/op
    BenchmarkAffineTypeSwitchInlined-8         10000            129791 ns/op
    BenchmarkAffineTypeSwitchCall-8            10000            228741 ns/op
    BenchmarkAffineInterface-8                  5000            298248 ns/op

Note that there's multiple calls through interfaces happening here, and
they're happening several times per loop, on a loop of 10,000 items,
so the actual interface-overhead appears to be about 1-1.5ns per call,
give or take. This is similar to things I've seen elsewhere, except that at
least one test (a derivative of Egon Elbre's interface benchmark,
found at https://github.com/egonelbre/exp/tree/master/bench/iface) I
ran showed *zero* overhead for an interface compared to direct function
calls!

The surprising big difference here is that the performance *gap* between
calls to a type-switched function and calls through an interface dropped
significantly when switching from passing a 6-item struct to passing a
pointer to it, so probably a large portion of what I was seeing originally
was actually just the cost of passing parameters around.

To compare, all of these are performing essentially the same math
operations. What changes is how those get called:

	ConcreteInline: Using inlineable functions, make direct calls
	on concrete types.

	Concrete: Using `go:noinline` functions, make direct calls
	on concrete types.

	TypeSwitchHand: Make a call on an interface-typed thing, which
	uses a type switch to directly inline the code.

	TypeSwitchInlined: Use a type switch to call inlineable functions
	on concrete types.

	TypeSwitchCall: Use a type switch to call `go:noinline` functions
	on concrete types.

	Interface: Use the interface methods.

