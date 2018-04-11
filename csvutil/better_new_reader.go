package csvutil

import (
	"bufio"
	"encoding/csv"
	"io"
)

// BetterNewReader returns a new Reader that reads from r, with some
// additional intermediatiary operations.
func BetterNewReader(r io.Reader) *csv.Reader {

  // Replace \r carriage returns with \n newlines.
  bnr := csv.NewReader(ReplaceSoloCarriageReturns(r))
  return bnr
}

// ReplaceSoloCarriageReturns wraps an io.Reader; for every call of Read,
// replacing instances of lonely \r with \r\n. Lots of files in the wild will
// come without "proper" line breaks, which irritates go's standard csv package.
// Use as, e.g.:
// rdr, err := csv.NewReader(ReplaceSoloCarriageReturns(r))
func ReplaceSoloCarriageReturns(data io.Reader) io.Reader {
	return crlfReplaceReader{
		rdr: bufio.NewReader(data),
	}
}

// crlfReplaceReader wraps a reader. (A user provided Read method is required.)
type crlfReplaceReader struct {
	rdr *bufio.Reader
}

// This Read method implements io.Reader required for type crlfReplaceReader.
func (c crlfReplaceReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}

	for {
		if n == len(p) {
			return
		}

		p[n], err = c.rdr.ReadByte()
		if err != nil {
			return
		}

		// Any time we encounter \r & still have space, check to see if \n follows.
		// If next char is not \n, add it in manually.
		if p[n] == '\r' && n < len(p) {
			if pk, err := c.rdr.Peek(1); (err == nil && pk[0] != '\n') || (err != nil && err.Error() == io.EOF.Error()) {
				n++
				p[n] = '\n'
			}
		}

		n++
	}
	return
}
