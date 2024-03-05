package controller

import (
	"BelajarAPI/controller/user"
	user_model "BelajarAPI/model/user"

	"BelajarAPI/controller/activity"
	activity_model "BelajarAPI/model/activity"
)

func User(m *user_model.UserModel) *user.UserController {
	return &user.UserController{Model: *m}
}

func Activity(m *activity_model.ActivityModel) *activity.ActivityController {
	return &activity.ActivityController{Model: *m}
}
