package main

import (
    "github.com/rynfkn/rest-api-medium/config"
    "github.com/rynfkn/rest-api-medium/app/controller"
    "github.com/rynfkn/rest-api-medium/app/repository"
    "github.com/rynfkn/rest-api-medium/app/service"
    "github.com/rynfkn/rest-api-medium/app/router"
    "log"
)

func main() {
    db := config.ConnectToDB()
    if db == nil {
        log.Fatal("Database connection failed.")
    }

    // Initialize repositories
    userRepo := repository.UserRepositoryInit(db)
    roleRepo := repository.RoleRepositoryInit(db)

    // Initialize services
    userService := service.NewUserService(userRepo)

    // Initialize controllers
    userCtrl := controller.NewUserController(userService)

    // Initialize configuration with all components
    init := config.NewInitialization(userRepo, userService, userCtrl, roleRepo)

    // Check if init is nil for debugging
    if init == nil {
        log.Fatal("Initialization struct is nil.")
    }

    // Set up router with initialized config
    r := router.Init(init)
    r.Run(":3000")
}
