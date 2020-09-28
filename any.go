package lazy

//go:generate go run github.com/cheekybits/genny -in=main.go -out=zmain.go gen "Item=string,Bytes,bool,int,int64,uint,uint64,uintptr,float64,Any"

type Any = interface{}
type Bytes = []byte
