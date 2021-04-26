package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Module struct {
	Id    int
	Description  string
}

func dbConn() (db *sql.DB) {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", uri)

	if err != nil {
			panic(err.Error())
	}
	return db
}

func DropTable() {
	db := dbConn()

	stmt, err := db.Prepare("DROP TABLE modules;")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func CreateTable() {
	log.Println("Preparing database...")

	db := dbConn()

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS modules (id int auto_increment primary key, description varchar(255));")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func Insert() {
	db := dbConn()

	descriptions := [5]string{
		"Padrões e Técnicas avançadas com Git e Github", 
		"Integração Contínua", 
		"Kubernetes", 
		"Observabilidade",
		"Fundamentos de Arquitetura de Software",
	}

	for i := 0; i < len(descriptions); i++ {
		description := descriptions[i]
		dml, err := db.Prepare("INSERT INTO modules(description) VALUES(?)")
	
		if err != nil {
			panic(err.Error())
		}
		
		dml.Exec(description)
	
		log.Println("INSERT: Description: " + description)
	}

	defer db.Close()
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("List modules from Mysql")

	db := dbConn()
	
	resultSet, err := db.Query("SELECT * FROM modules")

	if err != nil {
			panic(err.Error())
	}

	module := Module{}
	res := []Module{}

	for resultSet.Next() {
			var id int
			var description string
			
			err = resultSet.Scan(&id, &description)

			if err != nil {
					panic(err.Error())
			}
			
			module.Id = id
			module.Description = description
			res = append(res, module)
	}

  js, err := json.Marshal(res)

	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)

	defer db.Close()
}

func Start() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
	log.Println("Fullcycle PFA, started on: http://localhost:8000")
}

func main() {
	// DropTable()
	CreateTable()
	Insert()
	Start()
}