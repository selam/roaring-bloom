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
	"fmt"
	"hash/fnv"
	"math"
	"strconv"

	"github.com/RoaringBitmap/roaring/roaring64"
)

type bf struct {
	bitmap            *roaring64.Bitmap
	falsePositiveRate float64
	hashNum           int
	bitmapSize        uint64
	items             uint64
}

func New(maxSize uint64, falsePositiveRate float64) BloomFilter {
	hashNum := calculateHashNum(falsePositiveRate)
	bitmapSize := calculateBitmapSize(maxSize, falsePositiveRate)
	return &bf{
		bitmap:            roaring64.NewBitmap(),
		falsePositiveRate: falsePositiveRate,
		hashNum:           hashNum,
		bitmapSize:        bitmapSize,
		items:             0,
	}
}

func (bf *bf) Add(value interface{}) {
	bf.items = bf.items + 1
	for i := 0; i < bf.hashNum; i++ {
		key := getHash(value, i) % bf.bitmapSize
		bf.bitmap.Add(key)
	}
}

func (bf *bf) Contains(value interface{}) bool {
	// Implement the Contains method of your Bloom Filter
	for i := 0; i < bf.hashNum; i++ {
		key := getHash(value, i) % bf.bitmapSize
		if !bf.bitmap.Contains(key) {
			return false
		}
	}
	return true
}

func (bf *bf) FalsePositiveRate() float64 {
	return bf.falsePositiveRate
}

func (bf *bf) CurrentFalsePositiveRate() float64 {
	return float64(bf.bitmap.GetCardinality()) / float64(bf.bitmapSize) / math.Pow(float64(bf.hashNum), 2)
}

func (bf *bf) IsEmpty() bool {
	// Implement the IsEmpty method of your Bloom Filter
	return bf.bitmap.IsEmpty()
}

func (bf *bf) IsFull() bool {
	return bf.CurrentFalsePositiveRate() >= bf.FalsePositiveRate()
}

func (bf *bf) Size() uint64 {
	// Implement the Size method of your Bloom Filter
	return bf.items
}

func calculateBitmapSize(n uint64, f float64) uint64 {
	return uint64(math.Ceil(float64(n) * math.Log(f) * -1 / math.Pow(2*math.Log(2), 2)))
}

func calculateHashNum(f float64) int {
	return int(math.Ceil(math.Log2(f) * -1))
}

func getHash(value interface{}, seed int) uint64 {
	h := fnv.New64a()
	h.Write([]byte(strconv.Itoa(int(seed)) + fmt.Sprint(value)))
	return h.Sum64()
}
