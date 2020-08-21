package v4parser_test

import (
	"testing"

	"github.com/uvalib/virgo4-parser/v4parser"
)

func validate(query string) (bool, string) {
	return v4parser.Validate(query)
}

func expectValidationSuccess(t *testing.T, query string) {
	valid, errors := validate(query)

	if valid == false {
		t.Errorf("%s did not validate, but should have: %s", query, errors)
	}

	if errors != "" {
		t.Errorf("%s had errors, but should not have: %s", query, errors)
	}
}

func expectValidationFailure(t *testing.T, query string) {
	valid, errors := validate(query)

	if valid == true {
		t.Errorf("%s validated, but should not have: %s", query, errors)
	}

	if errors == "" {
		t.Errorf("%s did not have errors, but should have", query)
	}
}

func TestSimpleValid(t *testing.T) {
	q := "title: {bannanas}"

	expectValidationSuccess(t, q)
}

func TestSpecialCharsValid(t *testing.T) {
	q := "title:{A = B}"

	expectValidationSuccess(t, q)
}

func TestIdentifierValid(t *testing.T) {
	q := `identifier:{35007007606860  OR 9780754645733 OR 38083649 OR 2001020407  OR u5670758 OR "KJE5602.C73 2012"}`

	expectValidationSuccess(t, q)
}

func TestIdentifierInvalid(t *testing.T) {
	q := `identifier:{35007007606860  OR 9780754645733 OR 38083649 OR 2001020407  OR u5670758 OR KJE5602.C73 2012"}`

	expectValidationFailure(t, q)
}

func TestSimpleInvalidCrash(t *testing.T) {
	q := "title: bannanas"

	expectValidationFailure(t, q)
}

func TestComplexValid(t *testing.T) {
	q := `( title : {"susan sontag" OR music title} AND keyword:{ Maunsell } ) OR author:{ liberty } AND subject:{ poe } AND keyword:{ horror }`

	expectValidationSuccess(t, q)
}

func TestComplexInvalid(t *testing.T) {
	q := `( title : {"susan sontag" OR music title} AND keyword:{ Maunsell }  OR author:{ liberty } NOT author:{ poe }`

	expectValidationFailure(t, q)
}

func TestNotOpValid(t *testing.T) {
	q := `author:{ edgar } NOT author:{ poe }`

	expectValidationSuccess(t, q)
}

func TestNotOpInvalid(t *testing.T) {
	// Unary NOT operator is not supported
	q := `author:{ edgar } author:{ NOT poe }`

	expectValidationFailure(t, q)
}

func TestEmptyKeyword(t *testing.T) {
	q := `keyword:{}`

	expectValidationSuccess(t, q)
}

func TestEmptyQuotedKeyword(t *testing.T) {
	q := `keyword:{""}`

	expectValidationFailure(t, q)
}

func TestMismatchedBracket(t *testing.T) {
	q := `keyword:{`

	expectValidationFailure(t, q)
}

func TestExtraBracket(t *testing.T) {
	q := `keyword:{{}`

	expectValidationFailure(t, q)
}

func TestUnsupportedType(t *testing.T) {
	q := `color:{blue}`

	expectValidationFailure(t, q)
}

func TestStarQuery(t *testing.T) {
	q := `keyword:{"*"}`

	expectValidationSuccess(t, q)
}

func TestStarQueryNoQuotes(t *testing.T) {
	q := `keyword:{*}`

	expectValidationSuccess(t, q)
}

func TestBobRandomQuery(t *testing.T) {
	q := `keyword:{cincinnati, ohio (home of the :reds:)}`

	expectValidationSuccess(t, q)
}

func TestSearchTip1(t *testing.T) {
	q := `keyword: {"grapes of wrath"}`

	expectValidationSuccess(t, q)
}

func TestSearchTip2(t *testing.T) {
	q := `keyword: {kyoto NOT protocol}`

	expectValidationSuccess(t, q)
}

func TestSearchTip3(t *testing.T) {
	q := `keyword: {"frida kahlo" AND exhibitions}`

	expectValidationSuccess(t, q)
}

func TestSearchTip4a(t *testing.T) {
	q := `keyword: {(calico OR "tortoise shell") AND cats}`

	expectValidationSuccess(t, q)
}

func TestSearchTip4b(t *testing.T) {
	q := `(keyword: {calico OR "tortoise shell"})  AND keyword: {cats}`

	expectValidationSuccess(t, q)
}

func BenchmarkSlowValidation(b *testing.B) {
	v4parser.Validate(`keyword: { I have often thought that nothing would do more extensive good at small expense than the establishment of a small circulating library in every county, to consist of a few well-chosen books, to be lent to the people of the country under regulations as would secure their safe return in due time. }`)
}
