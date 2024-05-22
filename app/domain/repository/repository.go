package repository

type ListArgs[T1 any] struct {
	Limit  int `db:"limit"`
	Offset int `db:"offset"`
	Sort   SortStr
	Filter *T1
}

type SortStr string
