package bloomfilter

type BloomFilter interface {
	Add(value interface{})
	Contains(value interface{}) bool
	FalsePositiveRate() float64
	CurrentFalsePositiveRate() float64
	IsEmpty() bool
	IsFull() bool
	Size() uint64
}
