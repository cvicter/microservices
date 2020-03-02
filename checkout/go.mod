module checkout

go 1.12

require (
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71 // indirect
	microservices/checkout/queue v0.0.0 // indirect
)

replace microservices/checkout/queue v0.0.0 => ./queue
