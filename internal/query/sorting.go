package query

import (
	"github.com/Metadiv-Technology-Limited/sql/models"

	"gorm.io/gorm"
)

/*
ConsumeSorting consumes sorting information from models.ISorting and returns a
*gorm.DB.
*/
func ConsumeSorting(tx *gorm.DB, s models.ISorting) *gorm.DB {
	if s == nil {
		return tx
	}
	if s.GetBy() != "" {
		if s.GetAsc() {
			tx = tx.Order(s.GetBy())
		} else {
			tx = tx.Order(s.GetBy() + " DESC")
		}
	}
	return tx
}
