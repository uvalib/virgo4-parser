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
