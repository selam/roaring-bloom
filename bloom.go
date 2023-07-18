package bloomfilter

/*
Copyright 2023 timu eren

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/
import (
	"hash"
	"hash/fnv"
	"math"
	"sync"

	"github.com/RoaringBitmap/roaring/roaring64"
)

type bf struct {
	bitmaps           []*roaring64.Bitmap
	falsePositiveRate float64
	hashNum           int
	bitmapSize        uint64
	items             uint64
}

func New(maxSize uint64, falsePositiveRate float64) *bf {
	bitmapSize := calculateBitmapSize(maxSize, falsePositiveRate)
	hashNum := calculateHashNum(bitmapSize, maxSize)
	bitmaps := make([]*roaring64.Bitmap, hashNum)
	for i := 0; i < hashNum; i++ {
		bitmaps[i] = roaring64.NewBitmap()
	}
	return &bf{
		bitmaps:           bitmaps,
		falsePositiveRate: falsePositiveRate,
		hashNum:           hashNum,
		bitmapSize:        bitmapSize,
		items:             0,
	}
}

func (bf *bf) Add(value []byte) {
	bf.items = bf.items + 1
	for i, h := range getHash(value, bf.hashNum, bf.bitmapSize) {
		bf.bitmaps[i].Add(h)
	}
}

func (bf *bf) Contains(value []byte) bool {
	for i, h := range getHash(value, bf.hashNum, bf.bitmapSize) {
		if !bf.bitmaps[i].Contains(h) {
			return false
		}
	}
	return true
}

func (bf *bf) FalsePositiveRate() float64 {
	return bf.falsePositiveRate
}

func (bf *bf) IsEmpty() bool {
	// Implement the IsEmpty method of your Bloom Filter
	return bf.bitmaps[0].IsEmpty()
}

func (bf *bf) Size() uint64 {
	// Implement the Size method of your Bloom Filter
	return bf.items
}

func calculateBitmapSize(n uint64, f float64) uint64 {
	return uint64(math.Ceil(-float64(n) * math.Log(f) / math.Pow(math.Log(2), 2)))
}

func calculateHashNum(bitSize uint64, items uint64) int {
	return int(math.Ceil(math.Log(2) * float64(bitSize) / float64(items)))
}

var pool = sync.Pool{
	New: func() interface{} {
		return fnv.New64a()
	},
}

func getHash(value []byte, hashNum int, bitmapSize uint64) []uint64 {
	hashes := make([]uint64, hashNum)
	h := pool.Get().(hash.Hash64)
	defer pool.Put(h)
	for i := uint(0); i < uint(hashNum); i++ {
		h.Reset()
		h.Write(value)                                     // add a salt to the hash function
		hashes[i] = uint64(h.Sum64()) % uint64(bitmapSize) // map the hash value to the range [0, m)
	}
	return hashes
}
