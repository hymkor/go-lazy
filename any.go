package lazy

//go:generate go run github.com/cheekybits/genny -in=main.go -out=zmain.go gen "Item=string,int,bool,Any"

type Any = interface{}
