package cacheRepository

import (
	"fmt"
	"os"
	"stncCms/app/domain/entity"

	"stncCms/app/services"

	"github.com/hypnoglow/gormzap"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	_ "github.com/lib/pq" // here
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
)

var DB *gorm.DB

// Repositories strcut
type Repositories struct {
	User    services.UserAppInterface
	Lang    services.LanguageAppInterface
	Options services.OptionsAppInterface
	Media   services.MediaAppInterface
	Branch  services.BranchAppInterface
	Post    services.PostAppInterface
	Cat     services.CatAppInterface
	CatPost services.CatPostAppInterface
	Modules services.ModulesAppInterface
	DB      *gorm.DB
}

//DbConnect initial
/*TODO: burada db verisi pointer olarak işaretlenecek oyle gidecek veri*/
func DbConnect() *gorm.DB {
	dbdriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	gormAdvancedLogger := os.Getenv("GORM_ZAP_LOGGER")
	debug := os.Getenv("MODE")
	//	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword) //bu postresql

	//DBURL := "root:sel123C#@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local" //mysql
	var DBURL string

	if dbdriver == "mysql" {
		DBURL = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"
	} else if dbdriver == "postgres" {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", dbHost, dbPort, dbUser, dbPassword, dbName) //Build connection string
	}

	// dsn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s sslmode=disable",
	// HOST, PORT, username, password, database)

	db, err := gorm.Open(dbdriver, DBURL)
	db.Set("gorm:table_options", "charset=utf8")
	// }

	// db, err := gorm.Open(dbdriver, DBURL)
	//nunlar gorm 2 versionunda prfexi falan var
	// db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "krbn_", // table name prefix, table for `User` would be `t_users`
	// 		SingularTable: true,    // use singular table name, table for `User` would be `user` with this option enabled
	// 	},
	// 	// Logger: gorm_logrus.New(),
	// })

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	if debug == "DEBUG" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
		log := zap.NewExample()
		db.SetLogger(gormzap.New(log, gormzap.WithLevel(zap.DebugLevel)))
	} else if debug == "DEBUG" || debug == "TEST" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
	} else if debug == "RELEASE" {
		fmt.Println(debug)
		db.LogMode(false)
	}
	DB = db

	db.SingularTable(true)

	return db
}

//https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

// RepositoriesInit initial
func RepositoriesInit(db *gorm.DB) (*Repositories, error) {

	return &Repositories{
		User:    UserRepositoryInit(db),
		Lang:    LanguageRepositoryInit(db),
		Options: OptionRepositoryInit(db),
		Media:   MediaRepositoryInit(db),
		Branch:  BranchRepositoryInit(db),
		Post:    PostRepositoryInit(db),
		Cat:     CatRepositoryInit(db),
		CatPost: CatPostRepositoryInit(db),

		Modules: ModulesRepositoryInit(db),
		DB:      db,
	}, nil
}

//Close closes the  database connection
// func (s *Repositories) Close() error {
// 	return s.db.Close()
// }

// Automigrate This migrate all tables
func (s *Repositories) Automigrate() error {
	s.DB.AutoMigrate(&entity.Users{},
		&entity.Languages{}, &entity.Modules{}, &entity.Notes{},
		&entity.Options{}, &entity.Media{}, &entity.Region{}, &entity.Branches{}, &entity.Notification{}, &entity.NotificationTemplate{},
		&entity.Post{}, &entity.Categories{}, &entity.CategoryPosts{})
	return s.DB.Model(&entity.Branches{}).AddForeignKey("region_id", "br_region(id)", "CASCADE", "CASCADE").Error // one to many (one=region) (many=branches)
}
