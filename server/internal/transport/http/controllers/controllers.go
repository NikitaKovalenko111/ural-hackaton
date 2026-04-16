package controllers

import (
	"ural-hackaton/internal/services"
	admin_controller "ural-hackaton/internal/transport/http/controllers/admin"
	hub_controller "ural-hackaton/internal/transport/http/controllers/hub"
	mentor_controller "ural-hackaton/internal/transport/http/controllers/mentor"
	request_controller "ural-hackaton/internal/transport/http/controllers/request"
	user_controller "ural-hackaton/internal/transport/http/controllers/user"
)

type Controllers struct {
	AdminController   *admin_controller.AdminController
	UserController    *user_controller.UserController
	MentorController  *mentor_controller.MentorController
	HubController     *hub_controller.HubController
	RequestController *request_controller.RequestController
}

func Init(svc *services.Services) *Controllers {
	var userService user_controller.UserService
	if svc != nil {
		userService = svc.UserService
	}

	return &Controllers{
		AdminController:   admin_controller.Init(nil),
		UserController:    user_controller.Init(userService),
		MentorController:  mentor_controller.Init(nil),
		HubController:     hub_controller.Init(nil),
		RequestController: request_controller.Init(nil),
	}
}
