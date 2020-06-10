package stbz

// var (
// 	Method.
// )
type method uint

// UtilMethod UtilMethod
type UtilMethod struct {
	GET  method
	POST method
}

// Method Method
var Method = new(UtilMethod)

func init() {
	Method.GET = 1
	Method.POST = 2
}
