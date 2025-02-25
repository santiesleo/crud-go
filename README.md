# CRUD con Gin y PostgreSQL 🐹🐘

Este es un proyecto simple de CRUD (Crear, Leer, Actualizar, Eliminar) construido con el framework [Gin](https://gin-gonic.com/) para Go y una base de datos PostgreSQL. 🚀

## Requisitos 📋

- Go 1.21+
- PostgreSQL 14+
- [Git](https://git-scm.com/)

## Configuración del entorno 🛠️

1. Clona el repositorio:

```bash
git clone https://github.com/santiesleo/crud-go.git
cd crud-go
```

2. Crea la base de datos en PostgreSQL:

```sql
CREATE DATABASE gin_crud;
```

3. Crea el archivo `.env` en la raíz del proyecto:

```
DB_HOST=localhost
DB_USER=tu_usuario
DB_PASSWORD=tu_password
DB_NAME=gin_crud
DB_PORT=5432
SSL_MODE=disable
```

4. Instala las dependencias:

```bash
go mod init gin-crud-project
go get github.com/gin-gonic/gin github.com/lib/pq github.com/joho/godotenv
```

## Ejecución 🚀

```bash
go run main.go
```

La API estará disponible en `http://localhost:8080`.

## Endpoints 🌐

### Obtener todos los usuarios

```http
GET /users
```

### Crear un usuario

```http
POST /users
```

Body:

```json
{
  "name": "Santiago Escobar",
  "email": "santiesleo17@gmail.com"
}
```

### Actualizar parcialmente un usuario

```http
PATCH /users/{id}
```

Body (puede ser solo un campo):

```json
{
  "name": "Nuevo Nombre"
}
```

### Eliminar un usuario

```http
DELETE /users/{id}
```

## Estructura del proyecto 🗂️

```
.
├── db
│   └── connection.go
├── controllers
│   └── user_controller.go
├── models
│   └── user.go
├── routes
│   └── user_routes.go
├── main.go
├── .env
└── go.mod
```


---

✨ _Desarrollado con Go, Gin y PostgreSQL por Santiago Escobar_ ✨

