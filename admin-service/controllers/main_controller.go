package controllers

type MainController struct {
	User interface{ UserController }
	App  interface{ AppController }
}
