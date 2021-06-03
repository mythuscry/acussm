package dtoo

import "acussm/demo/model"

type userDto struct {
	Name string `json:"name"`
	Telephopne string `json:"telephopne"`
}

func Touser(user model.User) userDto {
	return userDto{
		Name: user.Name,
		Telephopne: user.Telephone,
	}
}