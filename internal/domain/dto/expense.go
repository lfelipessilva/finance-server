package domain

type ExpenseFilters struct {
	UserID         uint
	Category       string
	TimestampStart string
	TimestampEnd   string
	Name           string
	TagIds         []uint
	Page           int
	PageSize       int
	OrderBy        string
	OrderDirection string
}
