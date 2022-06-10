package main

import (
	"log"
	"os"
	"simple-blog/database"
	"simple-blog/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error .env files")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}

// var rnd *renderer.Render
// var db *my.Database
// var dsn = "root:@tcp(127.0.0.1:3306)/simple-blog?charset=utf8mb4"
// var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// type SimpleBlogModel struct {
// 	Name string
// 	Year string
// }

// func main() {
// 	http.HandleFunc("/createstuff", GoDatabaseCreate)

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func GoDatabaseCreate(w http.ResponseWriter, r *http.Request) {
// 	SimpleBlogModel := SimpleBlogModel{
// 		Name: "Jefta",
// 		Year: "2020",
// 	}
// 	db.Create(&SimpleBlogModel)
// 	if err := db.Create(&SimpleBlogModel).Error; err != nil {
// 		log.Fatalln((err))
// 	}

// 	json.NewEncoder(w).Encode(SimpleBlogModel)

// 	fmt.Println("Fields Added", SimpleBlogModel)
// }

// const (
// 	hostName       string = "127.0.0.1:3306"
// 	dbName         string = "simple-blog"
// 	collectionName string = "blog"
// 	port           string = ":8080"
// )

// type (
// 	postModel struct {
// 		ID 		int
// 		Title	string
// 		Post	string
// 		Author	string
// 		CreatedAt time.Time
// 		UpdatedAt time.Time
// 	}

// 	post struct{
// 		ID 		int `json:"id"`
// 		Title	string `json:"title"`
// 		Post	string `json:"post"`
// 		Author	string `json:"author"`
// 		CreatedAt time.Time `json:"created_at"`
// 		UpdatedAt time.Time `json:"updated_at"`
// 	}
// )

// func init()  {
// 	rnd = renderer.New()
// 	sess, err:= my.Dial(hostName)
// 	checkErr(err)
// 	sess.SetMode(my.Monotonic, true)
// }
