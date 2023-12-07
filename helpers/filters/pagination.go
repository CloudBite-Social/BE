package filters

type Pagination struct {
	Limit int `query:"limit"`
	Start int `query:"start"`
}
