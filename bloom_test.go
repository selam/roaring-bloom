package bloomfilter

import (
	"strconv"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	maxSize := uint64(1000)
	falsePositiveRate := 0.001
	bf := New(maxSize, falsePositiveRate)

	// Add values to the Bloom Filter
	values := []string{"value1", "value2", "value3"}
	for _, value := range values {
		bf.Add(value)
		if !bf.Contains(value) {
			t.Errorf("Failed to add value: %s", value)
		}
	}

	// Check if values exist in the Bloom Filter
	for _, value := range values {
		contains := bf.Contains(value)
		if !contains {
			t.Errorf("Failed to find value: %s", value)
		}
	}

	// Check false positive rate
	targetFPRate := bf.FalsePositiveRate()
	currentFPRate := bf.CurrentFalsePositiveRate()
	if currentFPRate > targetFPRate {
		t.Errorf("False positive rate exceeded the target. Target: %f, Current: %f", targetFPRate, currentFPRate)
	}

	// Check if Bloom Filter is empty
	isEmpty := bf.IsEmpty()
	if isEmpty {
		t.Errorf("Bloom Filter reported as empty, but it should have values")
	}

	// Check if Bloom Filter is full
	isFull := bf.IsFull()
	if isFull {
		t.Errorf("Bloom Filter reported as full, but it should not be full yet")
	}

	// Check the size of the Bloom Filter
	expectedSize := uint64(len(values))
	size := bf.Size()
	if size != expectedSize {
		t.Errorf("Unexpected size. Expected: %d, Actual: %d", expectedSize, size)
	}
}

func BenchmarkBloomFilter_Add(b *testing.B) {
	maxSize := uint64(1000000)
	falsePositiveRate := 0.001
	bf := New(maxSize, falsePositiveRate)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Add(strconv.Itoa(i))
	}
}
