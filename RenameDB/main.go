// In this we are trying to create renama the existing database available in Postgres
// So we will take below input from user
// Host,dbname,user,pass,newname

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



// Main function to create extension
func main(){
  // take the varibale from user
     fmt.Println("Enter the Hostname:")

     // var then variable name then variable type
     var host string

     // Taking input from user
     fmt.Scanln(&host)
     fmt.Println("Enter the Current Database Name: ")
     var dbname string
     fmt.Scanln(&dbname)
     fmt.Println("Enter the User Name: ")
     var user string
     fmt.Scanln(&user)
     fmt.Println("Enter the Password Name: ")
     var pgpass string
     fmt.Scanln(&pgpass)
     fmt.Println("Enter the New Database Name: ")
     var newdbname string
     fmt.Scanln(&newdbname)
     // Addition of two string
     fmt.Printf("Hostname:%s  DBname:%s  User:%s  Password:%s  Extension:%s", host, dbname, user, pgpass, newdbname)
     c := askForConfirmation("\n\nPlease validate the above DB details & Press y/yes to continue?")
     if c {
       fmt.Println("Details not are not validate.")
       return
     }
     fmt.Println("Connecting to the database")
     psqlInfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=5432 sslmode=disable", host, dbname, user, pgpass)
     db, err := sql.Open("postgres",psqlInfo)
      if err != nil {
        panic (err)
      }
     err = db.Ping()
      if err != nil {
        fmt.Println("Failed to Connect Database")
      }
     fmt.Println("Connected Sucessfully !!!!")
     //let's prepare the query to be excuete
     query, err := db.Prepare(fmt.Sprintf(`ALTER DATABASE %s RENAME TO %s;`,dbname, newdbname))
      if err != nil {
        log.Println(err)
      }
     _, err = query.Exec()
      if err != nil {
        fmt.Println(err)
      }else{
        fmt.Printf("Database Has been Rename  Successfully...!!!! ")
      }

}

