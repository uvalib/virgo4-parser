package v4parser_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/uvalib/virgo4-parser/v4parser"
)

// currently takes ~7 seconds
const slow1 = `keyword: { I have often thought that nothing would do more extensive good at small expense than the establishment of a small circulating library in every county, to consist of a few well-chosen books, to be lent to the people of the country under regulations as would secure their safe return in due time. }`

// currently takes ~33 seconds
const slow2 = `keyword: { If a nation expects to be ignorant and free, in a state of civilization, it expects what never was and never will be. The functionaries of every government have propensities to command at will the liberty and property of their constituents. There is no safe deposit for these but with the people themselves; nor can they be safe with them without information. Where the press is free, and every man able to read, all is safe. }`

type v4SolrQuery struct {
	query   string
	solr    string
	timeout int
}

func (v *v4SolrQuery) convert(t *testing.T) (v4parser.SolrParser, string, error) {
	sp := v4parser.SolrParser{Debug: true, DebugCallTree: false}

	t.Logf("QUERY: %s", v.query)

	var solr string
	var err error

	if v.timeout > 0 {
		solr, err = v4parser.ConvertToSolrWithParserAndTimeout(&sp, v.query, v.timeout)
	} else {
		solr, err = v4parser.ConvertToSolrWithParser(&sp, v.query)
	}

	return sp, solr, err
}

func (v *v4SolrQuery) expectSolrConversionSuccess(t *testing.T) v4parser.SolrParser {
	sp, solr, err := v.convert(t)

	if err != nil {
		t.Errorf("Conversion failed, but should have succeeded:")
		t.Errorf("ERROR: %s", err.Error())
		t.Errorf("WANT : %s", v.solr)

		return sp
	}

	if solr != v.solr {
		t.Errorf("Conversion succeeded, but result was not as expected:")
		t.Errorf("GOT  : %s", solr)
		t.Errorf("WANT : %s", v.solr)
	}

	counts := []string{}

	for field, values := range sp.FieldValues {
		counts = append(counts, fmt.Sprintf("%s: %d", field, len(values)))
	}

	t.Logf("FieldValue counts: %s", strings.Join(counts, "; "))

	return sp
}

func (v *v4SolrQuery) expectSolrConversionFailure(t *testing.T) {
	_, solr, err := v.convert(t)

	if err == nil {
		t.Errorf("Conversion succeeded, but should have failed:")
		t.Errorf("GOT  : %s", solr)
	}
}

func expectCounts(t *testing.T, field string, fieldValues map[string][]string, want int) {
	values := fieldValues[field]

	if values == nil {
		if want != 0 {
			t.Errorf("%s count fail. Expected %d, Actual: none", field, want)
		}
		return
	}

	got := len(values)

	if got != want {
		t.Errorf("%s count fail. Expected %d, Actual: %d, List: %v", field, want, got, values)
	}
}

func TestSolrShouldSucceed(t *testing.T) {
	type expectedResults struct {
		query       string // v4 query
		solr        string // expected solr conversion
		counts      bool   // check counts?
		titles      int    // expected title count
		authors     int    // expected author count
		subjects    int    // expected subject count
		keywords    int    // expected keyword count
		identifiers int    // expected identifier count
		dates       int    // expected date count
		publisheds  int    // expected published count
		filters     int    // expected filter count
	}

	tests := []expectedResults{
		{
			query: `keyword:{}`,
			solr:  `_query_:"{!edismax}(*)"`,
		},
		{
			query: `title: {bannanas}`,
			solr:  `_query_:"{!edismax qf=$title_qf pf=$title_pf}(bannanas)"`,
		},
		{
			query: `title:{A = B}`,
			solr:  `_query_:"{!edismax qf=$title_qf pf=$title_pf}(A = B)"`,
		},
		{
			query: `keyword:{digby OR duncan}`,
			solr:  `(_query_:"{!edismax}(digby)" OR _query_:"{!edismax}(duncan)")`,
		},
		{
			query: `keyword:{digby AND duncan}`,
			solr:  `(_query_:"{!edismax}(digby)" AND _query_:"{!edismax}(duncan)")`,
		},
		{
			query: `author:{Zhongguo Zang xue chu ban she  (publisher)}`,
			solr:  `_query_:"{!edismax qf=$author_qf pf=$author_pf}(Zhongguo Zang xue chu ban she publisher)"`,
		},
		{
			query: `author:{ ʼJam-dbyangs-nyi-ma }`,
			solr:  `_query_:"{!edismax qf=$author_qf pf=$author_pf}(ʼJam\\-dbyangs\\-nyi\\-ma)"`,
		},
		{
			query: `title:{Tragicheskai͡a istorii͡a kavkazskikh }`,
			solr:  `_query_:"{!edismax qf=$title_qf pf=$title_pf}(Tragicheskai͡a istorii͡a kavkazskikh)"`,
		},
		{
			query: `author:{Немирович-Данченко, Василий Иванович}`,
			solr:  `_query_:"{!edismax qf=$author_qf pf=$author_pf}(Немирович\\-Данченко, Василий Иванович)"`,
		},
		{
			query: `title : {"susan sontag" OR (music title)}   AND keyword:{ Maunsell }`,
			solr:  `((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)")`,
		},
		{
			query: `( title : {"susan sontag" OR music title}   AND keyword:{ Maunsell } ) OR author:{ liberty }`,
			solr:  `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)")) OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")`,
		},
		{
			query: `( title : {"susan sontag" OR (music title)}   AND keyword:{ Maunsell } ) OR author:{ liberty }`,
			solr:  `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND _query_:"{!edismax}(Maunsell)")) OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")`,
		},
		{
			query: `title : {"susan sontag" OR (music title)}   AND ( keyword:{ Maunsell } OR author:{ liberty })`,
			solr:  `((_query_:"{!edismax qf=$title_qf pf=$title_pf}(\"susan sontag\")" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(music title)") AND ((_query_:"{!edismax}(Maunsell)" OR _query_:"{!edismax qf=$author_qf pf=$author_pf}(liberty)")))`,
		},
		{
			query: ` author: {lovecraft} AND title: {madness NOT dunwich}`,
			solr:  `(_query_:"{!edismax qf=$author_qf pf=$author_pf}(lovecraft)" AND (_query_:"{!edismax qf=$title_qf pf=$title_pf}(madness)" NOT _query_:"{!edismax qf=$title_qf pf=$title_pf}(dunwich)"))`,
		},
		{
			query: `identifier:{35007007606860  OR 9780754645733 OR 38083649 OR 2001020407  OR u5670758 OR "KJE5602.C73 2012"}`,
			solr:  `(((((_query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(35007007606860)" OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(9780754645733)") OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(38083649)") OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(2001020407)") OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(u5670758)") OR _query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(\"KJE5602.C73 2012\")")`,
		},
		{
			query: `keyword:{triceratops OR flameproof}`,
			solr:  `(_query_:"{!edismax}(triceratops)" OR _query_:"{!edismax}(flameproof)")`,
		},
		{
			query: `identifier:{35007007606860}`,
			solr:  `_query_:"{!edismax qf=$identifier_qf pf=$identifier_pf}(35007007606860)"`,
		},
		{
			query:    `( title: {pepperoni OR artichoke hearts} AND subject:{pizza} ) OR (subject:{calzone} AND (keyword:{italian} NOT author:{fieri}))`,
			solr:     `((((_query_:"{!edismax qf=$title_qf pf=$title_pf}(pepperoni)" OR _query_:"{!edismax qf=$title_qf pf=$title_pf}(artichoke hearts)") AND _query_:"{!edismax qf=$subject_qf pf=$subject_pf}(pizza)")) OR ((_query_:"{!edismax qf=$subject_qf pf=$subject_pf}(calzone)" AND ((_query_:"{!edismax}(italian)" NOT _query_:"{!edismax qf=$author_qf pf=$author_pf}(fieri)")))))`,
			counts:   true,
			titles:   2,
			authors:  1,
			subjects: 2,
			keywords: 1,
		},
		{
			query: `date:{1945}`,
			solr:  `_query_:"{!lucene df=published_daterange}(1945)"`,
		},
		{
			query: `date:{1945/12/07 TO 1949}`,
			solr:  `_query_:"{!lucene df=published_daterange}([1945-12-07 TO 1949])"`,
		},
		{
			query: `date:{BEFORE 1945-12-06}`,
			solr:  `_query_:"{!lucene df=published_daterange}([* TO 1945-12-05])"`,
		},
		{
			query: `date:{AFTER 1945}`,
			solr:  `_query_:"{!lucene df=published_daterange}([1946 TO *])"`,
		},
		{
			query: `date:{<1945} AND date:{>1932} AND author:{Shelly}`,
			solr:  `((_query_:"{!lucene df=published_daterange}([* TO 1944])" AND _query_:"{!lucene df=published_daterange}([1933 TO *])") AND _query_:"{!edismax qf=$author_qf pf=$author_pf}(Shelly)")`,
		},
		{
			query: `keyword:{cincinnati, ohio (home of the :reds:)}`,
			solr:  `_query_:"{!edismax}(cincinnati, ohio home of the \\:reds\\:)"`,
		},
		{
			query: `keyword: {"grapes of wrath"}`,
			solr:  `_query_:"{!edismax}(\"grapes of wrath\")"`,
		},
		{
			query: `keyword: {kyoto NOT protocol}`,
			solr:  `(_query_:"{!edismax}(kyoto)" NOT _query_:"{!edismax}(protocol)")`,
		},
		{
			query: `keyword: {"frida kahlo" AND exhibitions}`,
			solr:  `(_query_:"{!edismax}(\"frida kahlo\")" AND _query_:"{!edismax}(exhibitions)")`,
		},
		{
			query: `keyword: {(calico OR "tortoise shell") AND cats}`,
			solr:  `((_query_:"{!edismax}(calico)" OR _query_:"{!edismax}(\"tortoise shell\")") AND _query_:"{!edismax}(cats)")`,
		},
		{
			query: ` keyword : { (calico OR tortoise shell) AND cats } `,
			solr:  `((_query_:"{!edismax}(calico)" OR _query_:"{!edismax}(tortoise shell)") AND _query_:"{!edismax}(cats)")`,
		},
		{
			query: `(keyword: {calico OR "tortoise shell"})  AND keyword: {cats}`,
			solr:  `(((_query_:"{!edismax}(calico)" OR _query_:"{!edismax}(\"tortoise shell\")")) AND _query_:"{!edismax}(cats)")`,
		},
		{
			query: ` keyword : { ("tortoise shell" OR calico) AND cats } `,
			solr:  `((_query_:"{!edismax}(\"tortoise shell\")" OR _query_:"{!edismax}(calico)") AND _query_:"{!edismax}(cats)")`,
		},
		{
			query: `keyword: {(calico OR "tortoise shell") AND (cats OR dogs)}`,
			solr:  `((_query_:"{!edismax}(calico)" OR _query_:"{!edismax}(\"tortoise shell\")") AND (_query_:"{!edismax}(cats)" OR _query_:"{!edismax}(dogs)"))`,
		},
		{
			query: `keyword:{ \" rotunda \" }`,
			solr:  `_query_:"{!edismax}(\\\\ \" rotunda \\\\\")"`,
		},
		{
			query: `subject:{((((((((((turtles AND ((((((((((skateboarding))))))))))))))))))))}`,
			solr:  `(_query_:"{!edismax qf=$subject_qf pf=$subject_pf}(turtles)" AND _query_:"{!edismax qf=$subject_qf pf=$subject_pf}(skateboarding)")`,
		},
		{
			query: `subject:{"((((((((((turtles AND ((((((((((skateboarding))))))))))))))))))))"}`,
			solr:  `_query_:"{!edismax qf=$subject_qf pf=$subject_pf}(\"((((((((((turtles AND ((((((((((skateboarding))))))))))))))))))))\")"`,
		},
		{
			query: `keyword:{"calvin AND hobbes" OR (susie derkins)}`,
			solr:  `(_query_:"{!edismax}(\"calvin AND hobbes\")" OR _query_:"{!edismax}(susie derkins)")`,
		},
		{
			query: `keyword:{"calvin AND hobbes" OR "(susie derkins)"}`,
			solr:  `(_query_:"{!edismax}(\"calvin AND hobbes\")" OR _query_:"{!edismax}(\"(susie derkins)\")")`,
		},
		{
			query: `keyword:{"calvin and hobbes" or "(susie derkins)"}`,
			solr:  `_query_:"{!edismax}(\"calvin and hobbes\" or \"(susie derkins)\")"`,
		},
		{
			query: `subject:{"TRANSMUTATION (Chemistry)"}`,
			solr:  `_query_:"{!edismax qf=$subject_qf pf=$subject_pf}(\"TRANSMUTATION (Chemistry)\")"`,
		},
		{
			query: `keyword: {"transmutation AND (chemistry)"}`,
			solr:  `_query_:"{!edismax}(\"transmutation AND (chemistry)\")"`,
		},
		{
			query: `keyword: {("Death AND taxes" OR "war AND Peace") AND (cats OR dogs)}`,
			solr:  `((_query_:"{!edismax}(\"Death AND taxes\")" OR _query_:"{!edismax}(\"war AND Peace\")") AND (_query_:"{!edismax}(cats)" OR _query_:"{!edismax}(dogs)"))`,
		},
		{
			query: `keyword: {a \ b}`,
			solr:  `_query_:"{!edismax}(a \\\\ b)"`,
		},
		{
			query: `keyword: {a + b}`,
			solr:  `_query_:"{!edismax}(a \\+ b)"`,
		},
		{
			query: `keyword: {a - b}`,
			solr:  `_query_:"{!edismax}(a \\- b)"`,
		},
		{
			query: `keyword: {a && b}`,
			solr:  `_query_:"{!edismax}(a \\&& b)"`,
		},
		{
			query: `keyword: {a || b}`,
			solr:  `_query_:"{!edismax}(a \\|| b)"`,
		},
		{
			query: `keyword: {a ! b}`,
			solr:  `_query_:"{!edismax}(a \\! b)"`,
		},
		{
			query: `keyword: {a (b) c}`,
			solr:  `_query_:"{!edismax}(a b c)"`,
		},
		{
			query: `keyword: {a * b}`,
			solr:  `_query_:"{!edismax}(a * b)"`,
		},
		{
			query: `keyword: {a ^ b}`,
			solr:  `_query_:"{!edismax}(a \\^ b)"`,
		},
		{
			query: `keyword: {a " b " c}`,
			solr:  `_query_:"{!edismax}(a \" b \" c)"`,
		},
		{
			query: `keyword: {a ~ b}`,
			solr:  `_query_:"{!edismax}(a \\~ b)"`,
		},
		{
			query: `keyword: {a ? b}`,
			solr:  `_query_:"{!edismax}(a \\? b)"`,
		},
		{
			query: `keyword: {a : b}`,
			solr:  `_query_:"{!edismax}(a \\: b)"`,
		},
		{
			query: `keyword: {a / b}`,
			solr:  `_query_:"{!edismax}(a \\/ b)"`,
		},
		{
			query: `published: {New York City}`,
			solr:  `_query_:"{!edismax qf=$published_qf pf=$published_pf}(New York City)"`,
		},
		{
			query: `keyword: {"Organic chemistry"}  AND filter:{data_source_f:libraetd  OR  data_source_f:libraoc}`,
			solr:  `(_query_:"{!edismax}(\"Organic chemistry\")" AND (_query_:"(data_source_f:libraetd)" OR _query_:"(data_source_f:libraoc)"))`,
		},
		{
			query: `keyword: {chicago AND ("white sox" NOT “cubs”)}`,
			solr:  `(_query_:"{!edismax}(chicago)" AND (_query_:"{!edismax}(\"white sox\")" NOT _query_:"{!edismax}(\"cubs\")"))`,
		},
		{
			query: `journal_title:{nature} AND title:{orangutan}`,
			solr:  `(_query_:"{!edismax qf=$journal_title_qf pf=$journal_title_pf}(nature)" AND _query_:"{!edismax qf=$title_qf pf=$title_pf}(orangutan)")`,
		},
		{
			query: `fulltext:{Frank C.  MCCUE}`,
			solr:  `_query_:"{!edismax qf=$fulltext_qf pf=$fulltext_pf}(Frank C. MCCUE)"`,
		},
		{
			query: `series:{Early imprints}`,
			solr:  `_query_:"{!edismax qf=$series_qf pf=$series_pf}(Early imprints)"`,
		},
	}

	for _, test := range tests {
		v := v4SolrQuery{query: test.query, solr: test.solr}

		sp := v.expectSolrConversionSuccess(t)

		if test.counts == true {
			expectCounts(t, "title", sp.FieldValues, test.titles)
			expectCounts(t, "author", sp.FieldValues, test.authors)
			expectCounts(t, "subject", sp.FieldValues, test.subjects)
			expectCounts(t, "keyword", sp.FieldValues, test.keywords)
			expectCounts(t, "identifiers", sp.FieldValues, test.identifiers)
			expectCounts(t, "date", sp.FieldValues, test.dates)
		}
	}
}

func TestSolrShouldFail(t *testing.T) {
	tests := []string{
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
		`subject:{"(((turtles AND (((skateboarding))))))}"`,
		`subject:{"(((turtles AND (((skateboarding))))))}`,
		`keyword: {a [ b ] c}`,
	}

	for _, query := range tests {
		v := v4SolrQuery{query: query}

		v.expectSolrConversionFailure(t)
	}
}

// with SLL prediction mode, these remaining tests/benchmarks are now fast.  see:
// * https://github.com/antlr/antlr4/blob/master/doc/faq/general.md  (section "Why is my expression parser slow?")
// * https://github.com/antlr/antlr4/issues/374

func TestSlowConversion1(t *testing.T) {
	v := v4SolrQuery{
		query: slow1,
		solr:  `_query_:"{!edismax}(I have often thought that nothing would do more extensive good at small expense than the establishment of a small circulating library in every county, to consist of a few well\\-chosen books, to be lent to the people of the country under regulations as would secure their safe return in due time.)"`,
	}

	v.expectSolrConversionSuccess(t)
}

func TestSlowConversion2(t *testing.T) {
	v := v4SolrQuery{
		query: slow2,
		solr:  `_query_:"{!edismax}(If a nation expects to be ignorant and free, in a state of civilization, it expects what never was and never will be. The functionaries of every government have propensities to command at will the liberty and property of their constituents. There is no safe deposit for these but with the people themselves; nor can they be safe with them without information. Where the press is free, and every man able to read, all is safe.)"`,
	}

	v.expectSolrConversionSuccess(t)
}

/*
func TestSlowConversionWithTimeout(t *testing.T) {
	v := v4SolrQuery{
		query:   slow2,
		timeout: 5,
	}

	v.expectSolrConversionFailure(t)
}
*/

func BenchmarkSlowConversion1(b *testing.B) {
	v4parser.ConvertToSolr(slow1)
}

func BenchmarkSlowConversion2(b *testing.B) {
	v4parser.ConvertToSolr(slow2)
}
