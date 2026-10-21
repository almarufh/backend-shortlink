package router

import "os"

func (r *Router) staticRouter() {
	r.router.Static(os.Getenv("PROFILE_URL"), "public/photo_profiles")
}
