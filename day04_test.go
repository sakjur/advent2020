package advent2020

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDay4(t *testing.T) {
	fDemo, err := os.Open("testdata/day4_demo.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}
	fReal, err := os.Open("testdata/day4_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader       io.Reader
		fieldsExists int
		valid        int
	}{
		{
			reader:       fDemo,
			fieldsExists: 2,
			valid:        2,
		},
		{
			reader:       fReal,
			fieldsExists: 247,
			valid:        145,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.fieldsExists), func(t *testing.T) {
			documents := ScanDocuments(tc.reader)
			fieldsExists := 0
			valid := 0
			for _, document := range documents {
				if document.MissingFields() == nil {
					fieldsExists++
				}
				if document.Valid() {
					valid++
				}
			}

			if tc.fieldsExists != fieldsExists {
				t.Errorf("expected %d documents with the correct fields, found %d", tc.fieldsExists, fieldsExists)
			}
			if tc.valid != valid {
				t.Errorf("expected %d valid documents, found %d", tc.valid, valid)
			}
		})
	}
}
