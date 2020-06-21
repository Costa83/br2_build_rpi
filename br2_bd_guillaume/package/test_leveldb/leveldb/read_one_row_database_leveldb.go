package leveldb

import (
    "github.com/syndtr/goleveldb/leveldb"
    "fmt"
)
/*Lecture d'une ligne "nb_row_to_print" de la base de donnée "path" précisée dans le main avec un flag "print" permetant d'activer ou non l'affichage de la base de donnée sur le terminal*/
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
