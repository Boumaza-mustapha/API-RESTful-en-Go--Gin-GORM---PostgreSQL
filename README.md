# 📦 GO_CRUD_API – API RESTful en Go

Ce projet est une API simple de type **CRUD (Create, Read, Update, Delete)** développée en **Go**, utilisant :

- Le framework **Gin** pour le routage HTTP
- **GORM** pour l'accès à la base de données PostgreSQL
- Le chargement d’environnement avec `gotenv`
- Une architecture modulaire et claire

---

## 🎯 Objectif

Ce projet a été développé dans le cadre de l’apprentissage des bases du développement d’API REST avec Go, en mettant l’accent sur :

- La structuration propre d’un projet Go
- L’interaction avec une base de données via **GORM**
- L’utilisation de **Supabase** comme backend PostgreSQL

---

## 📁 Structure du projet

```

GO\_CRUD\_API/
├── controllers/     # Fonctions CRUD (Create, Read, Update, Delete)
├── initializers/    # Connexion à la base de données + chargement des variables d’environnement
├── models/          # Définition des modèles (Post)
├── migrate.go       # Script de migration de la base de données
├── main.go          # Point d’entrée principal de l’application
├── .env             # Fichier de configuration (non versionné)

````

---

## ⚙️ Installation et exécution

### 1. Cloner le projet

```bash
git clone https://github.com/TON-UTILISATEUR/GO_CRUD_API.git
cd GO_CRUD_API
````

### 2. Configurer la base de données

Crée un fichier `.env` à la racine du projet contenant :

```env
DB_URL=postgres://user:password@host:port/dbname
```

> Remplace les valeurs par celles de ta base de données (ex. Supabase ou local).

### 3. Installer les dépendances

```bash
go mod tidy
```

### 4. Migrer la base de données

```bash
go run migrate.go
```

### 5. Lancer l'application

```bash
go run main.go
```

L’API sera disponible sur :
📍 `http://localhost:8080`

---

## ✅ Endpoints disponibles

| Méthode | Route        | Description              |
| ------- | ------------ | ------------------------ |
| POST    | `/posts`     | Créer un nouveau post    |
| GET     | `/posts`     | Récupérer tous les posts |
| GET     | `/posts/:id` | Récupérer un post par ID |
| PUT     | `/posts/:id` | Mettre à jour un post    |
| DELETE  | `/posts/:id` | Supprimer un post        |

---

## 🛠️ Technologies utilisées

* [Go](https://golang.org/)
* [Gin](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io/)
* [gotenv](https://github.com/subosito/gotenv)
* [Supabase](https://supabase.com/) (ou PostgreSQL local)

