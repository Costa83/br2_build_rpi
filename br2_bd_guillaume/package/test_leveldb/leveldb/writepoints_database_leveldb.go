package leveldb

import (
    "github.com/syndtr/goleveldb/leveldb"
    "fmt"
    "strconv"
    "math/rand"
)

/*Ecriture en base de données de n lignes "nb_row" dans la base de donnée nommée "path "*/
func Write_database_leveldb(nb_row *string, path *string){

	    fmt.Println("Insetion de %s lignes au chemin %s : ", *nb_row, *)
	    source_tags := [3]string{"FAKE0","FAKE1","FAKE2"}
	    type_tags := [11]string{"Vitesse","Frequence cardiaque", "Puissance développé", "Cadence de pédalage", "Distance", "Temperature", "Dépense énergétique", "Pourcentage de pente", "Dénivelé positif", "Vitesse du vent", "Luminosité/Panneau solaire"}
	    unit_tags := [11]string{"km_h","bpm","Watts","tmp","km","°C","Kcal","%","D+", "m/s","lux"}
	    
	    db, err := leveldb.OpenFile(*path, nil)
	    if err != nil {
	    
	    	fmt.Println("Error creating LevelDB database: ", err.Error())
	    }
	    defer db.Close()
	    
	    nb_row_int, err := strconv.Atoi(*nb_row)
	    for i:=0;i<(nb_row_int);i++ {
	    	
	    	batch := new(leveldb.Batch)
		batch.Put([]byte(strconv.Itoa(i) + "Source"), []byte(source_tags[i%3]))
		batch.Put([]byte(strconv.Itoa(i) +"Type"), []byte(type_tags[i%11]))
		batch.Put([]byte(strconv.Itoa(i) + "Unit"), []byte(unit_tags[i%11]))
		batch.Put([]byte(strconv.Itoa(i) +"Value"), []byte(strconv.Itoa(rand.Intn(1000))))
		batch.Put([]byte(strconv.Itoa(i) +"Id"), []byte(strconv.Itoa(i)))

		err = db.Write(batch, nil)
		if err != nil {
			fmt.Println("Error: ", err.Error())
   		}
    }
}


