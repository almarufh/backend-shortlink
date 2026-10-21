package router

func (r *Router) AuthRouter() {
	auth := r.router.Group("auth")
	auth.POST("register", r.ac.RegisterAC)
	auth.POST("login", r.ac.LoginAC)
	auth.DELETE("logout", r.m.AuthMiddlewarer, r.ac.LogoutAC)
}
