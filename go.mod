module github.com/sageflow/sageauth

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.3
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/sageflow/sageapi v0.0.0-20210310195905-496b28dffb15 // indirect
	github.com/sageflow/sageflow v0.0.0-20210310195444-994b73c0ab77
	github.com/soheilhy/cmux v0.1.4
	github.com/vektah/gqlparser v1.3.1 // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/examples v0.0.0-20210305213134-61f0b5fa7c1c // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/sageflow/sageflow v0.0.0-20210310195444-994b73c0ab77 => ../sageflow
