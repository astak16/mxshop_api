package forms

type GoodsForm struct {
	Name        string   `form:"name" json:"name" binding:"required,min=2,max=100"`
	GoodsSn     string   `form:"goods_sn" json:"goods_sn" binding:"required,min=2,lt=20"`
	Stocks      int32    `form:"stocks" json:"stocks" binding:"required,min=1"`
	CategoryId  int32    `form:"category" json:"category" binding:"required"`
	MarketPrice float32  `form:"market_price" json:"market_price" binding:"required,min=0"`
	ShopPrice   float32  `form:"shop_price" json:"shop_price" binding:"required,min=0"`
	GoodsBrief  string   `form:"goods_brief" json:"goods_brief" binding:"required,min=3"`
	Images      []string `form:"images" json:"images" binding:"required,min=1"`
	DescImages  []string `form:"desc_images" json:"desc_images" binding:"required,min=1"`
	ShipFree    *bool    `form:"ship_free" json:"ship_free" binding:"required"`
	FrontImage  string   `form:"front_image" json:"front_image" binding:"required,url"`
	Brand       int32    `form:"brand" json:"brand" binding:"required"`
}

type GoodsStatusForm struct {
	IsNew  *bool `form:"new" json:"new" binding:"required"`
	IsHot  *bool `form:"hot" json:"hot" binding:"required"`
	OnSale *bool `form:"sale" json:"sale" binding:"required"`
}

type CategoryForm struct {
	Name           string `form:"name" json:"name" binding:"required,min=3,max=20"`
	ParentCategory int32  `form:"parent" json:"parent"`
	Level          int32  `form:"level" json:"level" binding:"required,oneof=1 2 3"`
	IsTab          *bool  `form:"is_tab" json:"is_tab" binding:"required"`
}

type UpdateCategoryForm struct {
	Name  string `form:"name" json:"name" binding:"required,min=3,max=20"`
	IsTab *bool  `form:"is_tab" json:"is_tab"`
}

type BannerForm struct {
	Image string `form:"image" json:"image" binding:"url"`
	Index int    `form:"index" json:"index" binding:"required"`
	Url   string `form:"url" json:"url" binding:"url"`
}

type BrandForm struct {
	Name string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Logo string `form:"logo" json:"logo" binding:"url"`
}

type CategoryBrandForm struct {
	CategoryId int `form:"category_id" json:"category_id" binding:"required"`
	BrandId    int `form:"brand_id" json:"brand_id" binding:"required"`
}
