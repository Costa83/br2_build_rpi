package leveldb

import (
    "github.com/syndtr/goleveldb/leveldb"
    "fmt"
)


func Delete_database_leveldb(path string) {

  db, err := leveldb.OpenFile(path, nil)
	
	if err != nil {
		
		fmt.Println("Error creating LevelDB database: ", err.Error())
	}
	defer db.Close()
	
	iter := db.NewIterator(nil, nil)
        
        for iter.Next() {
        	
        	key := iter.Key()
   		key_string := string(key)
	 	err = db.Delete([]byte(key_string), nil)

        		
        		
	}
	
	
  	iter.Release()
  	err = iter.Error()


}
