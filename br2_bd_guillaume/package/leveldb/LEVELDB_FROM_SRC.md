# INSTALLER LEVELDB SUR LE PC HOTE, DEPUIS ZERO :


* Outils : LevelDB



## ETAPE 2: Installer LevelDB: sur Ubuntu 

* Source : https://github.com/syndtr/goleveldb
* Source : https://godoc.org/github.com/syndtr/goleveldb/leveldb
* Source : https://pypi.org/project/leveldb-cli/
* Source : https://computingforgeeks.com/how-to-install-leveldb-on-ubuntu-18-04-ubuntu-16-04/


$ echo "deb https://repos.influxdata.com/ubuntu bionic stable" | sudo tee /etc/apt/sources.list.d/influxdb.list


Update apt index and install influxdb 

$ sudo apt-get install -y influxdb

$ sudo apt-get install influxdb-client

Start and enable the service to start on boot up : 

$ sudo systemctl enable --now influxdb

$ sudo systemctl is-enabled influxdb

Check service status:

$ systemctl status influxdb

Open a terminal and check if influxdb is installed :

$ influxd




# INSTALLER GOLANG et les tests de LevelDB en GO sur  UBUNTU HOST et Cible ARM :

- On est sur une architecture client-serveur où ici le client et le serveur est situé sur le PC hote.


## Prérequis 2: Installer la commande go pourconstruire le package depuis l’hote:

$ sudo apt install golang-go

$ git clone https://go.googlesource.com/go goroot

$ cd goroot/

$ git checkout go1.14.2

$ git status

$ cd src/

$ ./all.bash 
 
$  go version


$ mkdir test_leveldb

$ git clone .....

$ cp test_leveldb/* test_leveldb   (leveldb/ measuretime/ go.mod go.sum main.go test_leveldb LEVELDB_FROM_SRC.md)

$ cd test_leveldb

$ rm -rf go.* 
$ rm -rf test_leveldb

$ go mod init test_leveldb
$ go mod tidy

(génère un go.mod et un go.sum)

$ env GOOS=linux GOARCH=arm GOARM=7 go build                         (ou go build sur pc hote)

Génère un exécutable arm

$ ./test_leveldb

On retrouve la base de donnée dans testleveldb/ (pc hote)

$ scp test_leveldb root@192.168.0.11:/usr/


On lance l’exécutable avec ./test_leveldb et les options du programme :

./main --Points=10 --Ip=localhost –Port=8086

########## Test LEVELDB DB in Go  ########################
--- Enter the IpAdress of rows to insert : --Ip=...
--- Enter the PortAddress of rows to insert : --Port=...
--- Enter the number of rows to insert : --Points=...
--- Enter the Middle of rows to insert : --Middle=...
--- Enter the Last of rows to insert : --Last=...
##########################################################






