package dsl

// requirements

// 1. All allocations happen at construction time, not inside the closure.
// 2. The closure returns a NodeOutput without allocating (except the unavoidable slice header).
// 3. The closure does not build new slices, strings, or byte arrays.
// 4. The closure does not call other nodes multiple times.
// 5. All byte fragments are prebuilt once and reused.
