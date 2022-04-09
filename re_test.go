package re_test

import (
	"testing"

	"github.com/mlampret/re"
)

var text1 = `Brown fox is a forest animal. Red fox is a desert animal.`

func TestSimpleInterface(t *testing.T) {
	// Matches
	m := re.Matches(text1, `fox.+animal`)

	if !m {
		t.Error("Matches not working correctly")
	}

	// Submatch
	s1 := re.Submatch(text1, `(\w+)\sfox`, 1)

	if s1 != "Brown" {
		t.Error("Submatch 1 not working correctly", s1)
	}

	s2 := re.Submatch(text1, `(Brown)\s(\w+)`, 2)

	if s2 != "fox" {
		t.Error("Submatch 2 not working correctly", s2)
	}

	s3 := re.Submatch(text1, `(\w+)\sfox`, 3)

	if s3 != "" {
		t.Error("Submatch 3 not working correctly", s3)
	}

	// Replace
	text2 := re.Replace(text1, `(\S+)\s(fox|cat)`, `$2 ($1)`)

	if text2 != "fox (Brown) is a forest animal. fox (Red) is a desert animal." {
		t.Error("Replace not working correctly", text2)
	}
}

func TestObjectOrientedInterface(t *testing.T) {
	// Matches
	m := re.String(text1).Pattern(`fox.+animal`).Matches()

	if !m {
		t.Error("Matches not working correctly")
	}

	// Submatch
	s1 := re.String(text1).Pattern(`(\w+)\sfox`).Submatch(1)

	if s1 != "Brown" {
		t.Error("Submatch 1 not working correctly", s1)
	}

	s2 := re.Bytes([]byte(text1)).Pattern(`(Brown)\s(\w+)`).Submatch(2)

	if s2 != "fox" {
		t.Error("Submatch 2 not working correctly", s2)
	}

	s3 := re.String(text1).Pattern(`(\w+)\sfox`).Submatch(3)
	if s3 != "" {
		t.Error("Submatch 3 not working correctly", s3)
	}

	// Replace
	text2 := re.String(text1).Pattern(`(\S+)\s(fox|cat)`).Replace(`$2 ($1)`)

	if text2 != "fox (Brown) is a forest animal. fox (Red) is a desert animal." {
		t.Error("Replace not working correctly", text2)
	}
}
