package v4parser_test

import (
	"testing"

	"github.com/uvalib/virgo4-parser/v4parser"
)

type v4query struct {
	query   string
	timeout int
}

func (v *v4query) validate() (bool, string) {
	if v.timeout > 0 {
		return v4parser.ValidateWithTimeout(v.query, v.timeout)
	}

	return v4parser.Validate(v.query)
}

func (v *v4query) expectValidationSuccess(t *testing.T) {
	valid, errors := v.validate()

	if valid == false {
		t.Errorf("%s did not validate, but should have: %s", v.query, errors)
	}

	if errors != "" {
		t.Errorf("%s had errors, but should not have: %s", v.query, errors)
	}
}

func (v *v4query) expectValidationFailure(t *testing.T) {
	valid, errors := v.validate()

	if valid == true {
		t.Errorf("%s validated, but should not have: %s", v.query, errors)
	}

	if errors == "" {
		t.Errorf("%s did not have errors, but should have", v.query)
	}
}

func TestSimpleValid(t *testing.T) {
	v := v4query{query: "title: {bannanas}"}

	v.expectValidationSuccess(t)
}

func TestSpecialCharsValid(t *testing.T) {
	v := v4query{query: "title:{A = B}"}

	v.expectValidationSuccess(t)
}

func TestIdentifierValid(t *testing.T) {
	v := v4query{query: `identifier:{35007007606860  OR 9780754645733 OR 38083649 OR 2001020407  OR u5670758 OR "KJE5602.C73 2012"}`}

	v.expectValidationSuccess(t)
}

func TestIdentifierInvalid(t *testing.T) {
	v := v4query{query: `identifier:{35007007606860  OR 9780754645733 OR 38083649 OR 2001020407  OR u5670758 OR KJE5602.C73 2012"}`}

	v.expectValidationFailure(t)
}

func TestSimpleInvalidCrash(t *testing.T) {
	v := v4query{query: "title: bannanas"}

	v.expectValidationFailure(t)
}

func TestComplexValid(t *testing.T) {
	v := v4query{query: `( title : {"susan sontag" OR music title} AND keyword:{ Maunsell } ) OR author:{ liberty } AND subject:{ poe } AND keyword:{ horror }`}

	v.expectValidationSuccess(t)
}

func TestComplexInvalid(t *testing.T) {
	v := v4query{query: `( title : {"susan sontag" OR music title} AND keyword:{ Maunsell }  OR author:{ liberty } NOT author:{ poe }`}

	v.expectValidationFailure(t)
}

func TestNotOpValid(t *testing.T) {
	v := v4query{query: `author:{ edgar } NOT author:{ poe }`}

	v.expectValidationSuccess(t)
}

func TestNotOpInvalid(t *testing.T) {
	// Unary NOT operator is not supported
	v := v4query{query: `author:{ edgar } author:{ NOT poe }`}

	v.expectValidationFailure(t)
}

func TestEmptyKeyword(t *testing.T) {
	v := v4query{query: `keyword:{}`}

	v.expectValidationSuccess(t)
}

func TestEmptyQuotedKeyword(t *testing.T) {
	v := v4query{query: `keyword:{""}`}

	v.expectValidationFailure(t)
}

func TestMismatchedBracket(t *testing.T) {
	v := v4query{query: `keyword:{`}

	v.expectValidationFailure(t)
}

func TestExtraBracket(t *testing.T) {
	v := v4query{query: `keyword:{{}`}

	v.expectValidationFailure(t)
}

func TestUnsupportedType(t *testing.T) {
	v := v4query{query: `color:{blue}`}

	v.expectValidationFailure(t)
}

func TestStarQuery(t *testing.T) {
	v := v4query{query: `keyword:{"*"}`}

	v.expectValidationSuccess(t)
}

func TestStarQueryNoQuotes(t *testing.T) {
	v := v4query{query: `keyword:{*}`}

	v.expectValidationSuccess(t)
}

func TestBobRandomQuery(t *testing.T) {
	v := v4query{query: `keyword:{cincinnati, ohio (home of the :reds:)}`}

	v.expectValidationSuccess(t)
}

func TestSearchTip1(t *testing.T) {
	v := v4query{query: `keyword: {"grapes of wrath"}`}

	v.expectValidationSuccess(t)
}

func TestSearchTip2(t *testing.T) {
	v := v4query{query: `keyword: {kyoto NOT protocol}`}

	v.expectValidationSuccess(t)
}

func TestSearchTip3(t *testing.T) {
	v := v4query{query: `keyword: {"frida kahlo" AND exhibitions}`}

	v.expectValidationSuccess(t)
}

func TestSearchTip4a(t *testing.T) {
	v := v4query{query: `keyword: {(calico OR "tortoise shell") AND cats}`}

	v.expectValidationSuccess(t)
}

func TestSearchTip4b(t *testing.T) {
	v := v4query{query: `(keyword: {calico OR "tortoise shell"})  AND keyword: {cats}`}

	v.expectValidationSuccess(t)
}

func TestSlowValidation(t *testing.T) {
	v := v4query{query: `keyword: { I have often thought that nothing would do more extensive good at small expense than the establishment of a small circulating library in every county, to consist of a few well-chosen books, to be lent to the people of the country under regulations as would secure their safe return in due time. }`}

	v.expectValidationSuccess(t)
}

func TestSlowValidationWithTimeout(t *testing.T) {
	v := v4query{
		query:   `keyword: { I have often thought that nothing would do more extensive good at small expense than the establishment of a small circulating library in every county, to consist of a few well-chosen books, to be lent to the people of the country under regulations as would secure their safe return in due time. }`,
		timeout: 5,
	}

	v.expectValidationFailure(t)
}

func BenchmarkSlowValidation(b *testing.B) {
	v4parser.Validate(`keyword: { I have often thought that nothing would do more extensive good at small expense than the establishment of a small circulating library in every county, to consist of a few well-chosen books, to be lent to the people of the country under regulations as would secure their safe return in due time. }`)
}
