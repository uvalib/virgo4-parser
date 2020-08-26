// Code generated from VirgoQuery.g4 by ANTLR 4.8. DO NOT EDIT.

package v4parser // VirgoQuery
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseVirgoQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVirgoQueryVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitQuery_parts(ctx *Query_partsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitField_query(ctx *Field_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitField_type(ctx *Field_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitRange_field_type(ctx *Range_field_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitBoolean_op(ctx *Boolean_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitRange_search_string(ctx *Range_search_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitDate_string(ctx *Date_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitSearch_string(ctx *Search_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitSearch_part(ctx *Search_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitQuoted_search_part(ctx *Quoted_search_partContext) interface{} {
	return v.VisitChildren(ctx)
}
