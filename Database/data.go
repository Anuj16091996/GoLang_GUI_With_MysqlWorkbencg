package Database

import (
	"Final_Project/Model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnect() sql.DB {

	db, err := sql.Open("mysql", "farmer:farmer@tcp(localhost)/ferme_bd")

	if err != nil {
		panic(err.Error())
	}

	return *db

}

func CloseDatabase(db *sql.DB) {
	defer db.Close()
}

func Print() {
	db := GetConnect()
	results, err := db.Query("SELECT fname FROM bank.customer")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag string

		err = results.Scan(&tag)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Println(tag)

	}

	CloseDatabase(&db)
}



func CheckLoginCredt(id string, password string) bool{
	db := GetConnect()
	results, err := db.Query("SELECT custid,Online_Password FROM bank.customer");

	if err != nil {
		panic(err.Error())
	}


	for results.Next() {

		var tag Model.Body

		err = results.Scan(&tag.Custometr_ID,&tag.Online_Password)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		if id ==tag.Custometr_ID{
			if password==tag.Online_Password{
				return true
			}
			CloseDatabase(&db)
			return false
		}

	}

	CloseDatabase(&db)



	return false;
}


func CustomerWindow(Customer_id string) Model.Body{
	var tag Model.Body

	db := GetConnect()
	results, err := db.Query("select distinct c.custid, c.fname, c.ltname,c.Online_Password, c.opening_balance  from  bank.customer c where c.custid=?",Customer_id);

	if err != nil {
		panic(err.Error())
	}

	for results.Next(){
		err = results.Scan(&tag.Custometr_ID,&tag.First_Name,&tag.Last_Name,&tag.Online_Password,&tag.Balace)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}


	CloseDatabase(&db)

	return tag;

}


func UpdateBankBalance(tag Model.Body){
	db := GetConnect()
	_, err := db.Query("update bank.customer set opening_balance=? where custid=?",tag.Balace,tag.Custometr_ID);

	if err != nil {
		panic(err.Error())
	}

	CloseDatabase(&db)
}


func CheckNumber(number int) bool{

	db := GetConnect()
	n, err := db.Query("SELECT mobileno from bank.customer where mobileno=?",number);

	if err != nil {
		panic(err.Error())
	}
	CloseDatabase(&db)



	for n.Next(){
		return true
	}
		return false


}

func InsertUser(tag Model.Body) string{
	db := GetConnect()
	_, err := db.Query("INSERT INTO bank.customer(fname,ltname,Online_Password,opening_balance,mobileno) VALUES (?,?,?,?,?)",tag.First_Name,tag.Last_Name,tag.Online_Password,tag.Balace,tag.MobileNumber);

	if err != nil {
		panic(err.Error())
	}
	CloseDatabase(&db)


	value:=GiveCustomerId(tag.MobileNumber)
	return value

}


func GiveCustomerId(number int) string{
	db := GetConnect()
	n, err := db.Query("SELECT custid from bank.customer where mobileno=?",number);

	if err != nil {
		panic(err.Error())
	}

	var value string
	for n.Next(){
		err=n.Scan(&value)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	CloseDatabase(&db)
	return value
}


