package domain

type ExpenseFilters struct {
	Category       string
	TimestampStart string
	TimestampEnd   string
	Page           int
	PageSize       int
}
