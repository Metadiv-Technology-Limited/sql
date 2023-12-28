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

/*
Sort creates a new models.ISorting.
*/
func Sort(by string, asc bool) models.ISorting {
	return &models.Sorting{
		By:  by,
		Asc: asc,
	}
}
