package suite

import (
	"math"
	"testing"

	"github.com/ajzaff/fastlog"
)

func TestMSESmallX(t *testing.T) {
	// Intuitively, MSE for small X should be lower than the full suite MSE.
	var mse float32
	for i := range 256 {
		x := float32(i + 1)
		res := fastlog.Log2(x)
		want := float32(math.Log2(float64(i + 1)))
		delta := want - res
		mse += delta * delta
	}
	mse /= 256
	suiteMSE := CalculateLog2MSE(1+fastlog.K, -fastlog.K)

	t.Log("Small X MSE:", mse)
	t.Log("Suite MSE:", suiteMSE)
	if suiteMSE < mse {
		t.Errorf("TestMSESmallX(): Log2 MSE on small X should not be worse than suite MSE: %f < %f", mse, suiteMSE)
	}
}

func TestNewConstantsBeatQuarterK(t *testing.T) {
	// Where log2(2^N*(1+f)) ~ N+f*(1+k-k*f) where N is the
	// exponent and f the fraction (mantissa), f>=0.
	// We show that the new constants c0, c1 beat the 1+k-k*f approximation on this test suite.
	const quarter = .25

	mse := CalculateLog2MSE(1+fastlog.K, -fastlog.K)
	kMSE := CalculateLog2MSE(1+quarter, -quarter)

	t.Log("MSE is", mse)
	t.Log("kMSE is", kMSE)

	if kMSE < mse {
		t.Errorf("TestNewConstantsBeatQuarterK(): new constants do not meet or exceed the k=1/4 approximation: %f < %f", kMSE, mse)
	}
}

func TestNewConstantsBeatLeelaApprox(t *testing.T) {
	// Where log2(2^N*(1+f)) ~ N+f*(1+k-k*f) where N is the
	// exponent and f the fraction (mantissa), f>=0.
	// We show that the our constant K beats the Leela approximation on this test suite.
	const lc0, lc1 = 1.3465552, -0.34655523

	mse := CalculateLog2MSE(1+fastlog.K, -fastlog.K)
	lMSE := CalculateLog2MSE(lc0, lc1)

	t.Log("MSE is", mse)
	t.Log("Leela MSE is", lMSE)

	if lMSE < mse {
		t.Errorf("TestNewConstantsBeatLeelaApprox(): new constants do not meet or exceed the Leela approximation: %f < %f", lMSE, mse)
	}
}

func TestFastLog2ImplSanity(t *testing.T) {
	for _, example := range suite {
		if got1, got2 := fastlog.Log2(example.X), FastLog2(example.X, 1+fastlog.K, -fastlog.K); got1 != got2 {
			t.Fatalf("TestFastLog2ImplSanity(%f): Log2 != FastLog2(c0,c1): %f != %f", example.X, got1, got2)
		}
	}
}

func TestOptimalLog2MSE(t *testing.T) {
	var mse float64
	for _, example := range suite {
		got := fastlog.Log2(example.X)
		delta := float64(got - example.Log2)
		mse += float64(delta * delta)
	}
	mse /= float64(len(suite))

	t.Log("MSE is", mse)
	if wantMse := 2.9706729847024478e-05; mse != wantMse {
		t.Errorf("TestOptimalLog2MSE(): Log2 does not match the advertized MSE against the test suite (got, want): %f != %f", mse, wantMse)
	}
}

func TestOptimalLogMSE(t *testing.T) {
	var mse float64
	for _, example := range suite {
		got := fastlog.Log(example.X)
		delta := float64(got - example.Log)
		mse += float64(delta * delta)
	}
	mse /= float64(len(suite))

	// MSE[Log(X)] ~ (1/ln 2)**2 * MSE[Log2(X)].
	t.Log("MSE is", mse)
	if wantMse := 1.4272725200787152e-05; mse != wantMse {
		t.Errorf("TestOptimalLogMSE(): Log does not match the advertized MSE against the test suite (got, want): %f != %f", mse, wantMse)
	}
}
