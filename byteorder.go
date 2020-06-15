package bom

import (
	"bytes"
	"errors"

	"encoding/binary"

	"github.com/dsoprea/go-logging"
)

var (
	// ErrNotValidEncoding is returned if the bytes are not valid BOM-bytes for
	// any encoding.
	ErrNotValidEncoding = errors.New("not a valid encoding")
)

// Encoding describes an encoding-type determined by a BOM.
type Encoding int

const (
	// Utf8Encoding represents the UTF-8 encoding. This starts at one so an
	// accidental zero value is not accidentally interpreted as a valid
	// encoding.
	Utf8Encoding Encoding = iota + 1

	// Utf16Encoding represents the UTF-16 encoding.
	Utf16Encoding Encoding = iota + 1

	// Utf32Encoding represents the UTF-32 encoding.
	Utf32Encoding Encoding = iota + 1

	// Utf7Encoding represents the UTF-7 encoding.
	Utf7Encoding Encoding = iota + 1

	// Utf1Encoding represents the UTF-1 encoding.
	Utf1Encoding Encoding = iota + 1

	// UtfEbcdicEncoding represents the UTF-EBCDIC encoding.
	UtfEbcdicEncoding Encoding = iota + 1

	// ScsuEncoding represents the SCSU encoding.
	ScsuEncoding Encoding = iota + 1

	// Bocu1Encoding represents the BOCU-1 encoding.
	Bocu1Encoding Encoding = iota + 1

	// Gb18030Encoding represents the GB-18030 encoding.
	Gb18030Encoding Encoding = iota + 1
)

// String returns a string describing the encoding.
func (encoding Encoding) String() string {
	switch encoding {
	case Utf8Encoding:
		return "UTF-8"
	case Utf16Encoding:
		return "UTF-16"
	case Utf32Encoding:
		return "UTF-32"
	case Utf7Encoding:
		return "UTF-7"
	case Utf1Encoding:
		return "UTF-1"
	case UtfEbcdicEncoding:
		return "UTF-EBCDIC"
	case ScsuEncoding:
		return "SCSU"
	case Bocu1Encoding:
		return "BOCU-1"
	case Gb18030Encoding:
		return "GB-18030"
	}

	log.Panicf("encoding not valid")
	panic(nil)
}

var (
	// Utf8BomBytes represents the UTF-8 BOM bytes.
	Utf8BomBytes = []byte{0xEF, 0xBB, 0xBF}

	// Utf16BomBytesBigEndian represents the UTF-16 BOM bytes if in the BE byte-
	// order.
	Utf16BomBytesBigEndian = []byte{0xFE, 0xFF}

	// Utf16BomBytesLittleEndian represents the UTF-16 BOM bytes if in the LE byte-
	// order.
	Utf16BomBytesLittleEndian = []byte{0xFF, 0xFE}

	// Utf32BomBytesBigEndian represents the UTF-32 BOM bytes if in the BE byte-
	// order.
	Utf32BomBytesBigEndian = []byte{0, 0, 0xFE, 0xFF}

	// Utf32BomBytesLittleEndian represents the UTF-32 BOM bytes if in the LE
	// byte-order.
	Utf32BomBytesLittleEndian = []byte{0xFE, 0xFF, 0, 0}

	// Utf7_1BomBytes represents the UTF-7 BOM bytes (sequence 1).
	Utf7_1BomBytes = []byte{0x2B, 0x2F, 0x76, 0x38}

	// Utf7_2BomBytes represents the UTF-7 BOM bytes (sequence 1).
	Utf7_2BomBytes = []byte{0x2B, 0x2F, 0x76, 0x39}

	// Utf7_3BomBytes represents the UTF-7 BOM bytes (sequence 1).
	Utf7_3BomBytes = []byte{0x2B, 0x2F, 0x76, 0x2B}

	// Utf7_4BomBytes represents the UTF-7 BOM bytes (sequence 1).
	Utf7_4BomBytes = []byte{0x2B, 0x2F, 0x76, 0x2F}

	// Utf7_5BomBytes represents the UTF-7 BOM bytes (sequence 1).
	Utf7_5BomBytes = []byte{0x2B, 0x2F, 0x76, 0x38, 0x2D}

	// Utf1BomBytes represents the UTF-1 BOM bytes.
	Utf1BomBytes = []byte{0xF7, 0x64, 0x4C}

	// UtfEbcdicBomBytes represents the UTF-EBCDIC BOM bytes.
	UtfEbcdicBomBytes = []byte{0xDD, 0x73, 0x66, 0x73}

	// ScsuBomBytes represents the SCSU BOM bytes.
	ScsuBomBytes = []byte{0x0E, 0xFE, 0xFF}

	// Bocu1BomBytes represents the BOCU-1 BOM bytes.
	Bocu1BomBytes = []byte{0xFB, 0xEE, 0x28}

	// Gb18030BomBytes represents the GB-18030 BOM bytes.
	Gb18030BomBytes = []byte{0x84, 0x31, 0x95, 0x33}
)

func getEncodingWihtByteOrder(s []byte) (encoding Encoding, byteOrder binary.ByteOrder, err error) {
	if bytes.Equal(s, Utf16BomBytesBigEndian) == true {
		return Utf16Encoding, binary.BigEndian, nil
	} else if bytes.Equal(s, Utf16BomBytesLittleEndian) == true {
		return Utf16Encoding, binary.LittleEndian, nil
	} else if bytes.Equal(s, Utf32BomBytesBigEndian) == true {
		return Utf32Encoding, binary.BigEndian, nil
	} else if bytes.Equal(s, Utf32BomBytesLittleEndian) == true {
		return Utf32Encoding, binary.LittleEndian, nil
	}

	return 0, nil, ErrNotValidEncoding
}

// GetEncoding match the bytes against known BOM sequences, and return the
// encoding and byte-order they represent. If not a known sequence of bytes,
// returns ErrNotValidEncoding.
func GetEncoding(s []byte) (encoding Encoding, byteOrder binary.ByteOrder, err error) {
	// We dispatch to a second method just to break up this function and
	// eliminate the cyclomatic complexity complaints.
	encoding, byteOrder, err = getEncodingWihtByteOrder(s)
	if err == nil {
		return encoding, byteOrder, nil
	}

	if bytes.Equal(s, Utf8BomBytes) == true {
		return Utf8Encoding, nil, nil
	} else if bytes.Equal(s, Utf7_1BomBytes) == true {
		return Utf7Encoding, nil, nil
	} else if bytes.Equal(s, Utf7_2BomBytes) == true {
		return Utf7Encoding, nil, nil
	} else if bytes.Equal(s, Utf7_3BomBytes) == true {
		return Utf7Encoding, nil, nil
	} else if bytes.Equal(s, Utf7_4BomBytes) == true {
		return Utf7Encoding, nil, nil
	} else if bytes.Equal(s, Utf7_5BomBytes) == true {
		return Utf7Encoding, nil, nil
	} else if bytes.Equal(s, Utf1BomBytes) == true {
		return Utf1Encoding, nil, nil
	} else if bytes.Equal(s, UtfEbcdicBomBytes) == true {
		return UtfEbcdicEncoding, nil, nil
	} else if bytes.Equal(s, ScsuBomBytes) == true {
		return ScsuEncoding, nil, nil
	} else if bytes.Equal(s, Bocu1BomBytes) == true {
		return Bocu1Encoding, nil, nil
	} else if bytes.Equal(s, Gb18030BomBytes) == true {
		return Gb18030Encoding, nil, nil
	}

	return 0, nil, ErrNotValidEncoding
}
