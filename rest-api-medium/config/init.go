package config

import (
   "github.com/rynfkn/rest-api-medium/app/controller"
   "github.com/rynfkn/rest-api-medium/app/repository"
   "github.com/rynfkn/rest-api-medium/app/service"
)

type Initialization struct {
   userRepo repository.UserRepository
   userSvc  service.UserService
   UserCtrl controller.UserController
   RoleRepo repository.RoleRepository
}

func NewInitialization(userRepo repository.UserRepository,
   userService service.UserService,
   userCtrl controller.UserController,
   roleRepo repository.RoleRepository) *Initialization {
   return &Initialization{
      userRepo: userRepo,
      userSvc:  userService,
      UserCtrl: userCtrl,
      RoleRepo: roleRepo,
   }
}