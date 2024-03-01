package parser

import (
	"encoding/binary"
	"fmt"
	"io"
)

var (
	// ErrInvalidLength is returned when decoding a negative length.
	ErrInvalidLength = fmt.Errorf("proto: negative length found during unmarshaling")
	// ErrIntOverflow is returned when decoding a varint representation of an integer that overflows 64 bits.
	ErrIntOverflow = fmt.Errorf("proto: integer overflow")
)

// Cursors holds start and end indexes for each known field in a message.
type Cursors []byte

func NewCursors(size int) Cursors {
	return make(Cursors, size)
}

func (c Cursors) AtOffset(cursorOffset uint32) (uint32, uint32) {
	return binary.LittleEndian.Uint32(c[cursorOffset : cursorOffset+4]), binary.LittleEndian.Uint32(c[cursorOffset+4 : cursorOffset+8])
}

func (c Cursors) IsPresent(cursorOffset uint32) bool {
	// Use end cursor as start cursor may be 0 even if field is in fact present
	return binary.LittleEndian.Uint32(c[cursorOffset+4:cursorOffset+8]) != 0
}

func FindCursors(
	cursorOffsets map[int32]uint32,
	payload []byte,
	cursors Cursors,
	fromOffset int,
	untilFieldNum int32,
) (int, error) {
	l := len(payload)
	i := fromOffset
	for i < l {
		// Read field tag
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow
			}
			if i >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := payload[i]
			i++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}

		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)

		// Store start cursor
		cursorOffset, fieldKnown := cursorOffsets[fieldNum]
		if fieldKnown {
			binary.LittleEndian.PutUint32(cursors[cursorOffset:cursorOffset+4], uint32(i))
		}

		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if i >= l {
					return 0, io.ErrUnexpectedEOF
				}
				i++
				if payload[i-1] < 0x80 {
					break
				}
			}
		case 1:
			i += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if i >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := payload[i]
				i++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLength
			}
			i += length
		case 5:
			i += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}

		// Store end cursor
		if fieldKnown {
			binary.LittleEndian.PutUint32(cursors[cursorOffset+4:cursorOffset+8], uint32(i))
		}

		if fieldNum == untilFieldNum {
			break
		}
	}

	return i, nil
}

func ParseBool(data []byte) bool {
	// Bools are represented as a varint, but we only bother to look at the first byte since the value should be either 0 or 1
	return len(data) >= 1 && data[0] == 1
}
