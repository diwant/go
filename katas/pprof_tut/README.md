# Shows Golang Profiling Using pprof

`myprogram.go' has three functions:

* SuperCompute - high compute intensity, burdens the cpu
* SuperHeap - high heap allocation/deallocation, burdens the heap and garbage collector
* SuperEfficientHeap - Same as SuperHeap but using bytes.Buffer vs string concatenation to demonstrate better heap usage.

# How to Use

Running `go run myprogram.go` will run all three above functions and create a `cpu.prof` and `mem.prof` file in the pprof_tut folder (or where you ran the go binary from).

## Go tool pprof
Now, run `go tool pprof cpu.prof` to check CPU profiling
Or, run `go tool pprof mem.prof` to check Memory (Heap) profiling

Now, you are in the pprof console. Some cool things to try in the pprof console:
* `topK` - type in `top` or `top5` or `top10` to see the top things using the resource you are profiling (CPU or Heap mem)
* `gif` - renders a callgraph that shows usage of the resource you are profiling (CPU or Heap mem) visually.  Will save a `profileNN.gif` file to the folder you ran `go tool pprof` from where `NN` is the index of the file (so you can do this several times without overwriting previous renders)

## Heap Memory Stats
Also, the console will print out heap statistics.  Here's what each bit means:
```
mem.Alloc - these are the bytes that were allocated and still in use
mem.TotalAlloc - what we allocated throughout the lifetime
mem.HeapAlloc - what’s being used on the heap right now
mem.HeapSys - this includes what is being used by the heap and what has been reclaimed but not given back out
```
