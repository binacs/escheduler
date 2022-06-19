package core

//generate core mock code

//go:generate  mockgen -destination ../mock/core/mock_escheduler.go -package mock_core -source escheduler.go
//go:generate  mockgen -destination ../mock/core/mock_interface.go -package mock_core -source interface.go
