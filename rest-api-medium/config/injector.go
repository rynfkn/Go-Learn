package config

import (
	"github.com/google/wire"
	"github.com/rynfkn/rest-api-medium/app/controller"
	"github.com/rynfkn/rest-api-medium/app/repository"
	"github.com/rynfkn/rest-api-medium/app/service"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, roleRepoSet)
	return nil
}
