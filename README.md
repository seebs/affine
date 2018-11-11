Affine Benchmark

This is a trivial benchmark that came into existence because I
was thinking about the API for doing this, and wanted to know about
performance of several options.

Results:

    AffineTypeSwitch-8     10000  141219 ns/op       0 B/op      0 allocs/op
    AffineInterface-8       2000  785927 ns/op  240001 B/op  30000 allocs/op
    AffineTypeSwitchPre-8  10000  154993 ns/op       0 B/op      0 allocs/op
    AffineInterfacePre-8    5000  317816 ns/op       0 B/op      0 allocs/op
    AffineConcrete-8      200000   10532 ns/op       0 B/op      0 allocs/op

So even when they don't require an allocation, interfaces can have a
very significant performance penalty. Also, using an interface's
functions can be significantly more expensive than doing a type switch.

I don't know whether any of these are easily improved.

