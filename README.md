# fastlog

Fast log is used in my interactive Monte Carlo (MCIT) work.
There we expect to call `Log` many thousand times per second
to compute queue priorities. We happen to use 32 bit floats
there, so this package is written around `float32`.

```golang
import "github.com/ajzaff/fastlog"

fastlog.Log2(x)
fastlog.Log(x)  // base e
```

The package is optimized for calling `Log` on small integers
between 1 and 10000 but has been tested up to 1e7.

This package is cool because I used a pretty naive but simple
Monte Carlo tree search routine to minimize the MSE of `Log2`
parameters on a representative test suite.

The products of this package are then used by MCIT itself to
improve search times!

The test suite can be viewed in [suite.go](suite.go).

## Benchmark

```text
goos: linux
goarch: amd64
pkg: github.com/ajzaff/fastlog/suite
cpu: Intel(R) Core(TM) i5-2520M CPU @ 2.50GHz
BenchmarkFastLog2-4      7883703              2971 ns/op
BenchmarkMathLog2-4       566901             38410 ns/op
```

## Caveats

The `math.Log` implementation does a lot of important special cases and bounds
checking which accounts for some of the differences in benchmark time.
Calling `Log` or `Log2` with values <= 0, +/-inf, NaN, etc. has undefined behavior.

It's worth stating again that this package was optimized for small integers.
Therefore, it may perform unpredcably for very large values or values very close to 0.

## References

The idea for this kind of fast Log was taken from the Leela chess project (lc0)
which in turn comes from the bit twiddling literature [TODO: Insert Source].

This package achieves slightly better MSE on our test suite. It's not exactly
a fair comparison but I'm happy with the result.

Go's own implementation uses a similar principle but with more estimation terms.