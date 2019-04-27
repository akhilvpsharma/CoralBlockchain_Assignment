package main
import
(
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func SaveService(webForm WebForm) (string, string){
	fmt.Println("Saving...")
	fmt.Printf(webForm.UserName)
	fmt.Printf(webForm.EmailID)
	fmt.Printf(webForm.PhoneNumber)
	fmt.Printf(webForm.Password)
	db, err := sql.Open("mysql", "dummyUser:dummyUser01@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/db_intern")
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()

	searchQuery:="SELECT * FROM db_intern.userData WHERE emailId='"+webForm.EmailID+"';"
	results, err := db.Query(searchQuery)
	if err != nil {
			panic(err.Error())
	}
	
	var flag =false
	var x WebForm
	for results.Next() {
		flag=true
		err = results.Scan(&x.UserName, &x.EmailID, &x.PhoneNumber, &x.Password, &x.DateTime)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(x.UserName)
	}
	
	if(flag==false){	
		insertQuery:="INSERT INTO `db_intern`.`userData` (`username`, `emailId`, `phoneNo`, `password`, `dateTime`) VALUES ('"+webForm.UserName+"', '"+webForm.EmailID+"', '"+webForm.PhoneNumber+"', '"+webForm.Password+"', NOW());"
		insert, err := db.Query(insertQuery)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		return "true", "Inserted Record"
	}else{
		updateQuery:="UPDATE `db_intern`.`userData` SET `username`='"+webForm.UserName+"', `emailId`='"+webForm.EmailID+"', `phoneNo`='"+webForm.PhoneNumber+"', `password`='"+webForm.Password+"', `dateTime`=NOW() WHERE `emailId`='"+x.EmailID+"'";
		update, err := db.Query(updateQuery)

		if err != nil {
			panic(err.Error())
		}

		defer update.Close()
		return "true", "Updated Exisiting Record"
	}	
}

func SearchService(emailId string) (WebForm,string, string, error){
	fmt.Println("Searching...", emailId)
	db, err := sql.Open("mysql", "dummyUser:dummyUser01@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/db_intern")
    
    if err != nil {
        panic(err.Error())
	}
	defer db.Close()
	
	searchQuery:="SELECT * FROM db_intern.userData WHERE emailId='"+emailId+"';"
	results, err := db.Query(searchQuery)
	if err != nil {
			panic(err.Error())
	}

	var x WebForm
	for results.Next() {
		err = results.Scan(&x.UserName, &x.EmailID, &x.PhoneNumber, &x.Password, &x.DateTime)
		if err != nil {
			panic(err.Error())
		}
		return x, "true", "true", nil
	}
	return x, "false", "Record Doesn't Exist", err
}

func DeleteService(emailId string) (string, string, error){
	fmt.Println("Deleting...", emailId)
	db, err := sql.Open("mysql", "dummyUser:dummyUser01@tcp(db-intern.ciupl0p5utwk.us-east-1.rds.amazonaws.com:3306)/db_intern")
    if err != nil {
        panic(err.Error())
	}
	defer db.Close()
	
	searchQuery:="SELECT * FROM db_intern.userData WHERE emailId='"+emailId+"';"
	results, err := db.Query(searchQuery)
	if err != nil {
			panic(err.Error())
	}

	var flag =false
	for results.Next() {
		flag=true
		break
	}
	if(flag==true){
		deleteQuery:="DELETE FROM db_intern.userData WHERE emailId='"+emailId+"';"
		db.Query(deleteQuery)
		return "true", "Deleteted record successfully", err
	}
	return "false", "Record Doesn't Exist", err
}