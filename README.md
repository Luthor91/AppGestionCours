# AppGestionCours

## Description

Ce projet est une plateforme éducative interactive utilisant Redis comme système de stockage de données.


## Spécifications de l’application

1. La plateforme éducative doit proposer des cours, des professeurs et des étudiants. Chaque
cours a un ID, un titre, un enseignant et une liste d’étudiants inscrits. Les cours peuvent également avoir d’autres propriétés, telles que le résumé du cours, le niveau du cours
(débutant, intermédiaire, avancé), et le nombre de places disponibles.

2. Les professeurs et les étudiants ont des profils contenant des informations telles que le nom,
l’ID, les cours auxquels ils sont inscrits (pour les étudiants) ou les cours qu’ils enseignent
(pour les professeurs). Les profils doivent également inclure des fonctionnalités pour mettre
à jour ces informations.

3. Implémentez un système de nouvelles de type publish-subscribe qui :
    - Du côté de l’éditeur, permet à l’enseignant de publier des mises à jour de cours, de créer
de nouveaux cours et d’émettre un message de nouvelles contenant l’ID du cours mis à
jour ou nouvellement créé.
    - Du côté de l’abonné, permet aux étudiants de s’abonner aux mises à jour du cours,
de récupérer les détails du cours à partir d’une nouvelle par l’ID et d’afficher l’entrée
complète du cours à partir de la base de données. Les étudiants doivent également
pouvoir s’inscrire à des cours via la plateforme.
    - Faites expirer les cours après un certain temps (si le cours n’est pas mis à jour ou si
personne ne s’y inscrit) : ces cours ne sont plus disponibles pour l’inscription.
    - Si un étudiant s’inscrit à un cours (par exemple, en définissant un certain champ dans la base de données), fait rafraîchir la date d’expiration du cours.

4. La plateforme doit également inclure une fonction de recherche qui permet aux utilisateurs
de chercher des cours par titre, enseignant, niveau, ou d’autres critères pertinents.

## Pré-requis

Vous devez avoir ces logiciels sur votre machine : 

### Golang
  - Installable depuis ce lien : https://go.dev/doc/install

### Postgresql
  - Installable depuis ce lien : https://www.enterprisedb.com/downloads/postgres-postgresql-downloads

### Redis
  - Installable sur machine Linux uniquement avec ce script 
```
  curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

  echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

  sudo apt-get update
  sudo apt-get install redis
```

Si vous êtes sur Windows, faites attention à exécuter cette ligne de commande avant
```
wsl --install -d Ubuntu
```

Pour finir, sur la machine devant héberger Redis, veuillez utiliser cette commande :
```
redis-server
```

### Installation de Make

Installer Make n'est pas obligatoire, mais fortement conseillé


Sur une machine Ubuntu vous allez devoir l'installer avec cette commande :

```
sudo apt install make
```

Sur une machine Windows vous allez devoir installer chocolatey avec cette commande sur un terminal Powershell :

```
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

Ensuite vous pourrez installer make avec cette commande dans un terminal ayant des permissions d'administrateur : 

```
choco install make
```

## Exécution du projet

### Avec make

Ouvrez un terminal situé dans le projet, exécutez cette commande :
```
make all
```

Le projet devrait être lancé sur l'url http://127.0.0.1:8080/index

### Sans make

Pour les 3 commandes ci-dessous, nous assumons que vous êtes à la racine du projet, dans le répertoire "TP5_Redis"

Tout d'abord nous allons exécuter la commande pour charger les dépendences du projet
```
cd main_module && go mod tidy
```
Ensuite nous allons compiler le projet
```
cd main && go build
```
Pour finir, nous allons exécuter le projet
```
go run main.go
```

Vérifiez bien que vous êtes dans le dossier main_module/main avant d'effectuer la compilation.


L'application devrait être lancé sur l'url http://127.0.0.1:8080/index.

## A faire

### FRONTEND

 - Une page de login
    - La connexion se fera avec un email et un mot de passe

 - Une page de création de compte
    - La création de compte se fera avec un nom, un email et deux champs mots de passe qui devront correspondre

 - Une page permettant la création + modification d'un cours
    - Une fois validé, une notification est envoyé à tout le monde, elle s'affichera dans un coin de l'écran pendant un laps de temps fixe
    - Les champs utilisés seront : 
        - Identifiant (peut-être générable aléatoirement ?)
        - Nom du cours
        - Résumé du cours
        - Niveau du cours, débutant, intermédiaire, avancé sont les seules valeurs autorisés
        - Place disponible, ça sera le nombre de place maximum
        - Temps d'expiration, c'est le temps d'abonnement maximum pour un étudiant, pas le temps d'expiration du cours sur la plateforme

 - Une page de détail du cours avec toutes les informations du cours et la possibilité de s'y inscrire
    - Les informations suivantes seront affichés : 
        - Identifiant
        - Nom du cours
        - Résumé du cours
        - Niveau du cours
        - Place disponible
        - Temps d'expiration

 - Une page avec la possibilité de chercher un cours en fonction du nom du cours, de l'enseignant, du niveau et d'autres critères.

### BACKEND

 - L'objet Utilisateur devras être stocké dans la session (ou en cache)

 - Faire en sortes que les données récupérés par la base de données soient mis en cache avec Redis

 - Priorité à la récupération de données sur Redis, ensuite sur Postgres

 - Les données de l'utilisateurs doivent être stockés en session et/ou cache pour vérifier sur chaque page si il est Etudiant ou Professeur

 - Au moment de l'inscription à un cours, on mettra un timer qui défini la durée d'abonnement au cours
    - Utiliser EXPIRE de Redis
    - Utiliser PUBLISH et SUBSCRIBE de Redis pour l'abonnement
    - Donc pour cette données on ira pas chercher dans Postgresql
 
 - L'inscription échoue si :
    - Le nombre de place disponible est égal à 0

 - Envoyer un mail sur l'adresse utilisé pour l'inscription


## Structure du projet

### Frontend

Le frontend se situe à la racine du projet dans le répertoire "public".

### Backend

Le backend se situe à la racine du projet dans le répertoire "main_module".
Dans le répertoire "main_module" nous avons 5 autres répertoires et deux autres fichiers servant à la compilation du projet.

- controller : Contient les fichiers de contrôleur, la logique métier, permettent d'interroger la base de données.

- database : Contient les ficheirs permettant de faire la création, configuration et connexion aux bases de données.

- main : Contient le fichier principal de l'application, le fichier "main.go".

- models : Contient les fichiers de modèles entité, des fonctions de formattage de données.

- server_config : Contient le fichier "routes.go" permettant de définir les routes de l'application.


## Tâches à réaliser

### Frontend

- Pages de l'applications
  - login.html
    - page de connexion
    - demandera email et mot de passe
  - creation_compte.html
    - page de création de compte
    - demandera email, 2 fois le mot de passe
  - myaccount.html
    - page affichant les information de l'utilisateur


### Backend

- utiliser Publish/Subscribe

## Contact

Si l'installation ne fonctionne pas, ou si vous constatez des problèmes, contactez luthorino sur Discord
