package main
import
(
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
) 	

func main() {
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_intern")
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }
    // defer the close till after the main function has finished
    // executing
    defer db.Close()
    
    // Execute the query
    results, err := db.Query("SELECT * FROM db_intern.userData;")
    if err != nil {
         panic(err.Error()) // proper error handling instead of panic in your app
    }
 
    for results.Next() {
        var x WebForm
        // for each row, scan the result into our tag composite object
        err = results.Scan(&x.UserName, &x.EmailID, &x.PhoneNumber, &x.Password, &x.DateTime)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        fmt.Println(x.UserName)
    }
    Serve()
}