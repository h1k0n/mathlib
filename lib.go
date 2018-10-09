package mathlib // import "github.com/h1k0n/mathlib"

/*
float invSqrt(float x)
{
	float xhalf = 0.5 * x;
	int i = *(int*)&x;
	i = 0x5f3759df - (i >> 1);
	x = *(float*)&i;
	x = x * (1.5 - xhalf * x * x);
	return x;
}
*/
import "C"

// Add return sum of two numbers
func Add(num1, num2 int) int {
	return num1 + num2
}

// InvSprt call this classic C function
func InvSprt(x float32) float32 {
	return float32(C.invSqrt(C.float(x)))
}
