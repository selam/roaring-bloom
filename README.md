# Bloom Filter

This repository contains a simple implementation of a Bloom Filter using with roaring bitmaps

## Overview

A Bloom Filter is a probabilistic data structure used for efficient membership testing. It can quickly determine whether an element is likely to be a member of a set or not. However, it may introduce false positives (indicating that an element is a member when it's not) with a certain probability.

This Bloom Filter implementation provides the following features:

- `Add`: Adds an element to the Bloom Filter.
- `Contains`: Checks if an element is likely to be a member of the Bloom Filter.
- `FalsePositiveRate`: Returns the target false positive rate of the Bloom Filter.
- `CurrentFalsePositiveRate`: Returns the current false positive rate of the Bloom Filter.
- `IsEmpty`: Checks if the Bloom Filter is empty.
- `IsFull`: Checks if the Bloom Filter is full.
- `Size`: Returns the number of elements inserted into the Bloom Filter.
- `Len`: Returns the number of bits set in the Bloom Filter.

## Getting Started

### Installation

To use this Bloom Filter implementation, you need to have Go installed. You can clone this repository or use Go modules to import it into your project.
```
go get github.com/username/repo/bloomfilter
```

### Usage

Here is an example of how to use the Bloom Filter:

```go
package main

import (
	"fmt"

	"github.com/selam/roaring-bloom"
)

func main() {
	maxSize := uint64(1000)
	falsePositiveRate := 0.001
	bf := bloomfilter.New(maxSize, falsePositiveRate)

	bf.Add("value1")
	bf.Add("value2")

	fmt.Println(bf.Contains("value1")) // true
	fmt.Println(bf.Contains("value3")) // false

	fmt.Println("False Positive Rate:", bf.FalsePositiveRate())
	fmt.Println("Current False Positive Rate:", bf.CurrentFalsePositiveRate())
	fmt.Println("Is Empty?", bf.IsEmpty())
	fmt.Println("Is Full?", bf.IsFull())
	fmt.Println("Size:", bf.Size())
	fmt.Println("Len:", bf.Len())
}
```

## Contributing

Contributions to this Bloom Filter implementation are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## Performance

The performance of the Bloom Filter can vary based on the size of the filter and the false positive rate. It provides a trade-off between space efficiency and accuracy. It is recommended to adjust the size and false positive rate according to the specific use case to achieve optimal performance.

## License
This Bloom Filter implementation is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.