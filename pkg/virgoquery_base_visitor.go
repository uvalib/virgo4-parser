// Generated from VirgoQuery.g4 by ANTLR 4.7.

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

func (v *BaseVirgoQueryVisitor) VisitBoolean_op(ctx *Boolean_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitSearch_string(ctx *Search_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVirgoQueryVisitor) VisitSearch_part(ctx *Search_partContext) interface{} {
	return v.VisitChildren(ctx)
}
