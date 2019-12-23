package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/register", showRegisterPage)
		userRoutes.POST("/register", register)
		userRoutes.GET("/login", showLoginPage)
		userRoutes.POST("/login", login)
		userRoutes.GET("/logout", logout)
	}
}