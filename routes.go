package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/register", showRegisterPage)
		userRoutes.POST("/register", register)
		userRoutes.GET("/login", showLoginPage)
		userRoutes.POST("/login", login)
		userRoutes.GET("/logout", logout)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", getArticle)
		articleRoutes.GET("/create", getArticle)
		articleRoutes.POST("/create", createArticle)
	}
}