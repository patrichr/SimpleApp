module github.com/patrichr/SimpleApp/server

go 1.16

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.10.9
	github.com/patrichr/SimpleApp v0.0.0-20230607104238-685bf322c511
	google.golang.org/grpc v1.55.0
)

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab
