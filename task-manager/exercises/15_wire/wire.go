//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

var ConfigSet = wire.NewSet(NewConfig)

var DatabaseSet = wire.NewSet(NewDatabase, NewCache)

var RepositorySet = wire.NewSet(NewUserRepository)

var ServiceSet = wire.NewSet(NewUserService)

var AppSet = wire.NewSet(
	ConfigSet,
	DatabaseSet,
	RepositorySet,
	ServiceSet,
)



func InitializeUserService() *UserService {
	wire.Build(AppSet)
	return nil
}