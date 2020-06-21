# Installer l'environnement buildroot sur la raspberry PI 4


 (Possibilité d'utiliser le fichier DEFCONFIG propre à la RPI4 et le fichier .config).

* 1 Se positionner dans ~/SMILE_STAGE/SOURCES_GIT/BR2_BUILD_RPI$


* 2 git clone https://github.com/Costa83/br2_build_rpi.git


* 3 cd br2_build_rpi/br2_bd_guillaume/

* 4 br2_build_rpi/br2_bd_guillaume$ make O=${PWD}/output BR2_EXTERNAL=${PWD} -C /opt/buildroot br2_install_rpi4_defconfig


* 5 br2_build_rpi/br2_bd_guillaume/output$ make menuconfig


* 6 On sélectionne les options suivantes ou alors on charge un fichier .config permettant de charger toutes ses configurations.

On sélectionne différentes fonctionnalités nécessaires au projet :
- Target packags → Networking applications : cocher « dropbear »
On configure un mot de passe Root : 
- System configuration → Root password :
« gucos83»
On configure l’interface réseau pas défault :
- System configuration :
(eth0) Network interface to configure through DHCP.
On configure les paramètres propes à la rpi 
- Target packages → Harware handling → firmware 
« rpi-wifi-firmware ».
« rpi-bt-firmware ».
On active le support du C++ : 
- Toolchain
« Enable C++ support »
.config contient toute la configuration.
project_defconfig contient le minimum pour reconstruite .config
Permet de réutiliser une configuration avec une version postérieur de Buildroot

* 7 Sauvegarde du fichier .config résultant de la configuration de l'outil menuconfig
/br2_bd_guillaume# cp .config ../configs/.br2_intall_rpi4_config

* 8 Lancer la compilation : 

/SMILE_STAGE/SOURCES_GIT/BR2_BUILD_RPI/br2_build_rpi/br2_bd_guillaume/output$ make

* 9 Téléversement de l’image générée sur une carte uSD :

Insérer la carte uSD sur le PC.

S’assurer que la carte SD est bien démontée.

[repertoire projet]$  sudo umount /dev/sdb1
[repertoire projet]$  sudo umount /dev/sdb2

Ecrire l’image « sdcard.img » sur la carte uSD dev/mmcblk0 avec un bloc size de 1M

[repertoire projet]$ sudo dd if=output/images/sdcard.img of=/dev/sdb bs=1M

29 secondes environ sur ubunto 18.4.

On *ure en partant du cable (coté opposé au porrs USB) et on visualise le lien série avec l’outil minicom : 


Vérier que le système a bien détecté le port sur lequel on vient de se brancher avec dmesg | grep tty
On configure ensuite minicom avec la commande :

[repertoire projet]$  sudo su
[repertoire projet]$  minicom -s
On mets dans configuration du port série -> port série : /dev/ttyUSB0 puis entrée , puis configurer la vitesse de transmission à 115200 ainsi que veiller à ce que hardware flow control soit à 
enregistrer la configuration sous : "usb0" et entrer et sortir de minicom

[repertoire projet]$  sudo minicom -D /dev/ttyUSB0 usb0
 
 
 ## Connection de la raspberry PI 4 au réseau

Connecter un cable ethernet du PC host à la carte RPI4 (port ethernet).

Attribution d'une IP statique sur la RPI4 ainsi que sur le PC host.
Sur le PC host : 192.168.10 (0 pour O de host)
Dans parametre filaire sous ubuntu-> Fenetre IPV4 -> Décocher automatique et cocher manuel pour la méthode IPV4, puis attribuer une IP fixe : 
	Adresse 192.168.0.10   Maque de réseau 255.255.255.0  Passerelle : 192.168.0.1

Sur la RPI4 : 192.168.0.90
Editer le fichier /etc/network/interfaces

[repertoire projet]$  vi /etc/network/interfaces

iface eth0 inet static
  address 192.168.0.11
  netmask 255.255.255.0

Sauvegarder avec vi le fichier interfaces (: puis x! et entrée)

Activer l'interface eth0 avec la commande :

[repertoire projet]$  ifup eth0

Vérifiez que l'interface eth0 est bien montée: inet addr doit etre fixé à 192.168.0.11

[repertoire projet]$  ifconfig

# ifconfig                                                                      
eth0      Link encap:Ethernet  HWaddr DC:A6:32:04:A4:8A                         
          inet addr:192.168.0.11  Bcast:0.0.0.0  Mask:255.255.255.0             
          inet6 addr: fe80::dea6:32ff:fe04:a48a/64 Scope:Link 

Tester la connexion sur le sous réseau avec la commande ping :

Depuis le PC host : costa@costa-Latitude-5590:~$ ping 192.168.0.11

Depuis la RPI4 : # ping 192.168.0.10


## Connection en SSH à la carte RPI buildroot :

Depuis le PC host :

$ ssh root@192.168.0.11
tapez entrée puis le mot de passe root de la rpi4 : "gucos83"

Vous avez maintenant la main sur le terminal de la rpi4.



