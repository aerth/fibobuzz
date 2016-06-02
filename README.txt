Simple! I wanted to run a primal fizzbuzz test with fibonacci input, but ran into trouble at the uint64 
limit.

So, "fibonize", "increment", and "fizzbuzz" all use Go's "math/big" package to count REALLY high!

And fizzbuzz checks for primes.  The technique I am using for primality check is probably not going to win 
you a record. This project was me learning the "math/big" and integrating os.Stdin/os.Stdout pipelines.

There is a small delay in the counting to go easy on the CPU when checking very large numbers for 
primality.
