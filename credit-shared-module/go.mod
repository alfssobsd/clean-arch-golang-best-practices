module clean-arch-golang-best-practices/credit-shared-module

require (
	clean-arch-golang-best-practices/credit-library v0.0.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

go 1.17

replace clean-arch-golang-best-practices/credit-library => ../credit-library
