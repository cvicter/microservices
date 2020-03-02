module order

go 1.12

require (
	github.com/go-redis/redis/v7 v7.2.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71 // indirect
	microservices/order/db v0.0.0 // indirect
	microservices/order/queue v0.0.0 // indirect
)

replace microservices/order/queue v0.0.0 => ./queue

replace microservices/order/db v0.0.0 => ./db
