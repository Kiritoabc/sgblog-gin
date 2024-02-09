package admin

type RouterGroup struct {
	LoginRouter
	ArticleRouter
	CategoryRouter
	LinkRouter
	TagRouter
	UserRouter
	MenuRouter
	RoleRouter
}
