package v4parser_test

import (
	"testing"

	"github.com/uvalib/virgo4-parser/v4parser"
)

func TestSimpleValid(t *testing.T) {
	q := "title: {bannanas}"
	valid, errors := v4parser.Validate(q)
	if valid == false {
		t.Errorf("%s did not validate, but should have: %s", q, errors)
	}
}

func TestSpecialCharsValid(t *testing.T) {
	q := "title:{A = B}"
	valid, errors := v4parser.Validate(q)
	if valid == false {
		t.Errorf("%s did not validate, but should have: %s", q, errors)
	}
}

func TestIdentifierValid(t *testing.T) {
	q := `identifier:{35007007606860  OR 9780754645733 OR 38083649 OR 2001020407  OR u5670758 OR "KJE5602.C73 2012"}`
	valid, errors := v4parser.Validate(q)
	if valid == false {
		t.Errorf("%s did not validate, but should have: %s", q, errors)
	}
}

func TestIdentifierInvalid(t *testing.T) {
	q := `identifier:{35007007606860  OR 9780754645733 OR 38083649 OR 2001020407  OR u5670758 OR KJE5602.C73 2012"}`
	valid, _ := v4parser.Validate(q)
	if valid == true {
		t.Errorf("%s validate, but should not have", q)
	}
}

func TestSimpleInvalidCrash(t *testing.T) {
	q := "title: bannanas"
	valid, _ := v4parser.Validate(q)
	if valid == true {
		t.Errorf("%s validate, but should not have", q)
	}
}

func TestComplexValid(t *testing.T) {
	q := `( title : {"susan sontag" OR music title} AND keyword:{ Maunsell } ) OR author:{ liberty } AND subject:{ poe } AND keyword:{ horror }`
	valid, errors := v4parser.Validate(q)
	if valid == false {
		t.Errorf("%s did not validate, but should have: %s", q, errors)
	}
}

func TestComplexInvalid(t *testing.T) {
	q := `( title : {"susan sontag" OR music title} AND keyword:{ Maunsell }  OR author:{ liberty } NOT author:{ poe }`
	valid, _ := v4parser.Validate(q)
	if valid == true {
		t.Errorf("%s validate, but should not have", q)
	}
}

func TestNotOpValid(t *testing.T) {
	q := `author:{ edgar } NOT author:{ poe }`
	valid, errors := v4parser.Validate(q)
	if valid == false {
		t.Errorf("%s did not validate, but should have: %s", q, errors)
	}
}

func TestNotOpInvalid(t *testing.T) {
	// Unary NOT operator is not supported
	q := `author:{ edgar } author:{ NOT poe }`
	valid, _ := v4parser.Validate(q)
	if valid == true {
		t.Errorf("%s validatef, but should not have", q)
	}
}
