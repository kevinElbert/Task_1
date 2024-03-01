package helper

import (
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollBack := tx.Rollback().Error
		PanicIfError(errorRollBack)
	} else {
		errorCommit := tx.Commit().Error
		PanicIfError(errorCommit)
	}
}
