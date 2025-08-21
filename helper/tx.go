package helper

import (
	"database/sql"
)

func CommitOrRollBack(tx *sql.Tx) {
	err := recover()
		if err != nil {
			errorRollBack := tx.Rollback()
			PanicIfErr(errorRollBack)
		} else {
			errorCommit := tx.Commit()
			PanicIfErr(errorCommit)
		}
}