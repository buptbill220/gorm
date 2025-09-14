package callbacks

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/apaas"
)

var dbNameCaller func(*gorm.DB) (string, error)

func SetDBNameCaller(fn func(*gorm.DB) (string, error)) {
	dbNameCaller = fn
}

func ApaasExtraCheckerCallBack(stage string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		if db.Statement.ApaasOff {
			return
		}
		ApaasWriteModeCheckerCallBack(stage)(db)
		if db.Error == nil && db.Statement.Schema != nil {
			db.Statement.ApaasMode = apaas.DirectMode
			dbName := db.Config.DBName
			if dbName == "" {
				//db.Error = db.AddError(GenError(fmt.Sprintf("%s ExtraCheckerCallBack(stage=%s) GetDBName nil", MSG_PREFIX, stage)))
				return
			}
			dbCol := apaas.GetDBCol()
			if dbCol == nil {
				//db.Error = db.AddError(GenError(fmt.Sprintf("%s ExtraCheckerCallBack(stage=%s) GetDBCollection nil ", MSG_PREFIX, stage)))
				return
			}
			dbMeta, ok := dbCol.GetDB(dbName)
			if !ok {
				//db.Error = db.AddError(GenError(fmt.Sprintf("%s ExtraCheckerCallBack(stage=%s) GetDB(db=%s) nil", dbName, MSG_PREFIX, stage)))
				return
			}
			tableMeta, ok := dbMeta.GetTableByName(db.Statement.Table)
			if !ok {
				//db.Error = db.AddError(GenError(fmt.Sprintf("%s ExtraCheckerCallBack(stage=%s) GetTable(db=%s,table=%s) GetDB nil", MSG_PREFIX, stage, dbName, db.Statement.Table)))
				return
			}
			db.Logger.Info(db.Statement.Context, "%s ExtraCheckerCallBack(stage=%s) db_name=%s, table=%s", apaas.MSG_PREFIX, stage, dbName, tableMeta.TableName)
			if tableMeta.ExtraMeta != nil {
				db.Logger.Info(db.Statement.Context, "%s ExtraCheckerCallBack(stage=%s) db_name=%s, table=%s, check extra[room_extra len=%d, contract_extra len=%d] begin", apaas.MSG_PREFIX, stage, dbName, db.Statement.Table, len(tableMeta.ExtraMeta.ExtraFields), len(tableMeta.ExtraMeta.ContrctExtra))
				dest := db.Statement.DestTOMap()
				extra := db.Statement.ApaasExtra
				db.Logger.Info(db.Statement.Context, "dest: %v, \nextra: %v", dest, extra)
				if err := tableMeta.ExtraMeta.Check(dest, extra); err != nil {
					db.Error = db.AddError(apaas.GenError(fmt.Sprintf("ExtraCheckerCallBack(stage=%s) db=%s,table=%s check error=(%s)", stage, dbName, db.Statement.Table, err.Error())))
					return
				}
			}
		}
	}
}

func ApaasWriteModeCheckerCallBack(stage string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		if db.Statement.ApaasOff {
			return
		}
		if db.Error == nil && db.Statement.Schema != nil {
			ApaasDBSetCallBack(db)
			if db.Statement.ApaasMode == apaas.ViewMode {
				db.AddError(apaas.GenError(fmt.Sprintf("%s ExtraCheckerCallBack(stage=%s) current model(db=%s, view table=%s) contains LookupView, don't allow write mode", apaas.MSG_PREFIX, stage, db.Config.DBName, db.Statement.Table)))
				return
			}
			db.Statement.ApaasMode = apaas.DirectMode
		}
	}
}

func ApaasDBSetCallBack(db *gorm.DB) {
	if db.Error == nil && db.Statement.Schema != nil {
		if db.Config.DBName == "" {
			if dbNameCaller != nil {
				db.Config.DBName, _ = dbNameCaller(db)
			} else {
				db.Config.DBName, _ = db.GetDBName()
			}
		}
	}
}
