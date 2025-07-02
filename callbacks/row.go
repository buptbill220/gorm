package callbacks

import (
	"gorm.io/gorm"
	"gorm.io/gorm/apaas"
)

func RowQuery(db *gorm.DB) {
	if db.Error == nil {
		BuildQuerySQL(db)
		if db.DryRun || db.Error != nil {
			return
		}
		needApaasServer := !db.Statement.ApaasServerMode && !db.Statement.ApaasOff && db.Statement.ApaasMode != apaas.DirectMode
		if isRows, ok := db.Get("rows"); ok && isRows.(bool) {
			db.Statement.Settings.Delete("rows")
			if needApaasServer {
				db.Statement.Dest, db.Error = apaas.NewApaasEngineClient().QueryContext(db.Statement.Context, db.Statement.ApaasDSLArgs)
			} else {
				db.Statement.Dest, db.Error = db.Statement.ConnPool.QueryContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
			}
		} else {
			if needApaasServer {
				db.Statement.Dest, db.Error = apaas.NewApaasEngineClient().QueryRowContext(db.Statement.Context, db.Statement.ApaasDSLArgs)
			} else {
				db.Statement.Dest = db.Statement.ConnPool.QueryRowContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
			}
		}
		db.RowsAffected = -1
	}
}
