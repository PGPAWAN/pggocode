package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
  "database/sql"
  _ "github.com/lib/pq"
)

// In this we are trying to create any pre defined extension available in postgres database.
// So we will take below input from user
// Host,dbname,user,pass


func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "n" || response == "no" {
			return true
		} else if response == "y" || response == "yes" {
			return false
		}
	}
}

func DBConnect(host,dbname,user,pgpass) (db *sql.DB, err error){
var psqlInfo
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " "password=%s dbname=%s ", host, port, user, password, dbname)
db, err := sql.Open("postgres", psqlInfo)
if err != nil {
  panic(err)
}
err = db.Ping()
if err != nil {
  log.Print("Failed to Connected Database")
  return err
}

fmt.Println("Successfully connected!")
return db

}

// Main function to create extension
func main(){
  // take the varibale from user
     fmt.Println("Enter the Hostname:")

     // var then variable name then variable type
     var host string

     // Taking input from user
     fmt.Scanln(&host)
     fmt.Println("Enter the Database Name: ")
     var dbname string
     fmt.Scanln(&dbname)
     fmt.Println("Enter the User Name: ")
     var user string
     fmt.Scanln(&user)
     fmt.Println("Enter the Password Name: ")
     var pgpass string
     fmt.Scanln(&pgpass)
     fmt.Println("Enter the Extension Name: ")
     var extname string
     fmt.Scanln(&extname)
     // Addition of two string
     fmt.Printf("Hostname:%s  DBname:%s  User:%s  Password:%s  Extension:%s", host, dbname, user, pgpass, extname)
     c := askForConfirmation("\n\nPlease validate the above DB details & Press y/yes to continue?")
     if c {
       fmt.Println("Details not are not validate.")
       return
     }
    
     fmt.Println("Connecting to database.... ")
     db=dbconnect(host,dbname,user,pgpass)
     if err !=nil {
      fmt.Println("Connection broken!!!!!!")
     }else{
      rows,err := db.Query('show server_version;')
      if err !=nil {
        fmt.Println(err)
      }
      for rows.Next(){
        err =rows.Scan(&version)
      }
      fmt.Printf(name)
  
     }

     

}

