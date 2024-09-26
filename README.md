After Death
Jeu créer un golang

Comment installer Visual Studio Code (éditeur de code)

Lien vers le site : https://code.visualstudio.com/download

Linux : 

Effectuez cette commande dans un terminal après avoir téléchargé VS Code

sudo apt install ./<file>.deb


Windows :

Téléchargez, exécutez le programme d'installation et suivez les instructions


Une fois dans VS Code, vous devez installer l'extension Golang pour utiliser le jeu
-Cliquez sur le panneau "Extensions" à gauche (Ctrl+Shift+X)
-Rechercher "Go" et installez l'extension




Comment installer Golang sur votre ordinateur (Language utiliser pour le jeu)

Lien vers le site : https://go.dev/dl/

Linux :

Supprimez toute installation Go précédente en supprimant le dossier /usr/local/go (s'il existe), puis extrayez l'archive que vous venez de télécharger dans /usr/local, en créant une nouvelle arborescence Go dans /usr/local/go :

$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz

(Vous devrez peut-être exécuter la commande en tant que root ou via sudo).

Ne décompressez pas l'archive dans une arborescence /usr/local/go existante. Cela est connu pour produire des installations Go défectueuses.
Ajoutez /usr/local/go/bin à la variable d'environnement PATH.

Vous pouvez le faire en ajoutant la ligne suivante à votre $HOME/.profile ou /etc/profile (pour une installation à l'échelle du système) :

export PATH=$PATH:/usr/local/go/bin

Remarque : les modifications apportées à un fichier de profil peuvent ne pas s'appliquer avant la prochaine fois que vous vous connectez à votre ordinateur. Pour appliquer les modifications immédiatement, exécutez simplement les commandes shell directement ou exécutez-les à partir du profil à l'aide d'une commande telle que source $HOME/.profile.
Vérifiez que vous avez installé Go en ouvrant une invite de commande et en tapant la commande suivante :

$ go version

Confirmez que la commande imprime la version installée de Go.



Windows :

Ouvrez le fichier MSI que vous avez téléchargé et suivez les instructions pour installer Go.

Par défaut, le programme d'installation installe Go dans Program Files ou Program Files (x86). Vous pouvez modifier l'emplacement selon vos besoins. Après l'installation, vous devrez fermer et rouvrir toutes les invites de commande ouvertes afin que les modifications apportées à l'environnement par le programme d'installation soient reflétées dans l'invite de commande.
Vérifiez que vous avez installé Go.

Sous Windows, cliquez sur le menu Démarrer.
Dans la zone de recherche du menu, saisissez cmd, puis appuyez sur la touche Entrée.
Dans la fenêtre d'invite de commande qui s'affiche, saisissez la commande suivante :

$ go version

Confirmez que la commande imprime la version installée de Go.




Comment installer Raylib (bibliothèque requise pour le jeu)


Linux : il faut effectuer les 2 commandes suivantes dans un terminal sur Visual Studio Code (explication pour le terminal en dessous)

sudo apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libwayland-dev libxkbcommon-dev 
go get -v -u github.com/gen2brain/raylib-go/raylib



Windows :
Il est déjà installer grâce au raylib.dll dans le projet


Comment installer Git (Logiciel qui servira a clôner le jeu)

Linux : il faut effectuer la commande suivantes dans un terminal sur Visual Studio Code (explication pour le terminal en dessous)

Effectuez cette commande dans un terminal après avoir téléchargé VS Code :
sudo apt install git-all



Windows : 

Téléchargez, exécutez le programme d'installation et suivez les instructions via ce site :
https://git-scm.com/download/win



Il faut maintenant clôner le projet
-Cliquez sur "Terminal", puis "New Terminal"
-Entrer cette commande dans le terminal :
git clone https://github.com/IliesBossuyt/arcadiaB1.git
-Maintenant que "ArcadiaB1" à été clôner sur votre ordinateur, ouvrez le dossier sur VS Code dans "File", puis "Open Folder"
-Vous pouvez maintenant lancer le jeu avec la commande suivante dans le terminal :
go run .


Membre du projet : 
Maxime : https://github.com/maxime150306
Léo : https://github.com/Leocollowald
Thomas : https://github.com/thom972
Ilies : https://github.com/IliesBossuyt