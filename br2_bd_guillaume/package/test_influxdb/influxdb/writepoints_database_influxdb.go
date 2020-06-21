package influxdb

import (
    "github.com/influxdata/influxdb1-client/v2"
    "fmt"
    "strconv"
    "time"
    "math/rand"
)

/*Ecriture en base de donnée de n lignes "nb_row"  en base de donnée sur une IP "IpAdress"  et port  "PortAddress" précisées dans le main*/

func Write_database_influxdb(nb_row *string, IpAdress *string, PortAddress *string){

    source_tags := [3]string{"FAKE0","FAKE1","FAKE2"}
    type_tags := [11]string{"Vitesse","Frequence cardiaque", "Puissance développé", "Cadence de pédalage", "Distance", "Temperature", "Dépense énergétique", "Pourcentage de pente", "Dénivelé positif", "Vitesse du vent", "Luminosité/Panneau solaire"}
    unit_tags := [11]string{"km_h","bpm","Watts","tmp","km","°C","Kcal","%","D+", "m/s","lux"}

    MyDB :=  "test_influxdb_go"
    Addr := "http://" + *IpAdress + ":" + *PortAddress

    c, err := client.NewHTTPClient(client.HTTPConfig{
    Addr: Addr,

  })

  if err != nil {
    fmt.Println("Error creating InfluxDB Client: ", err.Error())
  }
  defer c.Close()

  nb_row_int, err := strconv.Atoi(*nb_row)
  for i:=0;i<(nb_row_int);i++ {
      bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
  		Database: MyDB,
  	 })

     tags := map[string]string{"source": source_tags[i%3],"type":type_tags[i%11], "unit":unit_tags[i%11], "idtag":strconv.Itoa(i)}
	   fields := map[string]interface{}{
  	 	"value": rand.Intn(1000),
      "id": i,
	   }

     now :=  time.Now()
     measurement := "stt"
     pt, err := client.NewPoint(measurement, tags, fields, now)
     if err != nil {
         	fmt.Println("Error: ", err.Error())
     }
     bp.AddPoint(pt)
     c.Write(bp)
  }

}
