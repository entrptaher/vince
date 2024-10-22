/*
 * Copyright 2020 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package buffer

import (
	"encoding/binary"
	"fmt"
	"log"
	"slices"
	"sort"
	"sync/atomic"

	"github.com/pkg/errors"
)

const (
	defaultCapacity = 64
)

// Buffer is equivalent of bytes.Buffer without the ability to read. It is NOT thread-safe.
//
// In UseCalloc mode, z.Calloc is used to allocate memory, which depending upon how the code is
// compiled could use jemalloc for allocations.
//
// In UseMmap mode, Buffer  uses file mmap to allocate memory. This allows us to store big data
// structures without using physical memory.
//
// MaxSize can be set to limit the memory usage.
type Buffer struct {
	padding uint64 // number of starting bytes used for padding
	offset  uint64 // used length of the buffer
	buf     []byte // backing slice for the buffer
	curSz   int    // capacity of the buffer
	maxSz   int    // causes a panic if the buffer grows beyond this size
}

func NewBuffer(capacity int) *Buffer {
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	}
	return &Buffer{
		buf:     make([]byte, capacity),
		curSz:   capacity,
		offset:  8,
		padding: 8,
	}
}

func NewBufferSlice(slice []byte) *Buffer {
	return &Buffer{
		offset: uint64(len(slice)),
		buf:    slice,
	}
}

func (b *Buffer) WithMaxSize(size int) *Buffer {
	b.maxSz = size
	return b
}

func (b *Buffer) IsEmpty() bool {
	return int(b.offset) == b.StartOffset()
}

// LenWithPadding would return the number of bytes written to the buffer so far
// plus the padding at the start of the buffer.
func (b *Buffer) LenWithPadding() int {
	return int(atomic.LoadUint64(&b.offset))
}

// LenNoPadding would return the number of bytes written to the buffer so far
// (without the padding).
func (b *Buffer) LenNoPadding() int {
	return int(atomic.LoadUint64(&b.offset) - b.padding)
}

// Bytes would return all the written bytes as a slice.
func (b *Buffer) Bytes() []byte {
	off := atomic.LoadUint64(&b.offset)
	return b.buf[b.padding:off]
}

// Grow would grow the buffer to have at least n more bytes. In case the buffer is at capacity, it
// would reallocate twice the size of current capacity + n, to ensure n bytes can be written to the
// buffer without further allocation. In UseMmap mode, this might result in underlying file
// expansion.
func (b *Buffer) Grow(n int) {
	if b.buf == nil {
		panic("z.Buffer needs to be initialized before using")
	}
	if b.maxSz > 0 && int(b.offset)+n > b.maxSz {
		err := fmt.Errorf(
			"z.Buffer max size exceeded: %d offset: %d grow: %d", b.maxSz, b.offset, n)
		panic(err)
	}
	if int(b.offset)+n < b.curSz {
		return
	}

	// Calculate new capacity.
	growBy := b.curSz + n
	// Don't allocate more than 1GB at a time.
	if growBy > 1<<30 {
		growBy = 1 << 30
	}
	// Allocate at least n, even if it exceeds the 1GB limit above.
	if n > growBy {
		growBy = n
	}
	b.curSz += growBy
	b.buf = slices.Grow(b.buf, b.curSz)
}

// Allocate is a way to get a slice of size n back from the buffer. This slice can be directly
// written to. Warning: Allocate is not thread-safe. The byte slice returned MUST be used before
// further calls to Buffer.
func (b *Buffer) Allocate(n int) []byte {
	b.Grow(n)
	off := b.offset
	b.offset += uint64(n)
	return b.buf[off:int(b.offset)]
}

// AllocateOffset works the same way as allocate, but instead of returning a byte slice, it returns
// the offset of the allocation.
func (b *Buffer) AllocateOffset(n int) int {
	b.Grow(n)
	b.offset += uint64(n)
	return int(b.offset) - n
}

func (b *Buffer) writeLen(sz int) {
	buf := b.Allocate(8)
	binary.BigEndian.PutUint64(buf, uint64(sz))
}

// SliceAllocate would encode the size provided into the buffer, followed by a call to Allocate,
// hence returning the slice of size sz. This can be used to allocate a lot of small buffers into
// this big buffer.
// Note that SliceAllocate should NOT be mixed with normal calls to Write.
func (b *Buffer) SliceAllocate(sz int) []byte {
	b.Grow(8 + sz)
	b.writeLen(sz)
	return b.Allocate(sz)
}

func (b *Buffer) StartOffset() int {
	return int(b.padding)
}

func (b *Buffer) WriteSlice(slice []byte) {
	dst := b.SliceAllocate(len(slice))
	assert(len(slice) == copy(dst, slice))
}

func (b *Buffer) SliceIterate(f func(slice []byte) error) error {
	if b.IsEmpty() {
		return nil
	}

	next := b.StartOffset()
	var slice []byte
	for next >= 0 {
		slice, next = b.Slice(next)
		if len(slice) == 0 {
			continue
		}
		if err := f(slice); err != nil {
			return err
		}
	}

	return nil
}

type LessFunc func(a, b []byte) bool
type sortHelper struct {
	offsets []int
	b       *Buffer
	tmp     *Buffer
	less    LessFunc
	small   []int
}

func (s *sortHelper) sortSmall(start, end int) {
	s.tmp.Reset()
	s.small = s.small[:0]
	next := start
	for next >= 0 && next < end {
		s.small = append(s.small, next)
		_, next = s.b.Slice(next)
	}

	// We are sorting the slices pointed to by s.small offsets, but only moving the offsets around.
	sort.Slice(s.small, func(i, j int) bool {
		left, _ := s.b.Slice(s.small[i])
		right, _ := s.b.Slice(s.small[j])
		return s.less(left, right)
	})
	// Now we iterate over the s.small offsets and copy over the slices. The result is now in order.
	for _, off := range s.small {
		_, _ = s.tmp.Write(rawSlice(s.b.buf[off:]))
	}
	assert(end-start == copy(s.b.buf[start:end], s.tmp.Bytes()))
}

func assert(b bool) {
	if !b {
		log.Fatalf("%+v", errors.Errorf("Assertion failure"))
	}
}
func check(err error) {
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
func check2(_ interface{}, err error) {
	check(err)
}

func (s *sortHelper) merge(left, right []byte, start, end int) {
	if len(left) == 0 || len(right) == 0 {
		return
	}
	s.tmp.Reset()
	check2(s.tmp.Write(left))
	left = s.tmp.Bytes()

	var ls, rs []byte

	copyLeft := func() {
		assert(len(ls) == copy(s.b.buf[start:], ls))
		left = left[len(ls):]
		start += len(ls)
	}
	copyRight := func() {
		assert(len(rs) == copy(s.b.buf[start:], rs))
		right = right[len(rs):]
		start += len(rs)
	}

	for start < end {
		if len(left) == 0 {
			assert(len(right) == copy(s.b.buf[start:end], right))
			return
		}
		if len(right) == 0 {
			assert(len(left) == copy(s.b.buf[start:end], left))
			return
		}
		ls = rawSlice(left)
		rs = rawSlice(right)

		// We skip the first 4 bytes in the rawSlice, because that stores the length.
		if s.less(ls[8:], rs[8:]) {
			copyLeft()
		} else {
			copyRight()
		}
	}
}

func (s *sortHelper) sort(lo, hi int) []byte {
	assert(lo <= hi)

	mid := lo + (hi-lo)/2
	loff, hoff := s.offsets[lo], s.offsets[hi]
	if lo == mid {
		// No need to sort, just return the buffer.
		return s.b.buf[loff:hoff]
	}

	// lo, mid would sort from [offset[lo], offset[mid]) .
	left := s.sort(lo, mid)
	// Typically we'd use mid+1, but here mid represents an offset in the buffer. Each offset
	// contains a thousand entries. So, if we do mid+1, we'd skip over those entries.
	right := s.sort(mid, hi)

	s.merge(left, right, loff, hoff)
	return s.b.buf[loff:hoff]
}

// SortSlice is like SortSliceBetween but sorting over the entire buffer.
func (b *Buffer) SortSlice(less func(left, right []byte) bool) {
	b.SortSliceBetween(b.StartOffset(), int(b.offset), less)
}

func (b *Buffer) SortSliceBetween(start, end int, less LessFunc) {
	if start >= end {
		return
	}
	if start == 0 {
		panic("start can never be zero")
	}

	var offsets []int
	next, count := start, 0
	for next >= 0 && next < end {
		if count%1024 == 0 {
			offsets = append(offsets, next)
		}
		_, next = b.Slice(next)
		count++
	}
	assert(len(offsets) > 0)
	if offsets[len(offsets)-1] != end {
		offsets = append(offsets, end)
	}

	szTmp := int(float64((end-start)/2) * 1.1)
	s := &sortHelper{
		offsets: offsets,
		b:       b,
		less:    less,
		small:   make([]int, 0, 1024),
		tmp:     NewBuffer(szTmp),
	}
	defer func() { _ = s.tmp.Release() }()

	left := offsets[0]
	for _, off := range offsets[1:] {
		s.sortSmall(left, off)
		left = off
	}
	s.sort(0, len(offsets)-1)
}

func rawSlice(buf []byte) []byte {
	sz := binary.BigEndian.Uint64(buf)
	return buf[:8+int(sz)]
}

// Slice would return the slice written at offset.
func (b *Buffer) Slice(offset int) ([]byte, int) {
	if offset >= int(b.offset) {
		return nil, -1
	}

	sz := binary.BigEndian.Uint64(b.buf[offset:])
	start := offset + 8
	next := start + int(sz)
	res := b.buf[start:next]
	if next >= int(b.offset) {
		next = -1
	}
	return res, next
}

// SliceOffsets is an expensive function. Use sparingly.
func (b *Buffer) SliceOffsets() []int {
	next := b.StartOffset()
	var offsets []int
	for next >= 0 {
		offsets = append(offsets, next)
		_, next = b.Slice(next)
	}
	return offsets
}

func (b *Buffer) Data(offset int) []byte {
	if offset > b.curSz {
		panic("offset beyond current size")
	}
	return b.buf[offset:b.curSz]
}

// Write would write p bytes to the buffer.
func (b *Buffer) Write(p []byte) (n int, err error) {
	n = len(p)
	b.Grow(n)
	assert(n == copy(b.buf[b.offset:], p))
	b.offset += uint64(n)
	return n, nil
}

// Reset would reset the buffer to be reused.
func (b *Buffer) Reset() {
	clear(b.buf)
	b.offset = uint64(b.StartOffset())
}

// Release would free up the memory allocated by the buffer. Once the usage of buffer is done, it is
// important to call Release, otherwise a memory leak can happen.
func (b *Buffer) Release() error {
	b.buf = nil
	return nil
}