module github.com/neg0/go-oop

require (
	github.com/neg0/go-oop/src/channels/credit_account v0.0.0
	github.com/neg0/go-oop/src/composition/account v0.0.0
	github.com/neg0/go-oop/src/composition/credit_account v0.0.0
	github.com/neg0/go-oop/src/encapsulation/interfaces v0.0.0
	github.com/neg0/go-oop/src/encapsulation/interfaces/payment v0.0.0
	github.com/neg0/go-oop/src/encapsulation/package_oriented_design v0.0.0
	github.com/neg0/go-oop/src/polymorphism/pay v0.0.0
)

replace (
	github.com/neg0/go-oop/src/channels/credit_account => ./src/channels/credit_account
	github.com/neg0/go-oop/src/composition/account => ./src/composition/account
	github.com/neg0/go-oop/src/composition/credit_account => ./src/composition/credit_account
	github.com/neg0/go-oop/src/encapsulation/interfaces => ./src/encapsulation/interfaces
	github.com/neg0/go-oop/src/encapsulation/interfaces/payment => ./src/encapsulation/interfaces/payment
	github.com/neg0/go-oop/src/encapsulation/package_oriented_design => ./src/encapsulation/package_oriented_design
	github.com/neg0/go-oop/src/polymorphism/pay => ./src/polymorphism/pay
	github.com/neg0/go-oop/src/polymorphism/payment => ./src/polymorphism/payment
)

go 1.12
