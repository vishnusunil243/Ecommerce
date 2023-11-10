package interfaces

import (
	"main.go/internal/common/helperStruct"
	"main.go/internal/common/response"
)

type SuperAdminUseCase interface {
	SuperLogin(superadmin helperStruct.SuperLoginReq) (string, error)
	CreateAdmin(admin helperStruct.CreateAdmin) (response.AdminData, error)
	ListAllAdmins() ([]response.AdminData, error)
	DisplayAdmin(id int) (response.AdminData, error)
	BlockAdmin(id int) (response.AdminData, error)
	BlockUser(id int) (response.UserData, error)
}
