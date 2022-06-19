package framework

//generate framework mock code

//go:generate  mockgen -destination ../mock/framework/mock_interface.go -package mock_framework -source interface.go
//go:generate  mockgen -destination ../mock/framework/mock_relation.go -package mock_framework -source relation.go
