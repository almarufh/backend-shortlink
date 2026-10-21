package router

import "os"

func (r *Router) StaticRouter() {
	r.router.Static(os.Getenv("PROFILE_URL"), "public/photo_profiles")
}
