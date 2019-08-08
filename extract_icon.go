package main

import (
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "io/ioutil"
    "os"
)

func main() {
    db_host := os.Getenv("ISUBATA_DB_HOST")
    if db_host == "" {
        db_host = "127.0.0.1"
    }
    db_port := os.Getenv("ISUBATA_DB_PORT")
    if db_port == "" {
        db_port = "3306"
    }
    db_user := os.Getenv("ISUBATA_DB_USER")
    if db_user == "" {
        db_user = "root"
    }
    db_password := os.Getenv("ISUBATA_DB_PASSWORD")
    if db_password != "" {
        db_password = ":" + db_password
    }

    dsn := fmt.Sprintf("%s%s@tcp(%s:%s)/isubata?parseTime=true&loc=Local&charset=utf8mb4",
        db_user, db_password, db_host, db_port)

    db, _ := sqlx.Connect("mysql", dsn)

    defer db.Close()

    r, _ := db.Query("SELECT `name` , `data` FROM image")

    defer r.Close()

    var name string
    var data []byte
    for r.Next() {
        err := r.Scan(&name, &data)
        if err != nil {
            log.Fatalf(err.Error())
        }

        err = ioutil.WriteFile(name, data, 0666)
        if err != nil {
            log.Fatalf(err.Error())
        }
    }
}
