package leveldb

import (
    "github.com/syndtr/goleveldb/leveldb"
    "fmt"
)

/*Création de la base de donnée "path" précisée dans le main*/

func Create_database_leveldb(path string) {
   
  		fmt.Println("Création LevelDB: ")

        	c, err := leveldb.OpenFile(path, nil)
	
		if err != nil {
			fmt.Println("Error creating LevelDB database: ", err.Error())
		}
		defer c.Close()
}

