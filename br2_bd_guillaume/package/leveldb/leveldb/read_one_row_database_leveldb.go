package leveldb

import (
    "github.com/syndtr/goleveldb/leveldb"
    "fmt"
)

func Read_one_row_database_leveldb(nb_row_to_print string, path string, print string) {

  fmt.Println("Read one row nb#:  ", nb_row_to_print)
  db, err := leveldb.OpenFile(path, nil)
	
	if err != nil {
		
		fmt.Println("Error openning LevelDB database: ", err.Error())
	}
	defer db.Close()
	
	
	data, err := db.Get([]byte(nb_row_to_print), nil)

  	if print == "true" {
  		data_string := string(data)
		fmt.Println("Read one row nb#:  ", data_string)
	}
  
  
  
  
}
