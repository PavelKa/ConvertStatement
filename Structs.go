package main

type ConversionFieldFunction interface{}

type ConversionImpl struct {
	index     int
	fieldName string
	function  ConversionFieldFunction
}

func newConversionImpl(fieldName string, csvIndex int) ConversionImpl {
	mvf := ConversionImpl{}
	mvf.fieldName = fieldName
	mvf.function = getRecItem
	mvf.index = csvIndex
	return mvf

}
func newConversionImplF(fieldName string, csvIndex int, f ConversionFieldFunction) ConversionImpl {
	mvf := ConversionImpl{}
	mvf.fieldName = fieldName
	mvf.function = f
	mvf.index = csvIndex
	return mvf

}

type ConversionFields []ConversionImpl

type ItemList interface {
	AddItem(Item)
	CreateItem() Item
}

type Item interface {
	SetByName(fieldName string, value string) Item
}
