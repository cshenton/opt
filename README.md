# Opt
Thread-safe blackbox optimsers for Go.

## ToDo

- low dim benchmark functions, runner.
- single and multi thread examples using benchmark functions
- some introspection / stopping functions for SNES
- Travis, coveralls
-

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
    - When you update search params
    - Do Hypothetical update with increased search rate
    - Compute likelihood ratios for current generation
    - weighted-MW test between actual, hypothetical ranking
    - adjust learning rate.

So I'm thinking we first do a functional implementation of the updates. Then in
each main class, we just have a 'adaptive' param which activates a branch on update.

What about continuous learning? We can just use the active generation, but we'll have
to decide also how frequently to test for adjustments. Since we'll be using a smaller
LR in general, we could adjust every step?


## Algorithms

- SNES
- SNES-as
- Hyperprojection SNES
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