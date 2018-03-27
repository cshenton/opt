# Opt
Thread-safe blackbox optimsers for Go.

## Layout

- `.`: key interfaces, algorithms.
- `bench`: functions and dataset to benchmark optimisers
    - rastrigin etc. for low dim
    - neural network controller for high dim
- `examples`: examples of single-thread, multi-thread, and distributed use


## Internal

- batch utility computation
- online utility computation
- adaptive sampling


## Algorithms

- SNES
- SNES-as
- Block-free SNES (continuous generation)
- XNES
- XNES-as


## Internal, Float64

We want to share adaptive sampling in a nice way. Also since XNES, SNES are just
distribution choices, we should be able to, so there's a few components:

- Distribution choice
- Adaptive sampling, yes/no
- Generation management (storing generation results, blocking on Search)

## Internal, Byte

For another day.

## Notes

- split benchmarks up for different problem spaces
- have a table here recommending algorithms for each problem
- proper godoc