package main

// Product 表示要构建的复杂对象
type Product struct {
	Brand  string
	Model  string
	Color  string
	Option string
}

// Builder 是抽象的建造者接口
type Builder interface {
	SetBrand(brand string)
	SetModel(model string)
	SetColor(color string)
	SetOption(option string)
	GetProduct() Product
}

// ConcreteBuilder1 是具体的建造者实现
type ConcreteBuilder1 struct {
	product Product
}

// SetBrand 设置品牌
func (b *ConcreteBuilder1) SetBrand(brand string) {
	b.product.Brand = brand
}

// SetModel 设置型号
func (b *ConcreteBuilder1) SetModel(model string) {
	b.product.Model = model
}

// SetColor 设置颜色
func (b *ConcreteBuilder1) SetColor(color string) {
	b.product.Color = color
}

// SetOption 设置可选项
func (b *ConcreteBuilder1) SetOption(option string) {
	b.product.Option = option
}

// GetProduct 获取构建的产品
func (b *ConcreteBuilder1) GetProduct() Product {
	return b.product
}

// Director 是指导者，负责控制建造过程
type Director struct {
	builder Builder
}

// Construct 构建产品的过程
func (d *Director) Construct() Product {
	d.builder.SetBrand("BrandA")
	d.builder.SetModel("ModelA")
	d.builder.SetColor("Red")
	d.builder.SetOption("OptionA")
	return d.builder.GetProduct()
}
