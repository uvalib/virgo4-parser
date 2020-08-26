// Code generated from VirgoQuery.g4 by ANTLR 4.8. DO NOT EDIT.

package v4parser // VirgoQuery
import "github.com/antlr/antlr4/runtime/Go/antlr"

// VirgoQueryListener is a complete listener for a parse tree produced by VirgoQuery.
type VirgoQueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterQuery_parts is called when entering the query_parts production.
	EnterQuery_parts(c *Query_partsContext)

	// EnterField_query is called when entering the field_query production.
	EnterField_query(c *Field_queryContext)

	// EnterField_type is called when entering the field_type production.
	EnterField_type(c *Field_typeContext)

	// EnterRange_field_type is called when entering the range_field_type production.
	EnterRange_field_type(c *Range_field_typeContext)

	// EnterBoolean_op is called when entering the boolean_op production.
	EnterBoolean_op(c *Boolean_opContext)

	// EnterRange_search_string is called when entering the range_search_string production.
	EnterRange_search_string(c *Range_search_stringContext)

	// EnterDate_string is called when entering the date_string production.
	EnterDate_string(c *Date_stringContext)

	// EnterSearch_string is called when entering the search_string production.
	EnterSearch_string(c *Search_stringContext)

	// EnterSearch_part is called when entering the search_part production.
	EnterSearch_part(c *Search_partContext)

	// EnterQuoted_search_part is called when entering the quoted_search_part production.
	EnterQuoted_search_part(c *Quoted_search_partContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitQuery_parts is called when exiting the query_parts production.
	ExitQuery_parts(c *Query_partsContext)

	// ExitField_query is called when exiting the field_query production.
	ExitField_query(c *Field_queryContext)

	// ExitField_type is called when exiting the field_type production.
	ExitField_type(c *Field_typeContext)

	// ExitRange_field_type is called when exiting the range_field_type production.
	ExitRange_field_type(c *Range_field_typeContext)

	// ExitBoolean_op is called when exiting the boolean_op production.
	ExitBoolean_op(c *Boolean_opContext)

	// ExitRange_search_string is called when exiting the range_search_string production.
	ExitRange_search_string(c *Range_search_stringContext)

	// ExitDate_string is called when exiting the date_string production.
	ExitDate_string(c *Date_stringContext)

	// ExitSearch_string is called when exiting the search_string production.
	ExitSearch_string(c *Search_stringContext)

	// ExitSearch_part is called when exiting the search_part production.
	ExitSearch_part(c *Search_partContext)

	// ExitQuoted_search_part is called when exiting the quoted_search_part production.
	ExitQuoted_search_part(c *Quoted_search_partContext)
}
