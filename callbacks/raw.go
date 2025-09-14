package callbacks

import (
	"gorm.io/gorm"
)

func RawExec(db *gorm.DB) {
	if db.Error == nil && !db.DryRun {
		sql := db.Statement.SQL.String()
		db.Statement.ApaasDSLArgs.RawSQL = sql
		db.Statement.ApaasDSLArgs.SQL = db.Dialector.Explain(sql, db.Statement.Vars...)
		db.Statement.ApaasDSLArgs.Vars = db.Statement.Vars
		result, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, sql, db.Statement.Vars...)
		if err != nil {
			db.AddError(err)
			return
		}

		db.RowsAffected, _ = result.RowsAffected()
	}
}
