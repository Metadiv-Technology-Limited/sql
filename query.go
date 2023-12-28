package sql

import (
	"sql/internal/clause"
	"sql/internal/query"
	"sql/models"

	"gorm.io/gorm"
)

/*
Paginate creates a new models.IPagination.
*/
func Paginate(page, size int) models.IPagination {
	return &models.Pagination{
		Page: page,
		Size: size,
	}
}

/*
Sort creates a new models.ISorting.
*/
func Sort(by string, asc bool) models.ISorting {
	return &models.Sorting{
		By:  by,
		Asc: asc,
	}
}

/*
FindOne finds one record from the database.
*/
func FindOne[T any](tx *gorm.DB, cls models.IClause) (*T, error) {
	if cls != nil {
		tx = clause.Consume(tx, cls)
	}
	output := new(T)
	err := tx.First(output).Error
	if err != nil {
		return nil, err
	}
	return output, err
}

/*
FindAll finds all records from the database.
*/
func FindAll[T any](tx *gorm.DB, cls models.IClause) ([]T, error) {
	if cls != nil {
		tx = clause.Consume(tx, cls)
	}
	output := make([]T, 0)
	err := tx.Find(&output).Error
	if err != nil {
		return nil, err
	}
	return output, err
}

/*
FindAllComplex finds all records from the database with pagination and sorting.
*/
func FindAllComplex[T any](tx *gorm.DB, cls models.IClause, p models.IPagination, s models.ISorting) ([]T, models.IPagination, error) {

	// part 1: find all records
	tx1 := tx.Begin()
	if cls != nil {
		tx1 = clause.Consume(tx1, cls)
	}
	if p != nil && p.GetPage() > 0 && p.GetSize() > 0 {
		tx1 = query.ConsumePagination(tx1, p)
	}
	if s != nil && s.GetBy() != "" {
		tx1 = query.ConsumeSorting(tx1, s)
	}
	output := make([]T, 0)
	err := tx1.Find(&output).Error
	if err != nil {
		tx1.Rollback()
		return nil, nil, err
	}
	tx1.Commit()

	// part 2: count total records
	if p != nil && p.GetPage() > 0 && p.GetSize() > 0 {
		tx2 := tx.Begin()
		if cls != nil {
			tx2 = clause.Consume(tx2, cls)
		}
		var total int64
		err = tx2.Model(new(T)).Count(&total).Error
		if err != nil {
			tx2.Rollback()
			return nil, nil, err
		}
		tx2.Commit()
		p.SetTotal(total)
	} else {
		p = nil
	}

	return output, p, err
}

/*
Count counts the number of records from the database.
*/
func Count[T any](tx *gorm.DB, cls models.IClause) (int64, error) {
	if cls != nil {
		tx = clause.Consume(tx, cls)
	}
	var total int64
	err := tx.Model(new(T)).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, err
}
