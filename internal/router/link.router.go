package router

func (r *Router) LinksRouter() {
	r.home.GET("/:slug", r.lc.RedirectControl)

	link := r.router.Group("links")
	link.POST("", r.m.AuthMiddlewarer, r.lc.CreateShortURLcontrol)
	link.DELETE("/:slug", r.m.AuthMiddlewarer, r.lc.DelLinkRControl)
	link.GET("", r.m.AuthMiddlewarer, r.lc.GetMyLinksControl)
}
