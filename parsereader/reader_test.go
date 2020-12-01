package parsereader

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

const (
	basic = `test = 22

[h]
age = hoi

[a]
a=1

#
`
)

func TestEqualRW(t *testing.T) {
	var err1, err2 error

	file, err2 := os.Open("test.ini")
	if err2 != nil {
		t.Error(err2)
	}

	defer file.Close()

	reader1 := FromReadSeeker(strings.NewReader(basic))
	reader2 := FromReadSeeker(file)

	var r1, r2 rune

	runes1 := make([]rune, 0)
	runes2 := make([]rune, 0)

	i := 0

	for err1 == nil && err2 == nil {
		r1, _, err1 = reader1.ReadRune()
		r2, _, err2 = reader2.ReadRune()

		runes1 = append(runes1, r1)
		runes2 = append(runes2, r2)

		if r1 != r2 {
			t.Error(fmt.Errorf("runes not equal at %d: %d, %d", i, r1, r2))
		}

		i++
	}

	if !errors.Is(err1, io.EOF) {
		t.Error(fmt.Errorf("error 1 not eof: %w", err1))
	} else if !errors.Is(err2, io.EOF) {
		t.Error(fmt.Errorf("error 2 not eof: %w", err2))
	} else {
		t.Log("EOF: OK")
	}

	t.Log("READER1:", string(runes1))
	t.Log("READER2:", string(runes2))
}

func TestStringReader(t *testing.T) {
	var err error
	file := strings.NewReader(basic)

	parseReader := FromReadSeeker(file)

	var r rune

	runes := make([]rune, 0)

	for err == nil {
		r, _, err = parseReader.ReadRune()
		runes = append(runes, r)
	}

	if !errors.Is(err, io.EOF) {
		t.Error(err)
	} else {
		t.Log("EOF: OK")
	}

	t.Log(string(runes))
}

func TestFileReader(t *testing.T) {
	var err error

	file, err := os.Open("test.ini")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	parseReader := FromReadSeeker(file)

	var r rune

	runes := make([]rune, 0)

	for err == nil {
		r, _, err = parseReader.ReadRune()
		runes = append(runes, r)
	}

	if !errors.Is(err, io.EOF) {
		t.Error(err)
	} else {
		t.Log("EOF: OK")
	}

	t.Log(string(runes))
}
