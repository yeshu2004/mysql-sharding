package main

import (
	"database/sql"
	"fmt"
	"hash/fnv"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var db1 *sql.DB
var db2 *sql.DB

type Album struct {
	ID     string
	Title  string
	Artist string
	Price  float32
}

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// mysql setup
	cfg1 := mysql.NewConfig()
	cfg1.User = os.Getenv("DBUSER")
	cfg1.Passwd = os.Getenv("DBPASS")
	cfg1.Net = "tcp"
	cfg1.Addr = "127.0.0.1:3306"
	cfg1.DBName = "first_recording"

	cfg2 := mysql.NewConfig()
	cfg2.User = os.Getenv("DBUSER")
	cfg2.Passwd = os.Getenv("DBPASS")
	cfg2.Net = "tcp"
	cfg2.Addr = "127.0.0.1:3306"
	cfg2.DBName = "secound_recording"

	var err error

	//establish connection for database One
	db1, err = sql.Open("mysql", cfg1.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	// connected to db
	pingErrOne := db1.Ping()
	if pingErrOne != nil {
		log.Fatal(pingErrOne)
	}

	//establish connection for database Two
	db2, err = sql.Open("mysql", cfg2.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	// connected to db
	pingErrTwo := db2.Ping()
	if pingErrTwo != nil {
		log.Fatal(pingErrTwo)
	}

	fmt.Printf("connected to both the DB\n")

	// get albums
	res, err:= getAlbum("1acea3af-a752-4844-8d41-91d6d7025922")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", res)
	
	// fing album by artist
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// change the value 
	newAlbum := Album{
		Title:  "Kind of Blue",
		Artist: "Miles Davis",
		Price:  42.50,
	}

	albumTitle, err := addAlbum(newAlbum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album added: %v\n", albumTitle)

}

func getShardById(id string) (*sql.DB, error) {
	h := fnv.New32a()
	h.Write([]byte(id))
	hash := h.Sum32()
	if hash%2 == 0 {
		return db1, nil
	} else {
		return db2, nil
	}
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	// Query both shards since artist could exist in either i.e this makes this bad as if both db has 100k & 200k
	for _, db := range []*sql.DB{db1, db2} {
		rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
		if err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		defer rows.Close()

		for rows.Next() {
			var alb Album
			err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
			if err != nil {
				return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
			}
			albums = append(albums, alb)
		}
	}

	return albums, nil
}

func getAlbum(id string) ([]Album, error) {
	dbToUse, err := getShardById(id)
	if err != nil {
		return nil, fmt.Errorf("database connection not initialized for id: %v", id)
	}
	var albums []Album
	rows, err := dbToUse.Query("SELECT * FROM album WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("albumsByIdSharded %v: %v", id, err)
	}

	for rows.Next() {
		var alb Album
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			return nil, fmt.Errorf("albumsByIdSharded %q: %v", id, err)
		}
		albums = append(albums, alb)
	}

	return albums, nil
}

func addAlbum(alb Album) (string, error) {
	id := uuid.New().String()
	db, err := getShardById(id)
	if err != nil {
		return "", fmt.Errorf("database connection not initialized for id: %v", id)
	}
	alb.ID = id
	result, err := db.Exec("INSERT INTO album (id, title, artist, price) VALUES (?,?,?,?)", alb.ID, alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return "", fmt.Errorf("addAlbum %s: %v", alb.Title, err)
	}

	effRow, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("addAlbum %s: failed to get rows affected: %v", alb.Title, err)
	}

	if effRow == 0 {
		return "", fmt.Errorf("addAlbum %s: no rows inserted", alb.Title)
	}

	return alb.Title, nil
}
