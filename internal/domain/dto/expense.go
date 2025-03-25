package domain

type ExpenseFilters struct {
	Category       string
	TimestampStart string
	TimestampEnd   string
	Name           string
	Page           int
	PageSize       int
	OrderBy        string
	OrderDirection string
}
