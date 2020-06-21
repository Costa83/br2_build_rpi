# INSTALLER LEVELDB SUR LE PC HOTE, DEPUIS ZERO :

* Outils : LevelDB

## ETAPE 2: Installer LevelDB: sur Ubuntu 

* Source : https://projetsdiy.fr/influxdb-installation-ubuntu-16-04/
* Source : https://pimylifeup.com/raspberry-pi-influxdb/


$ echo "deb https://repos.influxdata.com/ubuntu bionic stable" | sudo tee /etc/apt/sources.list.d/influxdb.list


Mettre à jour l'index de apt et installer le paquet influxdb : 

$ sudo apt-get install -y influxdb

$ sudo apt-get install influxdb-client

Démarrer et permettre le démarage du service au démarrage : 

$ sudo systemctl enable --now influxdb

$ sudo systemctl is-enabled influxdb

Vérifier le statut du service : 

$ systemctl status influxdb

Ouvrir un terminal et vérifier si influxdb est bien installé:

$ influxd


# INSTALLER GOLANG et les tests de LevelDB en GO sur  UBUNTU HOST et Cible ARM :

- On est sur une architecture client-serveur où ici le client et le serveur est situé sur le PC hote.


## Prérequis 2: Installer la commande go pourconstruire le package depuis l’hote:

* Source : https://github.com/golang/go/wiki/Ubuntu
* Source : https://golang.org/doc/install/source

$ sudo apt install golang-go

$ git clone https://go.googlesource.com/go goroot

$ cd goroot/

$ git checkout go1.14.2

$ git status

$ cd src/

$ ./all.bash 
 
$ go version


$ mkdir test_influxdb

$ git clone .....

$ cp test_influxdb/* test_leveldb   (leveldb/ measuretime/ go.mod go.sum main.go test_influxdb INFLUXDB_FROM_SRC.md)

$ cd test_influxdb

$ rm -rf go.* 
$ rm -rf test_influxdb

$ go mod init test_influxdb
$ go mod tidy

(qui génère un go.mod et un go.sum permettant de construire l'executable)

$ env GOOS=linux GOARCH=arm GOARM=7 go build                         (ou go build sur pc hote)

Génère un exécutable arm qui peut être lancé avec :

$ ./test_influxdb

On retrouve la base de donnée dans testinfluxdb/ (sur pc hote)

$ scp test_influxdb root@192.168.0.11:/usr/

On lance l’exécutable avec ./test_influxdb et les options du programme.





