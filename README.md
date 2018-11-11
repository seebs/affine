Affine Benchmark

This is a trivial benchmark that came into existence because I
was thinking about the API for doing this, and wanted to know about
performance of several options.

Results:

    AffineTypeSwitch-8  10000  138034 ns/op       0 B/op      0 allocs/op
    AffineInterface-8    2000  773645 ns/op  240001 B/op  30000 allocs/op
    AffineConcrete-8   200000   11306 ns/op       0 B/op      0 allocs/op

So even when they don't require an allocation, interfaces can have a
very significant performance penalty. Also, using an interface's
functions can be significantly more expensive than doing a type switch.

I don't know whether any of these are easily improved.

