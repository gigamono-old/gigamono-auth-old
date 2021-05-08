module github.com/gigamono/gigamono-auth

go 1.15

require (
	github.com/gigamono/gigamono v0.0.0-20210503171043-f173ed5d20cd
	github.com/gin-gonic/gin v1.7.1
	github.com/soheilhy/cmux v0.1.5
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/grpc v1.37.0
)

replace github.com/gigamono/gigamono v0.0.0-20210503171043-f173ed5d20cd => ../gigamono
