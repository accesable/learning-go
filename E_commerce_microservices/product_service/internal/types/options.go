package types

type GetItemsOption func(*getItemsOptions)

type getItemsOptions struct {
	ShowCategoryName bool
	IncludeImgURLs   bool
}

func WithCategoryName() GetItemsOption {
	return func(opts *getItemsOptions) {
		opts.ShowCategoryName = true
	}
}

func WithImgURLs() GetItemsOption {
	return func(opts *getItemsOptions) {
		opts.IncludeImgURLs = true
	}
}

func NewGetItemsOptions() *getItemsOptions {
	return &getItemsOptions{}
}
