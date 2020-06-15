[![Build Status](https://travis-ci.org/dsoprea/go-unicode-byteorder.svg?branch=master)](https://travis-ci.org/dsoprea/go-unicode-byteorder)
[![Coverage Status](https://coveralls.io/repos/github/dsoprea/go-unicode-byteorder/badge.svg?branch=master)](https://coveralls.io/github/dsoprea/go-unicode-byteorder?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/dsoprea/go-unicode-byteorder)](https://goreportcard.com/report/github.com/dsoprea/go-unicode-byteorder)
[![GoDoc](https://godoc.org/github.com/dsoprea/go-unicode-byteorder?status.svg)](https://godoc.org/github.com/dsoprea/go-unicode-byteorder)

# Overview

This project will return the encoding associated with a sequence of Unicode
byte-order-mark (BOM) bytes. This sequence is typically used to identify the
encoding of a document and/or the byte-order of a document encoded with a
multibyte encoding.

This only works for encodings that have an associated BOM byte-sequence. For
raw-text encoded in an encoding that supports a BOM, it should be found at the
very beginning of the document. If there is a data-structure where Unicode-
encoded text is embedded within it, the BOM bytes, if supported, may be found as
a field or property.

For more information, see the [Unicode website](1).

[1]: https://unicode.org/faq/utf_bom.html#bom4 "When a BOM is used, is it only in 16-bit Unicode text?"
