package v4parser_test

import (
	"testing"

	"github.com/uvalib/virgo4-parser/v4parser"
)

func convert(t *testing.T, query string) (v4parser.SolrParser, string, error) {
	sp := v4parser.SolrParser{Debug: true}

	t.Logf("QUERY: %s", query)

	solr, err := v4parser.ConvertToSolrWithParser(&sp, query)

	return sp, solr, err
}

func expectSolrConversionSuccess(t *testing.T, query string, expected string) v4parser.SolrParser {
	sp, solr, err := convert(t, query)

	if err != nil {
		t.Errorf("Conversion failed, but should have succeeded:")
		t.Errorf("ERROR: %s", err.Error())
		t.Errorf("WANT : %s", expected)

		return sp
	}

	if solr != expected {
		t.Errorf("Conversion succeeded, but result was not as expected:")
		t.Errorf("GOT  : %s", solr)
		t.Errorf("WANT : %s", expected)
	}

	return sp
}

func expectSolrConversionFailure(t *testing.T, query string) {
	_, solr, err := convert(t, query)

	if err == nil {
		t.Errorf("Conversion succeeded, but should have failed:")
		t.Errorf("QUERY: %s", query)
		t.Errorf("GOT  : %s", solr)
	}
}

func expectCounts(t *testing.T, label string, values []string, want int) {
	got := len(values)

	if len(values) != want {
		t.Errorf("%s count fail. Expected %d, Actual: %d, List: %v", label, want, got, values)
	}
}

func TestSolrShouldSucceed(t *testing.T) {
	type expectedResults struct {
		query string // v4 query
		solr string // expected solr conversion
		counts bool // check counts?
		titles int // expected title count
		authors int // expected author count
		subjects int // expected subject count
		keywords int // expected keyword count
	}

	tests := []expectedResults {
		{
			query: `keyword:{}`,
			solr: `_query_:"{!edismax}(*)"`,
		},
		{
			query: `title: {bannanas}`,
			solr: `_query_:"{!edismax qf=$title_qf pf=$title_pf}(bannanas)"`,
		},
		{
			query: `title:{A = B}`,
			solr: `_query_:"{!edismax qf=$title_qf pf=$title_pf}(A = B)"`,
		},
		{
			query: `keyword:{digby OR duncan}`,
			solr: `(_query_:"{!edismax}(digby)" OR _query_:"{!edismax}(duncan)")`,
		},
		{
			query: `keyword:{digby AND duncan}`,
			solr: `(_query_:"{!edismax}(digby)" AND _query_:"{!edismax}(duncan)")`,
		},
		{
			query: `author:{Zhongguo Zang xue chu ban she  (publisher)}`,
			solr: `_query_:"{!edismax qf=$author_qf pf=$author_pf}(Zhongguo Zang xue chu ban she publisher)"`,
		},
		{
			query: `author:{ ʼJam-dbyangs-nyi-ma }`,
			solr: `_query_:"{!edismax qf=$author_qf pf=$author_pf}(ʼJam-dbyangs-nyi-ma)"`,
		},
		{
			query: `title:{Tragicheskai͡a istorii͡a kavkazskikh }`,
			solr: `_query_:"{!edismax qf=$title_qf pf=$title_pf}(Tragicheskai͡a istorii͡a kavkazskikh)"`,
		},
		{
			query: `author:{Немирович-Данченко, Василий Иванович}`,
			solr: `_query_:"{!edismax qf=$author_qf pf=$author_pf}(Немирович-Данченко, Василий Иванович)"`,
		},
		{
			query: `title : {"susan sontag" OR (music title)}   AND keyword:{ Maunsell }`,
			solr: `((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)")`,
		},
		{
			query: `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`,
			solr: `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)")) OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")`,
		},
		{
			query: `( title : {"susan sontag" OR (music title)}   AND keyword:{ Maunsell } ) OR author:{ liberty }`,
			solr: `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)")) OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")`,
		},
		{
			query: `title : {"susan sontag" OR (music title)}   AND ( keyword:{ Maunsell } OR author:{ liberty })`,
			solr: `((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND ((_query_:"{!edismax}(Maunsell)" OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")))`,
		},
		{
			query: ` author: {lovecraft} AND title: {madness NOT dunwich}`,
			solr: `(_query_:"{!edismax qf=$author_qf pf=$author_pf}(lovecraft)" AND (_query_:"{!edismax qf=$title_qf pf=$title_pf}(madness)" NOT _query_:"{!edismax qf=$title_qf pf=$title_pf}(dunwich)"))`,
		},
		{
			query: `identifier:{35007007606860  OR 9780754645733 OR 38083649 OR 2001020407  OR u5670758 OR "KJE5602.C73 2012"}`,
			solr: `(((((_query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(35007007606860)" OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(9780754645733)") OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(38083649)") OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(2001020407)") OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(u5670758)") OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(\"KJE5602.C73 2012\")")`,
		},
		{
			query: `keyword:{triceratops OR flameproof}`,
			solr: `(_query_:"{!edismax}(triceratops)" OR _query_:"{!edismax}(flameproof)")`,
		},
		{
			query: `identifier:{35007007606860}`,
			solr: `_query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(35007007606860)"`,
		},
		{
			query: `( title: {pepperoni OR artichoke hearts} AND subject:{pizza} ) OR (subject:{calzone} AND (keyword:{italian} NOT author:{fieri}))`,
			solr: `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(pepperoni)" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(artichoke hearts)") AND _query_:"{!edismax qf=$subject_qf pf=$subject_pf}(pizza)")) OR ((_query_:"{!edismax qf=$subject_qf pf=$subject_pf}(calzone)" AND ((_query_:"{!edismax}(italian)" NOT _query_:"{!edismax qf=$author_qf pf=$author_pf}(fieri)")))))`,
			counts: true,
			titles: 2,
			authors: 1,
			subjects: 2,
			keywords: 1,
		},
		{
			query: `date:{1945}`,
			solr: `_query_:"{!lucene df=published_daterange}(1945)"`,
		},
		{
			query: `date:{1945/12/07 TO 1949}`,
			solr: `_query_:"{!lucene df=published_daterange}([1945-12-07 TO 1949])"`,
		},
		{
			query: `date:{BEFORE 1945-12-06}`,
			solr: `_query_:"{!lucene df=published_daterange}([* TO 1945-12-06])"`,
		},
		{
			query: `date:{AFTER 1945}`,
			solr: `_query_:"{!lucene df=published_daterange}([1945 TO *])"`,
		},
		{
			query: `date:{<1945} AND date:{>1932} AND author:{Shelly}`,
			solr: `((_query_:"{!lucene df=published_daterange}([* TO 1945])" AND _query_:"{!lucene df=published_daterange}([1932 TO *])") AND _query_:"{!edismax qf=$author_qf pf=$author_pf}(Shelly)")`,
		},
		{
			query: `keyword:{cincinnati, ohio (home of the :reds:)}`,
			solr: `_query_:"{!edismax}(cincinnati, ohio home of the :reds:)"`,
		},
		{
			query: `keyword: {"grapes of wrath"}`,
			solr: `_query_:"{!edismax}(\"grapes of wrath\")"`,
		},
		{
			query: `keyword: {kyoto NOT protocol}`,
			solr: `(_query_:"{!edismax}(kyoto)" NOT _query_:"{!edismax}(protocol)")`,
		},
		{
			query: `keyword: {"frida kahlo" AND exhibitions}`,
			solr: `(_query_:"{!edismax}(\"frida kahlo\")" AND _query_:"{!edismax}(exhibitions)")`,
		},
		{
			query: `keyword: {(calico OR "tortoise shell") AND cats}`,
			solr: `((_query_:"{!edismax}(calico)" OR _query_:"{!edismax}(\"tortoise shell\")") AND _query_:"{!edismax}(cats)")`,
		},
		{
			query: ` keyword : { (calico OR tortoise shell) AND cats } `,
			solr: `((_query_:"{!edismax}(calico)" OR _query_:"{!edismax}(tortoise shell)") AND _query_:"{!edismax}(cats)")`,
		},
		{
			query: `(keyword: {calico OR "tortoise shell"})  AND keyword: {cats}`,
			solr: `(((_query_:"{!edismax}(calico)" OR _query_:"{!edismax}(\"tortoise shell\")")) AND _query_:"{!edismax}(cats)")`,
		},
		{
			query: ` keyword : { ("tortoise shell" OR calico) AND cats } `,
			solr: `((_query_:"{!edismax}(\"tortoise shell\")" OR _query_:"{!edismax}(calico)") AND _query_:"{!edismax}(cats)")`,
		},
		{
			query: `keyword: {(calico OR "tortoise shell") AND (cats OR dogs)}`,
			solr: `((_query_:"{!edismax}(calico)" OR _query_:"{!edismax}(\"tortoise shell\")") AND (_query_:"{!edismax}(cats)" OR _query_:"{!edismax}(dogs)"))`,
		},
		{
			query: `keyword:{ \" rotunda \" }`,
			solr: `_query_:"{!edismax}(\ \"rotunda \\")"`,
		},
		{
			query: `subject:{((((((((((turtles AND ((((((((((skateboarding))))))))))))))))))))}`,
			solr: `(_query_:"{!edismax qf=$subject_qf pf=$subject_pf}(turtles)" AND _query_:"{!edismax qf=$subject_qf pf=$subject_pf}(skateboarding)")`,
		},
		{
			query: `subject:{"TRANSMUTATION (Chemistry)"}`,
			solr: `_query_:"{!edismax qf=$subject_qf pf=$subject_pf}(\"TRANSMUTATION (Chemistry)\")"`,
		},
		{
			query: `keyword:{"calvin AND hobbes" OR (susie derkins)}`,
			solr: `(_query_:"{!edismax}(\"calvin AND hobbes\")" OR _query_:"{!edismax}(susie derkins)")`,
		},
		{
			query: `keyword:{"calvin AND hobbes" OR "(susie derkins)"}`,
			solr: `(_query_:"{!edismax}(\"calvin AND hobbes\")" OR _query_:"{!edismax}(\"(susie derkins)\")")`,
		},
	}

	for _, test := range tests {
		sp := expectSolrConversionSuccess(t, test.query, test.solr)

		if test.counts == true {
			expectCounts(t, "title", sp.Titles, test.titles)
			expectCounts(t, "author", sp.Authors, test.authors)
			expectCounts(t, "subject", sp.Subjects, test.subjects)
			expectCounts(t, "keyword", sp.Keywords, test.keywords)
		}
	}
}

func TestSolrShouldFail(t *testing.T) {
	tests := []string {
		``,
		`(`,
		`)`,
		`()`,
		`author:`,
		`subject:{`,
		`identifier:}`,
		`keyword:{tacos`,
		`(keyword:{derivatives}`,
		`title: {bannanas} OR author: bad`,
		`title:{bananas ( a fruit }`,
		`title:{bananas a fruit ) }`,
		`garbage:{1954}`,
		`rubbish:{bananas}`,
		`date:{1932 TO 1945} HELLOOOOO author:{Shelly}`,
		`date:{BadDate}`,
		`subject:{"((((((((((turtles AND ((((((((((skateboarding))))))))))))))))))))}"`,
		`subject:{"((((((((((turtles AND ((((((((((skateboarding))))))))))))))))))))"}`,
	}

	for _, query := range tests {
		expectSolrConversionFailure(t, query)
	}
}
