package router

func (r *Router) AuthRouter() {
	auth := r.router.Group("auth")
	auth.POST("register", r.as.RegisterAC)
	auth.POST("login", r.as.LoginAC)
	auth.DELETE("logout", r.m.AuthMiddlewarer, r.as.LogoutAC)
}
