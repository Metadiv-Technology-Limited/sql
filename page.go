package sql

import "sql/models"

/*
Paginate creates a new models.IPagination.
*/
func Paginate(page, size int) models.IPagination {
	return &models.Pagination{
		Page: page,
		Size: size,
	}
}
