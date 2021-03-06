package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		Authentication: false,
	},
	{
		Uri:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.SearchByNickOrName,
		Authentication: false,
	},
	{
		Uri:            "/users/{userID}",
		Method:         http.MethodGet,
		Function:       controllers.GetByID,
		Authentication: false,
	},
	{
		Uri:            "/users/{userID}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		Authentication: false,
	},
	{
		Uri:            "/users/{userID}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		Authentication: false,
	},
}
