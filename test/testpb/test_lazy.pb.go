package testpb

import (
	"github.com/ClaudioAlbertin/go-fast-protobuf/pkg/parser"
	"sync/atomic"
)

type LazyExample struct {
	payload      []byte
	cursors      parser.Cursors
	currentIndex uint32
	message      Example
}

func (l *LazyExample) GetField1() string {
	const fieldNum = 1
	if l.message.Field1 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field1 = string(l.payload[start:end])
	}

	return l.message.Field1
}

func (l *LazyExample) GetField2() string {
	const fieldNum = 2
	if l.message.Field2 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field2 = string(l.payload[start:end])
	}

	return l.message.Field2
}

func (l *LazyExample) GetField3() string {
	const fieldNum = 3
	if l.message.Field3 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field3 = string(l.payload[start:end])
	}

	return l.message.Field3
}

func (l *LazyExample) GetField4() string {
	const fieldNum = 4
	if l.message.Field4 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field4 = string(l.payload[start:end])
	}

	return l.message.Field4
}

func (l *LazyExample) GetField5() string {
	const fieldNum = 5
	if l.message.Field5 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field5 = string(l.payload[start:end])
	}

	return l.message.Field5
}

func (l *LazyExample) GetField6() string {
	const fieldNum = 6
	if l.message.Field6 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field6 = string(l.payload[start:end])
	}

	return l.message.Field6
}

func (l *LazyExample) GetField7() string {
	const fieldNum = 7
	if l.message.Field7 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field7 = string(l.payload[start:end])
	}

	return l.message.Field7
}

func (l *LazyExample) GetField8() string {
	const fieldNum = 8
	if l.message.Field8 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field8 = string(l.payload[start:end])
	}

	return l.message.Field8
}

func (l *LazyExample) GetField9() string {
	const fieldNum = 9
	if l.message.Field9 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field9 = string(l.payload[start:end])
	}

	return l.message.Field9
}

func (l *LazyExample) GetField10() string {
	const fieldNum = 10
	if l.message.Field10 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field10 = string(l.payload[start:end])
	}

	return l.message.Field10
}

func (l *LazyExample) GetField11() string {
	const fieldNum = 11
	if l.message.Field11 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field11 = string(l.payload[start:end])
	}

	return l.message.Field11
}

func (l *LazyExample) GetField12() string {
	const fieldNum = 12
	if l.message.Field12 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field12 = string(l.payload[start:end])
	}

	return l.message.Field12
}

func (l *LazyExample) GetField13() string {
	const fieldNum = 13
	if l.message.Field13 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field13 = string(l.payload[start:end])
	}

	return l.message.Field13
}

func (l *LazyExample) GetField14() string {
	const fieldNum = 14
	if l.message.Field14 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field14 = string(l.payload[start:end])
	}

	return l.message.Field14
}

func (l *LazyExample) GetField15() string {
	const fieldNum = 15
	if l.message.Field15 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field15 = string(l.payload[start:end])
	}

	return l.message.Field15
}

func (l *LazyExample) GetField16() string {
	const fieldNum = 16
	if l.message.Field16 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field16 = string(l.payload[start:end])
	}

	return l.message.Field16
}

func (l *LazyExample) GetField17() string {
	const fieldNum = 17
	if l.message.Field17 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field17 = string(l.payload[start:end])
	}

	return l.message.Field17
}

func (l *LazyExample) GetField18() string {
	const fieldNum = 18
	if l.message.Field18 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field18 = string(l.payload[start:end])
	}

	return l.message.Field18
}

func (l *LazyExample) GetField19() string {
	const fieldNum = 19
	if l.message.Field19 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field19 = string(l.payload[start:end])
	}

	return l.message.Field19
}

func (l *LazyExample) GetField20() string {
	const fieldNum = 20
	if l.message.Field20 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field20 = string(l.payload[start:end])
	}

	return l.message.Field20
}

func (l *LazyExample) GetField21() string {
	const fieldNum = 21
	if l.message.Field21 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field21 = string(l.payload[start:end])
	}

	return l.message.Field21
}

func (l *LazyExample) GetField22() string {
	const fieldNum = 22
	if l.message.Field22 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field22 = string(l.payload[start:end])
	}

	return l.message.Field22
}

func (l *LazyExample) GetField23() string {
	const fieldNum = 23
	if l.message.Field23 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field23 = string(l.payload[start:end])
	}

	return l.message.Field23
}

func (l *LazyExample) GetField24() string {
	const fieldNum = 24
	if l.message.Field24 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field24 = string(l.payload[start:end])
	}

	return l.message.Field24
}

func (l *LazyExample) GetField25() string {
	const fieldNum = 25
	if l.message.Field25 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field25 = string(l.payload[start:end])
	}

	return l.message.Field25
}

func (l *LazyExample) GetField26() string {
	const fieldNum = 26
	if l.message.Field26 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field26 = string(l.payload[start:end])
	}

	return l.message.Field26
}

func (l *LazyExample) GetField27() string {
	const fieldNum = 27
	if l.message.Field27 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field27 = string(l.payload[start:end])
	}

	return l.message.Field27
}

func (l *LazyExample) GetField28() string {
	const fieldNum = 28
	if l.message.Field28 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field28 = string(l.payload[start:end])
	}

	return l.message.Field28
}

func (l *LazyExample) GetField29() string {
	const fieldNum = 29
	if l.message.Field29 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field29 = string(l.payload[start:end])
	}

	return l.message.Field29
}

func (l *LazyExample) GetField30() string {
	const fieldNum = 30
	if l.message.Field30 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field30 = string(l.payload[start:end])
	}

	return l.message.Field30
}

func (l *LazyExample) GetField31() string {
	const fieldNum = 31
	if l.message.Field31 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field31 = string(l.payload[start:end])
	}

	return l.message.Field31
}

func (l *LazyExample) GetField32() string {
	const fieldNum = 32
	if l.message.Field32 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field32 = string(l.payload[start:end])
	}

	return l.message.Field32
}

func (l *LazyExample) GetField33() string {
	const fieldNum = 33
	if l.message.Field33 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field33 = string(l.payload[start:end])
	}

	return l.message.Field33
}

func (l *LazyExample) GetField34() string {
	const fieldNum = 34
	if l.message.Field34 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field34 = string(l.payload[start:end])
	}

	return l.message.Field34
}

func (l *LazyExample) GetField35() string {
	const fieldNum = 35
	if l.message.Field35 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field35 = string(l.payload[start:end])
	}

	return l.message.Field35
}

func (l *LazyExample) GetField36() string {
	const fieldNum = 36
	if l.message.Field36 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field36 = string(l.payload[start:end])
	}

	return l.message.Field36
}

func (l *LazyExample) GetField37() string {
	const fieldNum = 37
	if l.message.Field37 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field37 = string(l.payload[start:end])
	}

	return l.message.Field37
}

func (l *LazyExample) GetField38() string {
	const fieldNum = 38
	if l.message.Field38 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field38 = string(l.payload[start:end])
	}

	return l.message.Field38
}

func (l *LazyExample) GetField39() string {
	const fieldNum = 39
	if l.message.Field39 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field39 = string(l.payload[start:end])
	}

	return l.message.Field39
}

func (l *LazyExample) GetField40() string {
	const fieldNum = 40
	if l.message.Field40 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field40 = string(l.payload[start:end])
	}

	return l.message.Field40
}

func (l *LazyExample) GetField41() string {
	const fieldNum = 41
	if l.message.Field41 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field41 = string(l.payload[start:end])
	}

	return l.message.Field41
}

func (l *LazyExample) GetField42() string {
	const fieldNum = 42
	if l.message.Field42 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field42 = string(l.payload[start:end])
	}

	return l.message.Field42
}

func (l *LazyExample) GetField43() string {
	const fieldNum = 43
	if l.message.Field43 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field43 = string(l.payload[start:end])
	}

	return l.message.Field43
}

func (l *LazyExample) GetField44() string {
	const fieldNum = 44
	if l.message.Field44 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field44 = string(l.payload[start:end])
	}

	return l.message.Field44
}

func (l *LazyExample) GetField45() string {
	const fieldNum = 45
	if l.message.Field45 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field45 = string(l.payload[start:end])
	}

	return l.message.Field45
}

func (l *LazyExample) GetField46() string {
	const fieldNum = 46
	if l.message.Field46 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field46 = string(l.payload[start:end])
	}

	return l.message.Field46
}

func (l *LazyExample) GetField47() string {
	const fieldNum = 47
	if l.message.Field47 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field47 = string(l.payload[start:end])
	}

	return l.message.Field47
}

func (l *LazyExample) GetField48() string {
	const fieldNum = 48
	if l.message.Field48 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field48 = string(l.payload[start:end])
	}

	return l.message.Field48
}

func (l *LazyExample) GetField49() string {
	const fieldNum = 49
	if l.message.Field49 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field49 = string(l.payload[start:end])
	}

	return l.message.Field49
}

func (l *LazyExample) GetField50() string {
	const fieldNum = 50
	if l.message.Field50 == "" {
		if !l.cursors.IsPresent(lazyExampleCursorsOffsets[fieldNum]) {
			readUntil, _ := parser.FindCursors(lazyExampleCursorsOffsets, l.payload, l.cursors, int(atomic.LoadUint32(&l.currentIndex)), fieldNum)
			atomic.StoreUint32(&l.currentIndex, uint32(readUntil))
		}

		start, end := l.cursors.AtOffset(lazyExampleCursorsOffsets[fieldNum])
		l.message.Field50 = string(l.payload[start:end])
	}

	return l.message.Field50
}

func (l *LazyExample) UnmarshalFrom(payload []byte) error {
	l.payload = payload
	l.cursors = parser.NewCursors(400)
	return nil
}

var lazyExampleCursorsOffsets = map[int32]uint32{
	1:  0,
	2:  4,
	3:  8,
	4:  12,
	5:  16,
	6:  20,
	7:  24,
	8:  28,
	9:  32,
	10: 36,
	11: 40,
	12: 44,
	13: 48,
	14: 52,
	15: 56,
	16: 60,
	17: 64,
	18: 68,
	19: 72,
	20: 76,
	21: 80,
	22: 84,
	23: 88,
	24: 92,
	25: 96,
	26: 100,
	27: 104,
	28: 108,
	29: 112,
	30: 116,
	31: 120,
	32: 124,
	33: 128,
	34: 132,
	35: 136,
	36: 140,
	37: 144,
	38: 148,
	39: 152,
	40: 156,
	41: 160,
	42: 164,
	43: 168,
	44: 172,
	45: 176,
	46: 180,
	47: 184,
	48: 188,
	49: 192,
	50: 196,
}
