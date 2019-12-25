package main

func initializeRoutes() {
	router.Use(setUserStatus())
	router.GET("/", showIndexPage)
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegisterPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn(), login)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", getArticle)
		articleRoutes.GET("/create", ensureLoggedIn(), getArticle)
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}
}