package v4parser_test

import (
	"testing"

	"github.com/uvalib/virgo4-parser/v4parser"
)

func convert(query string) (v4parser.SolrParser, string, error) {
	sp := v4parser.SolrParser{Debug: true}

	solr, err := v4parser.ConvertToSolrWithParser(&sp, query)

	return sp, solr, err
}

func expectSolrConversionSuccess(t *testing.T, query string, expected string) v4parser.SolrParser {
	sp, solr, err := convert(query)

	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", query, err.Error())
	}

	if solr != expected {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", query, expected, solr)
	}

	return sp
}

func expectSolrConversionFailure(t *testing.T, query string) {
	_, solr, err := convert(query)

	if err == nil {
		t.Errorf("%s converted to %s, but should have failed", query, solr)
	}
}

func expectCounts(t *testing.T, label string, values []string, want int) {
	got := len(values)

	if len(values) != want {
		t.Errorf("%s count fail. Expected %d, Actual: %d, List: %v", label, want, got, values)
	}
}

func TestSolrSimpleValid(t *testing.T) {
	q := "title: {bannanas}"
	e := `_query_:"{!edismax qf=$title_qf pf=$title_pf}(bannanas)"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrSpecialCharsValid(t *testing.T) {
	q := "title:{A = B}"
	e := `_query_:"{!edismax qf=$title_qf pf=$title_pf}(A = B)"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrIdentifierValid(t *testing.T) {
	q := `identifier:{35007007606860}`
	e := `_query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(35007007606860)"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrSimpleInvalid(t *testing.T) {
	q := "title: {bannanas} OR author: bad"

	expectSolrConversionFailure(t, q)
}

func TestSolrValid(t *testing.T) {
	q := `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	e := `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)")) OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrValidCounts(t *testing.T) {
	q := `( title: {pepperoni OR artichoke hearts} AND subject:{pizza} ) OR (subject:{calzone} AND (keyword:{italian} NOT author:{fieri}))`
	e := `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(pepperoni)" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(artichoke hearts)") AND _query_:"{!edismax qf=$subject_qf pf=$subject_pf}(pizza)")) OR ((_query_:"{!edismax qf=$subject_qf pf=$subject_pf}(calzone)" AND ((_query_:"{!edismax}(italian)" NOT _query_:"{!edismax qf=$author_qf pf=$author_pf}(fieri)")))))`

	sp := expectSolrConversionSuccess(t, q, e)

	expectCounts(t, "title", sp.Titles, 2)
	expectCounts(t, "author", sp.Authors, 1)
	expectCounts(t, "subject", sp.Subjects, 2)
	expectCounts(t, "keyword", sp.Keywords, 1)
}

func TestSolrDateSingle(t *testing.T) {
	q := `date:{1945}`
	e := `_query_:"{!lucene df=published_daterange}(1945)"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrDateRange(t *testing.T) {
	q := `date:{1945/12/07 TO 1949}`
	e := `_query_:"{!lucene df=published_daterange}([1945-12-07 TO 1949])"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrDateBefore(t *testing.T) {
	q := `date:{BEFORE 1945-12-06}`
	e := `_query_:"{!lucene df=published_daterange}([* TO 1945-12-06])"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrDateAfter(t *testing.T) {
	q := `date:{AFTER 1945}`
	e := `_query_:"{!lucene df=published_daterange}([1945 TO *])"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrDateMixed(t *testing.T) {
	q := `date:{<1945} AND date:{>1932} AND author:{Shelly}`
	e := `((_query_:"{!lucene df=published_daterange}([* TO 1945])" AND _query_:"{!lucene df=published_daterange}([1932 TO *])") AND _query_:"{!edismax qf=$author_qf pf=$author_pf}(Shelly)")`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrEmptyQuery(t *testing.T) {
	q := ``

	expectSolrConversionFailure(t, q)
}

func TestSolrEmptyKeyword(t *testing.T) {
	q := `keyword:{}`
	e := `_query_:"{!edismax}(*)"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrBobRandomQuery(t *testing.T) {
	q := `keyword:{cincinnati, ohio (home of the :reds:)}`
	e := `_query_:"{!edismax}(cincinnati, ohio [home of the :reds:])"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrSearchTip1(t *testing.T) {
	q := `keyword: {"grapes of wrath"}`
	e := `_query_:"{!edismax}(\"grapes of wrath\")"`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrSearchTip2(t *testing.T) {
	q := `keyword: {kyoto NOT protocol}`
	e := `(_query_:"{!edismax}(kyoto)" NOT _query_:"{!edismax}(protocol)")`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrSearchTip3(t *testing.T) {
	q := `keyword: {"frida kahlo" AND exhibitions}`
	e := `(_query_:"{!edismax}(\"frida kahlo\")" AND _query_:"{!edismax}(exhibitions)")`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrSearchTip4a(t *testing.T) {
	q := `keyword: {(calico OR "tortoise shell") AND cats}`
	e := `(_query_:"{!edismax}([[calico  OR  \"tortoise shell\"]])" AND _query_:"{!edismax}(cats)")`

	expectSolrConversionSuccess(t, q, e)
}

func TestSolrSearchTip4b(t *testing.T) {
	q := `(keyword: {calico OR "tortoise shell"})  AND keyword: {cats}`
	e := `(((_query_:"{!edismax}(calico)" OR _query_:"{!edismax}(\"tortoise shell\")")) AND _query_:"{!edismax}(cats)")`

	expectSolrConversionSuccess(t, q, e)
}
