package main

import (
	"flag"
	"strings"

	"github.com/pingcap/tidb/planner/core"
	"github.com/s3nt3/sqlvine/internal/logger"
	"github.com/s3nt3/sqlvine/internal/session"
)

var (
	opt_sql = flag.String("sql", "SELECT c1, c2 FROM t1;", "SQL to build")
)

func main() {
	flag.Parse()

	parser := session.NewTiDBParser()
	stmts, warns, err := parser.Parse([]byte(*opt_sql))
	if err != nil {
		logger.L.Panic(err.Error())
	}

	if len(warns) > 0 {
		for _, warn := range warns {
			logger.L.Debug(warn.Error())
		}
	}

	schema := session.NewSchema(`[{
		"id": 1,
		"name": "t1",
		"charset": "utf8mb4",
		"collate": "utf8mb4_bin",

		"columns": [{
			"id": 1,
			"name": "c1",
			"table": "t1",
			"type": "int",
			"primary_key": true
		},{
			"id": 2,
			"name": "c2",
			"table": "t1",
			"type": "varchar"
		},{
			"id": 3,
			"name": "c3",
			"table": "t1",
			"type": "varchar"
		}],
		"indices": []
	}]`)
	builder := session.NewTiDBPlanBuilder(schema.GetSchemaInfo())
	for _, stmt := range stmts {
		plan, _, err := builder.Build(stmt)
		if err != nil {
			logger.L.Debug(err.Error())
		} else {
			PrintLogicalPlan(plan.(core.LogicalPlan), 0)
		}
	}
}

func PrintLogicalPlan(plan core.LogicalPlan, level int) {
	if plan != nil {
		logger.L.Debugf("|%s%T", strings.Repeat("-", level), plan)
		for _, p := range plan.Children() {
			PrintLogicalPlan(p, level+1)
		}
	}
}