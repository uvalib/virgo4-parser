package v4parser_test

import (
	"testing"

	"github.com/uvalib/virgo4-parser/v4parser"
)

func TestSolrSimpleValid(t *testing.T) {
	q := "title: {bannanas}"
	solr, err := v4parser.ConvertToSolr(q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	expected := `_query_:"{!edismax qf=$title_qf pf=$title_pf}(bannanas)"`
	if solr != expected {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", q, expected, solr)
	}
}

func TestSolrSpecialCharsValid(t *testing.T) {
	q := "title:{A = B}"
	expect := `_query_:"{!edismax qf=$title_qf pf=$title_pf}(A = B)"`

	solr, err := v4parser.ConvertToSolr(q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	if solr != expect {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", q, expect, solr)
	}
}

func TestSolrIdentifierValid(t *testing.T) {
	q := `identifier:{35007007606860}`
	expect := `_query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(35007007606860)"`

	solr, err := v4parser.ConvertToSolr(q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	if solr != expect {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", q, expect, solr)
	}
}

func TestSolrSimpleInvalid(t *testing.T) {
	q := "title: {bannanas} OR author: bad"
	solr, err := v4parser.ConvertToSolr(q)
	if err == nil {
		t.Errorf("%s converted to %s, but should have failed", q, solr)
	}
}

func TestSolrValid(t *testing.T) {
	q := `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`
	solr, err := v4parser.ConvertToSolr(q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	expected := `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)")) OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")`
	if solr != expected {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", q, expected, solr)
	}
}

func TestSolrValidCounts(t *testing.T) {
	q := `( title: {pepperoni OR artichoke hearts} AND subject:{pizza} ) OR (subject:{calzone} AND (keyword:{italian} NOT author:{fieri}))`
	sp := v4parser.SolrParser{}
	_, err := v4parser.ConvertToSolrWithParser(&sp, q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	tcnt := 2
	acnt := 1
	scnt := 2
	kcnt := 1
	if len(sp.Titles) != tcnt {
		t.Errorf("%s title count fail. Expected %d, Actual: %d, List: %v", q, tcnt, len(sp.Titles), sp.Titles)
	}
	if len(sp.Authors) != acnt {
		t.Errorf("%s author count fail. Expected %d, Actual: %d, List: %v", q, acnt, len(sp.Authors), sp.Authors)
	}
	if len(sp.Subjects) != scnt {
		t.Errorf("%s subject count fail. Expected %d, Actual: %d, List: %v", q, scnt, len(sp.Subjects), sp.Subjects)
	}
	if len(sp.Keywords) != kcnt {
		t.Errorf("%s keyword count fail. Expected %d, Actual: %d, List: %v", q, kcnt, len(sp.Keywords), sp.Keywords)
	}
}

func TestSolrDateSingle(t *testing.T) {
	q := `date:{1945}`
	expect := `_query_:"{!lucene df=published_daterange}(1945)"`

	solr, err := v4parser.ConvertToSolr(q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	if solr != expect {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", q, expect, solr)
	}
}

func TestSolrDateRange(t *testing.T) {
	q := `date:{1945/12/07 TO 1949}`
	expect := `_query_:"{!lucene df=published_daterange}([1945-12-07 TO 1949])"`

	solr, err := v4parser.ConvertToSolr(q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	if solr != expect {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", q, expect, solr)
	}
}

func TestSolrDateBefore(t *testing.T) {
	q := `date:{BEFORE 1945-12-06}`
	expect := `_query_:"{!lucene df=published_daterange}([* TO 1945-12-06])"`

	solr, err := v4parser.ConvertToSolr(q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	if solr != expect {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", q, expect, solr)
	}
}

func TestSolrDateAfter(t *testing.T) {
	q := `date:{AFTER 1945}`
	expect := `_query_:"{!lucene df=published_daterange}([1945 TO *])"`

	solr, err := v4parser.ConvertToSolr(q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	if solr != expect {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", q, expect, solr)
	}
}

func TestSolrDateMixed(t *testing.T) {
	q := `date:{<1945} AND date:{>1932} AND author:{Shelly}`
	expect := `((_query_:"{!lucene df=published_daterange}([* TO 1945])" AND _query_:"{!lucene df=published_daterange}([1932 TO *])") AND _query_:"{!edismax qf=$author_qf pf=$author_pf}(Shelly)")`

	solr, err := v4parser.ConvertToSolr(q)
	if err != nil {
		t.Errorf("%s couldn't convert, but should have: %s", q, err.Error())
	}
	if solr != expect {
		t.Errorf("%s convert fail. Expected %s, Actual: %s", q, expect, solr)
	}
}
