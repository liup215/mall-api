go build -o app/api/web/member/cmd/cmd app/api/web/member/cmd/main.go
go build -o app/api/web/finance/cmd/cmd app/api/web/finance/cmd/main.go
go build -o app/api/web/order/cmd/cmd app/api/web/order/cmd/main.go
go build -o app/api/web/user/cmd/cmd app/api/web/user/cmd/main.go
go build -o app/api/web/groups/cmd/cmd app/api/web/groups/cmd/main.go
go build -o app/api/web/system/cmd/cmd app/api/web/system/cmd/main.go
go build -o app/api/web/wechat/cmd/cmd app/api/web/wechat/cmd/main.go
go build -o app/api/web/pay/cmd/cmd app/api/web/pay/cmd/main.go

go build -o app/service/main/payment/cmd/cmd app/service/main/payment/cmd/main.go
go build -o app/service/main/sysset/cmd/cmd app/service/main/sysset/cmd/main.go
go build -o app/service/main/groups/cmd/cmd app/service/main/groups/cmd/main.go

go build -o app/job/main/groups/cmd/cmd app/job/main/groups/cmd/main.go