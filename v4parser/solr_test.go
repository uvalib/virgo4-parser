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

func expectSolrCoversionSuccess(t *testing.T, query string, expected string) v4parser.SolrParser {
	sp, solr, err := convert(query)

	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", query, err.Error())
	}

	if solr != expected {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", query, expected, solr)
	}

	return sp
}

func expectSolrCoversionFailure(t *testing.T, query string) {
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

	expectSolrCoversionSuccess(t, q, e)
}

func TestSolrSpecialCharsValid(t *testing.T) {
	q := "title:{A = B}"
	e := `_query_:"{!edismax qf=$title_qf pf=$title_pf}(A = B)"`

	expectSolrCoversionSuccess(t, q, e)
}

func TestSolrIdentifierValid(t *testing.T) {
	q := `identifier:{35007007606860}`
	e := `_query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(35007007606860)"`

	expectSolrCoversionSuccess(t, q, e)
}

func TestSolrSimpleInvalid(t *testing.T) {
	q := "title: {bannanas} OR author: bad"

	expectSolrCoversionFailure(t, q)
}

func TestSolrValid(t *testing.T) {
	q := `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	e := `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)")) OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")`

	expectSolrCoversionSuccess(t, q, e)
}

func TestSolrValidCounts(t *testing.T) {
	q := `( title: {pepperoni OR artichoke hearts} AND subject:{pizza} ) OR (subject:{calzone} AND (keyword:{italian} NOT author:{fieri}))`
	e := `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(pepperoni)" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(artichoke hearts)") AND _query_:"{!edismax qf=$subject_qf pf=$subject_pf}(pizza)")) OR ((_query_:"{!edismax qf=$subject_qf pf=$subject_pf}(calzone)" AND ((_query_:"{!edismax}(italian)" NOT _query_:"{!edismax qf=$author_qf pf=$author_pf}(fieri)")))))`

	sp := expectSolrCoversionSuccess(t, q, e)

	expectCounts(t, "title", sp.Titles, 2)
	expectCounts(t, "author", sp.Authors, 1)
	expectCounts(t, "subject", sp.Subjects, 2)
	expectCounts(t, "keyword", sp.Keywords, 1)
}

func TestSolrDateSingle(t *testing.T) {
	q := `date:{1945}`
	e := `_query_:"{!lucene df=published_daterange}(1945)"`

	expectSolrCoversionSuccess(t, q, e)
}

func TestSolrDateRange(t *testing.T) {
	q := `date:{1945/12/07 TO 1949}`
	e := `_query_:"{!lucene df=published_daterange}([1945-12-07 TO 1949])"`

	expectSolrCoversionSuccess(t, q, e)
}

func TestSolrDateBefore(t *testing.T) {
	q := `date:{BEFORE 1945-12-06}`
	e := `_query_:"{!lucene df=published_daterange}([* TO 1945-12-06])"`

	expectSolrCoversionSuccess(t, q, e)
}

func TestSolrDateAfter(t *testing.T) {
	q := `date:{AFTER 1945}`
	e := `_query_:"{!lucene df=published_daterange}([1945 TO *])"`

	expectSolrCoversionSuccess(t, q, e)
}

func TestSolrDateMixed(t *testing.T) {
	q := `date:{<1945} AND date:{>1932} AND author:{Shelly}`
	e := `((_query_:"{!lucene df=published_daterange}([* TO 1945])" AND _query_:"{!lucene df=published_daterange}([1932 TO *])") AND _query_:"{!edismax qf=$author_qf pf=$author_pf}(Shelly)")`

	expectSolrCoversionSuccess(t, q, e)
}

func TestSolrEmptyQuery(t *testing.T) {
	q := ``

	expectSolrCoversionFailure(t, q)
}
