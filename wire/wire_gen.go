// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"main.go/internal/infrastructure/config"
	"main.go/internal/infrastructure/persistence"
	"main.go/internal/repository"
	"main.go/internal/usecase"
	"main.go/internal/web"
	"main.go/internal/web/handler"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepo(gormDB)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	adminRepository := repository.NewAdminRepo(gormDB)
	adminUseCase := usecase.NewAdminUsecase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	productRepository := repository.NewProductRepo(gormDB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)
	superAdminRepository := repository.NewSuperRepo(gormDB)
	superAdminUseCase := usecase.NewSuperAdminUsecase(superAdminRepository)
	superAdminHandler := handler.NewSuperAdminHandler(superAdminUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, productHandler, superAdminHandler)
	return serverHTTP, nil
}
