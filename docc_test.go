package docc

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var want = []string{
	"Title",
	"Subtitle",
	"Here is a first row.",
	"Here is a second row.",
}

func FatalCloser(t *testing.T, r io.Closer) {
	err := r.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadAll(t *testing.T) {
	fp := filepath.Clean("./testdata/test.docx")
	r, err := NewReader(fp)
	if err != nil {
		t.Fatal(err)
	}
	defer FatalCloser(t, r)
	got, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestReadAllFromReader(t *testing.T) {
	fp := filepath.Clean("./testdata/test.docx")
	f, err := os.Open(fp)
	if err != nil {
		t.Fatal(err)
	}
	r, err := NewFromReader(f)
	if err != nil {
		t.Fatal(err)
	}
	defer FatalCloser(t, r)
	got, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestRead(t *testing.T) {
	fp := filepath.Clean("./testdata/test.docx")
	r, err := NewReader(fp)
	if err != nil {
		t.Fatal(err)
	}
	defer FatalCloser(t, r)
	got, err := r.Read()
	if err != nil {
		t.Fatal(err)
	}
	if want[0] != got {
		t.Errorf("want %s, got %s", want[0], got)
	}
}
