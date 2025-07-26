package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db1 *sql.DB
var db2 *sql.DB


type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func main(){
	// load .env file
	if err := godotenv.Load();
	err != nil{
		log.Fatal("Error loading .env file")
	}

	// mysql setup
	cfg1 := mysql.NewConfig();
	cfg1.User = os.Getenv("DBUSER")
	cfg1.Passwd = os.Getenv("DBPASS")
	cfg1.Net = "tcp"
	cfg1.Addr = "127.0.0.1:3306"
    cfg1.DBName = "first_recording"

	cfg2 := mysql.NewConfig();
	cfg2.User = os.Getenv("DBUSER")
	cfg2.Passwd = os.Getenv("DBPASS")
	cfg2.Net = "tcp"
	cfg2.Addr = "127.0.0.1:3306"
    cfg2.DBName = "secound_recording"

	var err error

	//establish connection for database One
	db1, err = sql.Open("mysql",cfg1.FormatDSN())
	if err != nil{
		log.Fatal(err)
	}
	// connected to db
	pingErrOne := db1.Ping()
	if pingErrOne != nil{
		log.Fatal(pingErrOne)
	}

	//establish connection for database Two
	db2, err = sql.Open("mysql",cfg2.FormatDSN())
	if err != nil{
		log.Fatal(err)
	}
	// connected to db
	pingErrTwo := db2.Ping()
	if pingErrTwo != nil{
		log.Fatal(pingErrTwo)
	}

	fmt.Printf("connected to both the DB\n")

	// sharding function envoke
	res, err:= albumsByIdSharded(3)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", res)

	// get albums by id
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

}

func albumsByArtist(name string) ([]Album, error){
	var albums []Album
	rows, err := db1.Query("SELECT * FROM album WHERE artist = ?", name);
	if err != nil{
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}		

	for rows.Next(){
		var alb Album;
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil{
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	return albums, nil;

}

func albumsByIdSharded(id int)([]Album, error){
	var dbToUse *sql.DB
    mod := id % 10
    if mod < 5 {
        dbToUse = db1
    } else if mod < 10 {
        dbToUse = db2
    } else {
        return nil, fmt.Errorf("invalid id for sharding: %d", id)
    }
    if dbToUse == nil {
        return nil, fmt.Errorf("database connection not initialized for id: %d", id)
    }
	var albums []Album
	rows, err := dbToUse.Query("SELECT * FROM album WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("albumsByIdSharded %d: %v", id, err)
	}

	for rows.Next(){
		var alb Album
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil{
			return nil, fmt.Errorf("albumsByIdSharded %q: %v", id, err)
		}
		albums = append(albums, alb)
	}

	return albums, nil;
}