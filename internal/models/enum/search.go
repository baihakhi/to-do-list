package enum

type SearchBy string

const (
	NAME SearchBy = "name"
	CODE SearchBy = "code"
)

var IsValidSearchBy = map[SearchBy]bool{
	NAME: true,
	CODE: true,
}
