// Generated from VirgoQuery.g4 by ANTLR 4.7.

package v4parser // VirgoQuery
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by VirgoQuery.
type VirgoQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VirgoQuery#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by VirgoQuery#query_parts.
	VisitQuery_parts(ctx *Query_partsContext) interface{}

	// Visit a parse tree produced by VirgoQuery#field_query.
	VisitField_query(ctx *Field_queryContext) interface{}

	// Visit a parse tree produced by VirgoQuery#field_type.
	VisitField_type(ctx *Field_typeContext) interface{}

	// Visit a parse tree produced by VirgoQuery#boolean_op.
	VisitBoolean_op(ctx *Boolean_opContext) interface{}

	// Visit a parse tree produced by VirgoQuery#search_string.
	VisitSearch_string(ctx *Search_stringContext) interface{}

	// Visit a parse tree produced by VirgoQuery#search_part.
	VisitSearch_part(ctx *Search_partContext) interface{}
}
