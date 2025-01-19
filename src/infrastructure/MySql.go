package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Cambia estos valores según tu configuración
	username := "root"
	password := "1234"
	hostname := "127.0.0.1:3306"
	dbname := "Go"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verifica la conexión
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Conexión exitosa a MySQL")
	return db, nil
}
