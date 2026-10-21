package router

func (r *Router) LinksRouter() {
	link := r.router.Group("links")
	link.POST("", r.m.AuthMiddlewarer, r.lc.CreateShortURLcontrol)
	r.home.GET("/:slug", r.lc.RedirectControl)
}
