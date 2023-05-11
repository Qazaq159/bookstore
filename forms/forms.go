package forms

type GetBooksParam struct {
	IsDesc    bool   `form:"desc"`
	Title     string `form:"query"`
	Published bool   `form:"published"`
}
