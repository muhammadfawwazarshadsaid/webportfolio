package migration

import "webportfolio/database"


func Migration(){
	
	db := database.ConnectDB()
	defer db.Close()
}