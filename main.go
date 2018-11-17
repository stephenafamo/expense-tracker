package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("EXPT") // will be uppercased automatically
	viper.AutomaticEnv()
}

func main() {

	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	dbname := viper.GetString("DB_NAME")

	connString := "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " sslmode=disable"

	DB, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	s := NewServer(DB)
	s.addAdminUser()

	http.Handle("/", s.router)
	fmt.Println("listening on port 80")
	log.Println(http.ListenAndServe(":80", s.router))
}

func checkError(err error) {
	if err != nil {
		log.Print(err)
		fmt.Printf("%#v", err)
	}
}
