package bloomfilter

type BloomFilter interface {
	Add(value []byte)
	Contains(value []byte) bool
	IsEmpty() bool
	Size() uint64
}
