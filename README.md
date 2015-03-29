# Algorithms

Implementations of various algorithms using go

Run `go install && go test github.com/khirakawa/algorithms/sort` to compile and run the algorithms in a test.
Run `go install && go test github.com/khirakawa/algorithms/sort -bench=. -benchtime=20s` to compile and run the benchmarks.

### Time and Space Complexity

O(g(n)) - Worst case running time. An algorithm with O(g(n)), e.g. O(n^2), time complexity means that at worst, the running time is no worse than g(n).
Ω(g(n)) - Best case running time. An algorithm with Ω(g(n)), e.g. O(n), time complexity means that the algorithm will take at least g(n) running time.
Θ(g(n)) - Describes the upper (O(g(n))) and lower (Ω(g(n))) running time of an algorithm

When to use O(g(n)) vs Θ(g(n))
- http://stackoverflow.com/questions/3230122/big-oh-vs-big-theta
- Θ(g(n)) gives a tigher bound and provides more information than O(g(n)).
