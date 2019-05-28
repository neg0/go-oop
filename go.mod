module github.com/neg0/go-oop

require (
	github.com/neg0/go-oop/polymorphism/pay v0.0.0
	github.com/neg0/go-oop/polymorphism/payment v0.0.0 // indirect
)

replace (
	github.com/neg0/go-oop/polymorphism/pay => ./src/polymorphism/pay
	github.com/neg0/go-oop/polymorphism/payment => ./src/polymorphism/payment
)

go 1.12
