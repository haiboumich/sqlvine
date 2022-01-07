package mutator

import (
	"github.com/pingcap/tidb/parser/ast"
	"github.com/s3nt3/sqlvine/internal/ir"
	"github.com/s3nt3/sqlvine/pkg/generator"
)

func (m *Mutator) mutateSelectStmt(node *ir.MutNode) {
	switch node.Node.(type) {
	case *ast.SelectStmt:
		m.MutateSelectStmtNode(node)
	}
}

func (m *Mutator) MutateSelectStmtNode(node *ir.MutNode) {
	stmt := node.Node.(*ast.SelectStmt)

	switch {
	case stmt.From == nil:
		stmt.From = generator.NewASTGenerator().TableRefsClause(m.RandomNum(1))
	case stmt.Where == nil:
		stmt.Where = generator.NewASTGenerator().BinaryOperationExpr(m.RandomNum(1))
	}

	switch m.RandomNum(6) {
	case 0:
		stmt.From = generator.NewASTGenerator().TableRefsClause(m.RandomNum(3))
	case 1:
		stmt.Where = generator.NewASTGenerator().BinaryOperationExpr(m.RandomNum(3))
	case 2:
		stmt.GroupBy = generator.NewASTGenerator().GroupByClause()
	case 3:
		if stmt.GroupBy == nil {
			stmt.GroupBy = generator.NewASTGenerator().GroupByClause()
		}
		stmt.Having = generator.NewASTGenerator().HavingClause(m.RandomNum(3))
	case 4:
		stmt.OrderBy = generator.NewASTGenerator().OrderByClause()
	case 5:
		stmt.Limit = generator.NewASTGenerator().Limit(0, 1)
	default:
		switch {
		case stmt.Limit != nil:
			stmt.Limit = nil
		case stmt.OrderBy != nil:
			stmt.OrderBy = nil
		case stmt.Having != nil:
			stmt.Having = nil
		case stmt.GroupBy != nil:
			stmt.GroupBy = nil
		}
	}
}