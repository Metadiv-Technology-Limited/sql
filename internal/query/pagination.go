package query

import (
	"github.com/Metadiv-Technology-Limited/sql/models"

	"gorm.io/gorm"
)

/*
ConsumePagination consumes pagination information from models.IPagination and
returns a *gorm.DB.
*/
func ConsumePagination(tx *gorm.DB, p models.IPagination) *gorm.DB {
	if p == nil {
		return tx
	}
	if p.GetPage() > 0 {
		tx = tx.Offset((p.GetPage() - 1) * p.GetSize())
	}
	if p.GetSize() > 0 {
		tx = tx.Limit(p.GetSize())
	}
	return tx
}
