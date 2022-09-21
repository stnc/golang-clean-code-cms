// golang gin framework mvc and clean code project
// Licensed under the Apache License 2.0
// @author Selman TUNÇ <selmantunc@gmail.com>
// @link: https://github.com/stnc/go-mvc-blog-clean-code
// @license: Apache License 2.0
package main

import (
	"github.com/flosch/pongo2/v5"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/leonelquinteros/gotext"
	myPongoGinRender "github.com/stnc/myPongoGinRender/v5"
	csrf "github.com/utrack/gin-csrf"
	"log"
	"net/http"
	"os"
	repository "stncCms/app/domain/repository/cacheRepository"
	"stncCms/app/web.api/controller/middleware"
	"stncCms/app/web/controller"
)

var cacheControlSelman = false

// https://github.com/stnc-go/gobyexample/blob/master/pongo2render/render.go
func init() {
	//To load our environmental variables.

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	/* //bu sunucuda çalışıyor
		    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	        if err != nil {
	            log.Fatal(err)
	        }
	        environmentPath := filepath.Join(dir, ".env")
	        err = godotenv.Load(environmentPath)
	        fatal(err)
	*/

}

func main() {

	// appEnv := os.Getenv("APP_ENV")
	// if appEnv == "local" {
	// 	err := beeep.Alert("Uygulama çalıştı", "Web Server Çalışmaya Başladı localhost:"+port, "assets/warning.png")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	//redis details

	// redisService, err := cache.RedisDBInit(redisHost, redisPort, redisPassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	debugMode := os.Getenv("MODE")

	//
	db := repository.DbConnect()
	services, err := repository.RepositoriesInit(db)
	if err != nil {
		panic(err)
	}
	//defer services.Close()
	services.Automigrate()

	// redisService, err := auth.RedisDBInit(redisHost, redisPort, redisPassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// token := auth.NewToken()

	posts := controller.InitPost(services.Post, services.CatPost, services.Cat, services.Lang, services.User)

	optionsHandle := controller.InitOptions(services.Options)

	userHandle := controller.InitUserControl(services.User, services.Branch, services.Role)

	login := controller.InitLogin(services.User)

	role := controller.InitRoles(services.Permission, services.Modules, services.Role, services.RolePermission)

	switch debugMode {
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)

	case "DEBUG":
		gin.SetMode(gin.DebugMode)

	case "TEST":
		gin.SetMode(gin.TestMode)

	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	store := cookie.NewStore([]byte("rodrigoHunter1"))
	////60 minutes olan 1 hours =  ( 60x60) 3600 seconds.
	//60 second * 60 = 1 saat //60*60
	//3600 (1 hour ) * 5 = 5 hour
	store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: 3600 * 8}) //Also set Secure: true if using SSL, you should though
	// store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: -1}) //Also set Secure: true if using SSL, you should though

	r.Use(sessions.Sessions("myCRM", store))

	r.Use(middleware.CORSMiddleware()) //For CORS

	r.Use(csrf.Middleware(csrf.Options{
		Secret: "rodrigoHunter2",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.HTMLRender = myPongoGinRender.TemplatePath("public/view")

	r.MaxMultipartMemory = 1 >> 20 // 8 MiB

	//TODO: html template make
	r.NoRoute(func(c *gin.Context) {
		var getText *gotext.Locale
		getText = gotext.NewLocale("public/locales", "tr_TR")
		getText.AddDomain("l404")
		viewData := pongo2.Context{
			"title":  "404",
			"locale": getText,
		}
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			"admin/viewComponents/404.html",
			viewData,
		)
	})

	r.Static("/assets", "./public/static")

	r.StaticFS("/upload", http.Dir("./public/upl"))
	//r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("/", controller.Index)
	r.GET("admin", controller.Index)
	r.GET("admin/", controller.Index)

	r.GET("optionsDefault", controller.OptionsDefault)
	r.GET("cacheReset", controller.CacheReset)

	optionsGroup := r.Group("/admin/options")
	{
		optionsGroup.GET("/", optionsHandle.Index)
		optionsGroup.POST("update", optionsHandle.Update)
		optionsGroup.GET("receiptNo", optionsHandle.ReceiptNo)
	}

	roleGroup := r.Group("/admin/roles")
	{
		roleGroup.GET("/knockout", role.IndexKnockout)
		roleGroup.GET("/", role.Index)
		roleGroup.GET("/create", role.Create)
		roleGroup.POST("/store", role.Store)
		roleGroup.GET("/edit/:ID", role.Edit)
		roleGroup.POST("/update", role.Update)
		roleGroup.GET("delete/:ID", role.Delete)
	}

	adminPost := r.Group("/admin/post")
	{
		adminPost.GET("/", posts.Index)
		adminPost.GET("index", posts.Index)
		adminPost.GET("create", posts.Create)
		adminPost.POST("store", posts.Store)
		adminPost.GET("edit/:postID", posts.Edit)
		adminPost.POST("update", posts.Update)
		adminPost.POST("upload", posts.Upload)
	}

	loginGroup := r.Group("/admin/login")
	{
		loginGroup.GET("/", login.Login)
		//loginGroup.GET("password", login.SifreVer)
		loginGroup.POST("loginpost", login.LoginPost)
		loginGroup.GET("logout", login.Logout)
	}

	userGroup := r.Group("/admin/user")
	{
		userGroup.GET("/", userHandle.Index)
		userGroup.GET("index", userHandle.Index)
		userGroup.GET("create", userHandle.Create)
		userGroup.POST("store", userHandle.Store)
		userGroup.GET("edit/:UserID", userHandle.Edit)
		// userGroup.GET("delete/:ID", userHandle.Delete)
		userGroup.POST("update", userHandle.Update)
		userGroup.GET("NewPasswordModalBox", userHandle.NewPasswordModalBox)
		userGroup.POST("NewPasswordAjax", userHandle.NewPasswordCreateModalBox)
	}
	//api routes
	// v1 := r.Group("/api/v1")
	// {
	// 	v1.POST("users", usersAPI.SaveUser)
	// 	v1.GET("users", usersAPI.GetUsers)

	// 	v1.GET("users/:user_id", usersAPI.GetUser)
	// 	v1.GET("postall", postsAPI.GetAllPaginationost)
	// 	v1.POST("post", postsAPI.SavePost)
	// 	v1.PUT("post/:post_id", middleware.AuthMiddleware(), postsAPI.UpdatePost)
	// 	v1.GET("post/:post_id", postsAPI.GetPostAndCreator)
	// 	v1.DELETE("post/:post_id", middleware.AuthMiddleware(), postsAPI.DeletePost)
	// 	// cs.GET("/allcoins", controller.AllCoins())
	// 	// cs.GET("/mycoins/:id", controller.MyCoins())
	// 	// cs.GET("/create", controller.CreateCoin())
	// 	// cs.POST("/store", controller.StoreCoin())
	// 	// cs.GET("/edit/:id", controller.EditCoin())
	// 	// cs.PUT("/update/:id", controller.UpdateCoin())
	// 	// e.GET("/cpr/:slug", controller.CoinPreview())
	// 	// cs.DELETE("/:id", controller.DeleteCoin())
	// 	//authentication routes
	// 	v1.POST("login", authenticate.Login)
	// 	v1.POST("logout", authenticate.Logout)
	// 	v1.POST("refresh", authenticate.Refresh)
	// }

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//Starting the application
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8080" //localhost
	}
	log.Fatal(r.Run(":" + appPort))

}
