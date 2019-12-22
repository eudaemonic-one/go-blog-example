package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}