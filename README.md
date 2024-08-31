# API SQL avec Go et Gin

Une API RESTful écrite en Go utilisant le framework Gin et une base de données MySQL.

## Table des matières

- [Pré-requis](#pré-requis)
- [Installation](#installation)
- [Configuration](#configuration)
- [Lancement du projet](#lancement-du-projet)
- [Endpoints de l'API](#endpoints-de-lapi)
- [Structures de données](#structures-de-données)
- [Contribuer](#contribuer)
- [License](#license)

## Pré-requis

Avant de commencer, assurez-vous d'avoir installé les éléments suivants sur votre machine :

- [Go](https://golang.org/doc/install) (version 1.16 ou supérieure)
- [MySQL](https://www.mysql.com/downloads/)
- [Git](https://git-scm.com/)

## Installation

Clonez le repository et installez les dépendances nécessaires :

```bash
git clone https://github.com/votre-utilisateur/api-sql.git
cd api-sql
go mod tidy
```

## Configuration

Avant de lancer le projet, configurez les variables d'environnement nécessaires pour la connexion à la base de données MySQL. Vous pouvez définir ces variables dans votre shell ou utiliser un fichier `.env` :

```bash
export DBUSER="votre_utilisateur"
export DBPASS="votre_mot_de_passe"
export DBNAME="nom_de_votre_base"
```

**Note :** Assurez-vous que MySQL est en cours d'exécution et que les informations de connexion sont correctes.

## Lancement du projet

Pour démarrer le serveur, exécutez la commande suivante :

```bash
go run main.go
```

Le serveur sera accessible sur `http://localhost:8080`.

## Endpoints de l'API

Voici les différents endpoints disponibles pour l'API :

- `GET /albums` : Récupère la liste de tous les albums.
- `GET /albums/:id` : Récupère les détails d'un album spécifique par ID.
- `POST /albums` : Ajoute un nouvel album.
- `PUT /albums/:id` : Met à jour un album existant par ID.
- `DELETE /albums/:id` : Supprime un album par ID.

### Exemples de requêtes

#### GET /albums

```bash
curl http://localhost:8080/albums
```

#### POST /albums

```bash
curl -X POST http://localhost:8080/albums -d '{"Title": "Nouvel Album", "Artist": "Artiste Inconnu", "Price": 12.99}'
```

#### PUT /albums/:id

```bash
curl -X PUT http://localhost:8080/albums/1 -d '{"Title": "Album Mis à Jour", "Artist": "Artiste Inconnu", "Price": 15.99}'
```

#### DELETE /albums/:id

```bash
curl -X DELETE http://localhost:8080/albums/1
```

## Structures de données

Voici la structure des données pour un album :

```go
type Album struct {
    ID     int64   `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float32 `json:"price"`
}
```
