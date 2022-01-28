module server

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
)

require (
	common v0.0.0
	github.com/apache/rocketmq-client-go/v2 v2.1.0
	github.com/spf13/viper v1.10.1
)

replace common => ../common
