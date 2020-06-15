package bom

import (
	"fmt"
	"testing"

	"encoding/binary"

	"github.com/dsoprea/go-logging"
)

func TestGetEncoding_Hit_Utf8(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf8BomBytes)
	log.PanicIf(err)

	if encoding != Utf8Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Utf16_BigEndian(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf16BomBytesBigEndian)
	log.PanicIf(err)

	if encoding != Utf16Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != binary.BigEndian {
		t.Fatalf("Byte order not correct.")
	}
}

func ExampleGetEncoding_utf16_bigEndian() {
	encoding, byteOrder, err := GetEncoding([]byte{0xFE, 0xFF})
	log.PanicIf(err)

	fmt.Printf("%s\n", encoding)
	fmt.Printf("%s\n", byteOrder)

	// Output:
	// UTF-16
	// BigEndian
}

func ExampleGetEncoding_utf16_littleEndian() {
	encoding, byteOrder, err := GetEncoding([]byte{0xFF, 0xFE})
	log.PanicIf(err)

	fmt.Printf("%s\n", encoding)
	fmt.Printf("%s\n", byteOrder)

	// Output:
	// UTF-16
	// LittleEndian
}

func TestGetEncoding_Hit_Utf16_LittleEndian(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf16BomBytesLittleEndian)
	log.PanicIf(err)

	if encoding != Utf16Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != binary.LittleEndian {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Utf32_BigEndian(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf32BomBytesBigEndian)
	log.PanicIf(err)

	if encoding != Utf32Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != binary.BigEndian {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Utf32_LittleEndian(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf32BomBytesLittleEndian)
	log.PanicIf(err)

	if encoding != Utf32Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != binary.LittleEndian {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Utf7_1(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf7_1BomBytes)
	log.PanicIf(err)

	if encoding != Utf7Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Utf7_2(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf7_2BomBytes)
	log.PanicIf(err)

	if encoding != Utf7Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Utf7_3(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf7_3BomBytes)
	log.PanicIf(err)

	if encoding != Utf7Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Utf7_4(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf7_4BomBytes)
	log.PanicIf(err)

	if encoding != Utf7Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Utf7_5(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf7_5BomBytes)
	log.PanicIf(err)

	if encoding != Utf7Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Utf1(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Utf1BomBytes)
	log.PanicIf(err)

	if encoding != Utf1Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_UtfEbcdic(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(UtfEbcdicBomBytes)
	log.PanicIf(err)

	if encoding != UtfEbcdicEncoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Scsu(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(ScsuBomBytes)
	log.PanicIf(err)

	if encoding != ScsuEncoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Bocu1(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Bocu1BomBytes)
	log.PanicIf(err)

	if encoding != Bocu1Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Hit_Gb18030(t *testing.T) {
	encoding, byteOrder, err := GetEncoding(Gb18030BomBytes)
	log.PanicIf(err)

	if encoding != Gb18030Encoding {
		t.Fatalf("Encoding not correct: (%d)", encoding)
	} else if byteOrder != nil {
		t.Fatalf("Byte order not correct.")
	}
}

func TestGetEncoding_Miss(t *testing.T) {
	_, _, err := GetEncoding([]byte{1, 2, 3})
	if err != ErrNotValidEncoding {
		t.Fatalf("Expected encoding-not-valid error: %v", err)
	}
}

func TestEncoding_String_Utf8(t *testing.T) {
	if Utf8Encoding.String() != "UTF-8" {
		t.Fatalf("String not correct.")
	}
}

func TestEncoding_String_Utf16(t *testing.T) {
	if Utf16Encoding.String() != "UTF-16" {
		t.Fatalf("String not correct.")
	}
}

func TestEncoding_String_Utf32(t *testing.T) {
	if Utf32Encoding.String() != "UTF-32" {
		t.Fatalf("String not correct.")
	}
}

func TestEncoding_String_Utf7(t *testing.T) {
	if Utf7Encoding.String() != "UTF-7" {
		t.Fatalf("String not correct.")
	}
}

func TestEncoding_String_Utf1(t *testing.T) {
	if Utf1Encoding.String() != "UTF-1" {
		t.Fatalf("String not correct.")
	}
}

func TestEncoding_String_UtfEbcdic(t *testing.T) {
	if UtfEbcdicEncoding.String() != "UTF-EBCDIC" {
		t.Fatalf("String not correct.")
	}
}

func TestEncoding_String_Scsu(t *testing.T) {
	if ScsuEncoding.String() != "SCSU" {
		t.Fatalf("String not correct.")
	}
}

func TestEncoding_String_Bocu1(t *testing.T) {
	if Bocu1Encoding.String() != "BOCU-1" {
		t.Fatalf("String not correct.")
	}
}

func TestEncoding_String_Gb18030(t *testing.T) {
	if Gb18030Encoding.String() != "GB-18030" {
		t.Fatalf("String not correct.")
	}
}

func TestEncoding_String_Invalid(t *testing.T) {
	defer func() {
		errRaw := recover()
		if errRaw == nil {
			t.Fatalf("Expected failure for invalid encoding.")
		}

		err := errRaw.(error)
		if err.Error() != "encoding not valid" {
			log.Panic(err)
		}
	}()

	Encoding(99).String()
}
