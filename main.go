package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "192.168.40.135"
	port     = 5432
	user     = "pgadmin"
	password = "159753"
	dbname   = "3dconst"
)

var db *sql.DB

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/getSeriesPrice", getSeriesPrice)
	router.Run("0.0.0.0:8100")

}

type product struct {
	id     int
	name   string
	code   string
	price  int
	name1c string
}

func getSeriesPrice(c *gin.Context) {

	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query(`SELECT id, name, code, price, name1c  FROM public."API_priceseries" WHERE name='экспресс/2400/2400/600/дубмолочный/3-дверныйшкаф/стандарт`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	product_ := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.name, &p.code, &p.price, &p.name1c)
		if err != nil {
			fmt.Println(err)
			continue
		}
		product_ = append(product_, p)
	}
	for i, p := range product_ {
		fmt.Println(p, i)
	}

}
