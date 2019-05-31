// Generated from VirgoQuery.g4 by ANTLR 4.7.

package v4parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVirgoQueryListener is a complete listener for a parse tree produced by VirgoQuery.
type BaseVirgoQueryListener struct{}

var _ VirgoQueryListener = &BaseVirgoQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVirgoQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVirgoQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVirgoQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVirgoQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseVirgoQueryListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseVirgoQueryListener) ExitQuery(ctx *QueryContext) {}

// EnterQuery_parts is called when production query_parts is entered.
func (s *BaseVirgoQueryListener) EnterQuery_parts(ctx *Query_partsContext) {}

// ExitQuery_parts is called when production query_parts is exited.
func (s *BaseVirgoQueryListener) ExitQuery_parts(ctx *Query_partsContext) {}

// EnterField_query is called when production field_query is entered.
func (s *BaseVirgoQueryListener) EnterField_query(ctx *Field_queryContext) {}

// ExitField_query is called when production field_query is exited.
func (s *BaseVirgoQueryListener) ExitField_query(ctx *Field_queryContext) {}

// EnterField_type is called when production field_type is entered.
func (s *BaseVirgoQueryListener) EnterField_type(ctx *Field_typeContext) {}

// ExitField_type is called when production field_type is exited.
func (s *BaseVirgoQueryListener) ExitField_type(ctx *Field_typeContext) {}

// EnterBoolean_op is called when production boolean_op is entered.
func (s *BaseVirgoQueryListener) EnterBoolean_op(ctx *Boolean_opContext) {}

// ExitBoolean_op is called when production boolean_op is exited.
func (s *BaseVirgoQueryListener) ExitBoolean_op(ctx *Boolean_opContext) {}

// EnterSearch_string is called when production search_string is entered.
func (s *BaseVirgoQueryListener) EnterSearch_string(ctx *Search_stringContext) {}

// ExitSearch_string is called when production search_string is exited.
func (s *BaseVirgoQueryListener) ExitSearch_string(ctx *Search_stringContext) {}

// EnterSearch_part is called when production search_part is entered.
func (s *BaseVirgoQueryListener) EnterSearch_part(ctx *Search_partContext) {}

// ExitSearch_part is called when production search_part is exited.
func (s *BaseVirgoQueryListener) ExitSearch_part(ctx *Search_partContext) {}
