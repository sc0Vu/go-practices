# strings

In this directory, there are some benchmarks to basic string function.

# concatenate.go

* Benchmark for concatenating two strings:
```GO
first := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco `
second := `laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
result := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
```

* Run:
```BASH
$go run concatenate.go
```

* Result:
```BASH
Test string length: 445
Success to concatenate two strings by bytes, time pass: 39.105µs
Success to concatenate two strings by plus, time pass: 375ns
Success to concatenate two strings by copy, time pass: 5.436µs
Success to concatenate two strings by string builder, time pass: 19.93µs
```