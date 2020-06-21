#  Installation de  SQLITE sur Linux embarqué sur raspberry PI       

## Prérequis : avoir installé BR2 sur une RPI. (boot de Linux)

* Se placer dans le dossier output du projet buildroot: 

br2_bd_guillaume$ cd output
br2_bd_guillaume/output$ make menuconfig

* Taper "MAJ "+ "/", dans le menuconfig puis tapez sqlite.

* Se déplacer dans le menu aller dans : Target packages → libraries → Database

* Puis cocher "sqlite" ainsi que toutes les options sous sqlite (qui lui sont liées)

* Faire Save puis exit.

* Sauvegarder la configuration pour SQLITE :

$ cp .config ../configs/.sqlite_config

$ make

SQLITE lib et l'executable vont etre installés dans le répertoire host et target de output/ grace à l’outil libtool.

* Après compîlation on peut vérifier que libsqlite3.so libsqlite3.so.0.8.6 libsqlite3.so.0 sont présents dans br2_bd_guillaume/output/target/usr/lib/


libtool: install: /usr/bin/install -c sqlite3 /home/costa/SMILE_STAGE/SOURCES_GIT/BR2_BUILD_RPI/br2_build_rpi/br2_bd_guillaume/output/host/arm-buildroot-linux-uclibcgnueabihf/sysroot/usr/bin/sqlite3



##  Réécriture de la carte SD 

* Insérer la carte uSD dans le PC
$lsblk

* Identifier le nom de la carte uSD

$ umount /dev/sdb1
$ umount dev/sdb2
$ sudo dd if=images/sdcard.img of=/dev/sdb bs=1M

* Retirer la carte uSD du PC et l’insérer la carte dans la RPI et la mettre sous tension.

## Utiliser SQLITE3 :

* S’identifier sur la session : user : root et mdp : gucos83

$ sqlite3

* Le prompt de sqlite3 devrait s’afficher sur le terminal de la raspberry pi 4.




##  Copie des excutables de tests développés pour raspbian sur br2

* Insertion d’une clé USB sur ubuntu, et formatage de celle ci avec l’outil disques. (Format FAT32).


* Modification de la taille du rootfilesystem depuis menuconfig :

* Depuis le PC Host ubuntu : 

$ cd output
$ make menuconfig

* Allez dans filesystem image puis remplir le champ «exact size», rentrez 500M

$ make


* Copie sur clé usb de br2_bd_guillaume/package/sqlite/

 ssh-keygen -f "/home/costa/.ssh/known_hosts" -R "192.168.0.11"
 br2_build_rpi/br2_bd_guillaume/$ cd package
 br2_build_rpi/br2_bd_guillaume/package$  scp sqlite/* root@192.168.0.11:/bin/

Tapez « yes »


* Pour voir le nom du block device :

$ fsdisk -l

sda1 dans mon cas

* Creer un point de montage avant montage de la clé USB :

$ mkdir /media/usb_test_sqlite/ 
# mount /dev/sda1 /media/usb_test_sqlite/ 

## Copie des excutables sur la cible:

br2_bd_guillaume/package $ scp sqlite/* root@192.168.0.11:/usr/bin
# chmod +x /usr/bin/sttTest


* Pour installer les executables il va falloir les compiler avec les librairies qui ont été généré grace à l’activation de sqlite3 avec menuconfig Pour cela on copie depuis le dossier /output/host/bin/arm-buildroot-linux-uclibcgnueabihf-gcc les .so dans br2_bd_guillaume/package/src/


* Copie de libsqlite3.so.0.8.6 , libsqlite3.so.0 ,  libsqlite3.so

* On modifie aussi le makefile pour appeler la librairie et le cross compilateur de br2.


* Modification du makefile :

CC = ../../../output/host/bin/arm-buildroot-linux-uclibcgnueabihf-gcc
CFLAGS  = 
RM = rm
lIBPATH =
LIB = -lsqlite3

	

* -On lance ensuite la compilation du makefile :

~/SMILE_STAGE/SOURCES_GIT/BR2_BUILD_RPI/br2_build_rpi/br2_bd_guillaume/package/sqlite/src$ make all


scp sqlite/script_test/* root@192.168.0.11:/usr/test_sqlite


rpi# cd usr/test_sqlite/
rpi# chmod +x *
rpi# ./sttTestInsertRowsSqlite -i10
rpi# sqlite3 test_sqlite.db












