package data

import "greenlight/alan.com/internal/validator"

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func ValidateFilters(v *validator.Validator, f Filters) {

	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10000000, "page_size", "must be less than 10,000,000")
	v.Check(f.PageSize > 0 && f.PageSize <= 100, "page_size", "must be between 1 and 100")
	v.Check(validator.In(f.Sort, f.SortSafelist...), "sort", "invalid sort value")

}
