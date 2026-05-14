# Matrix Factorization API (Go)

API profesional desarrollada en **Go** utilizando el framework **Fiber v3** para realizar la factorización QR de matrices rectangulares. Incluye un sistema de autenticación robusto basado en JWT.

## Tecnologías
- [Go](https://go.dev/) (v1.26+)
- [Fiber v3](https://docs.gofiber.io/v3/)
- [GORM](https://gorm.io/) & [PostgreSQL](https://www.postgresql.org/)
- [JWT](https://jwt.io/) & [Bcrypt](https://en.wikipedia.org/wiki/Bcrypt)
- [Swaggo](https://github.com/swaggo/swag) & [Scalar](https://scalar.com/)

---

## Configuración (Variables de Entorno)
El proyecto utiliza un archivo `.env` para la configuración.

| Variable | Descripción | Por defecto |
| :--- | :--- | :--- |
| `PORT` | Puerto del servidor | `3001` |
| `DB_HOST` | Host de la base de datos | `localhost` |
| `DB_PORT` | Puerto de Postgres | `5432` |
| `DB_USER` | Usuario de Postgres | `postgres` |
| `DB_PASSWORD`| Contraseña de Postgres | `postgres` |
| `DB_NAME` | Nombre de la base de datos | `matrices_db` |
| `JWT_SECRET` | Clave secreta para firmar tokens | `tu_secreto` |

---

## Ejecución

### Infraestructura (Base de Datos)
Para levantar la base de datos PostgreSQL usando Docker:
```bash
docker compose up -d
```

### Instalar dependencias
```bash
go mod tidy
```

### Generar documentación (Swagger)
```bash
swag init -g cmd/api/main.go --parseInternal --parseDependency
```

### Iniciar el servidor
```bash
go run cmd/api/main.go
```

---

## Endpoints de Autenticación

### Registro de Usuario
**POST** `/auth/register`

**Payload (JSON):**
```json
{
  "email": "usuario@ejemplo.com",
  "password": "mi_password_seguro_123"
}
```

### Inicio de Sesión
**POST** `/auth/login`

**Payload de Entrada (JSON):**
```json
{
  "email": "usuario@ejemplo.com",
  "password": "mi_password_seguro_123"
}
```

**Respuesta Exitosa (JSON):**
```json
{
  "status": 200,
  "message": "Login exitoso",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "a48dff30-5510-45f1-86bd-cab52e0846c6",
      "email": "usuario@ejemplo.com",
      "created_at": "2024-05-13T21:21:28.11Z"
    }
  }
}
```

---

## Endpoints Protegidos

### Factorización QR
**POST** `/matrix/factorize`
- **Requisito**: Header `Authorization: Bearer <TOKEN_JWT>`

**Payload de Entrada (JSON):**
```json
{
  "matrix": [
    [12, -51, 4],
    [6, 167, -68],
    [-4, 24, -41]
  ]
}
```

**Respuesta Exitosa (JSON):**
```json
{
  "status": 200,
  "message": "Factorización QR completada",
  "data": {
    "q": [
      [0.85714, -0.39428, 0.33142],
      [0.42857, 0.90285, -0.03428],
      [-0.28571, 0.17142, 0.94285]
    ],
    "r": [
      [14, 21, -14],
      [0, 175, -70],
      [0, 0, -35]
    ]
  }
}
```

---

## Documentación
Accede a la consola de Scalar para probar todos los endpoints y la autenticación:
**[http://localhost:3001/docs](http://localhost:3001/docs)**

---

## Estructura del Proyecto
```text
.
├── cmd/api/            # Punto de entrada
├── internal/
│   ├── controllers/    # Manejadores HTTP
│   ├── services/       # Lógica de negocio (Algoritmo QR / Auth)
│   ├── repositories/   # Acceso a base de datos (GORM)
│   ├── models/         # Entidades de base de datos (User)
│   ├── routes/         # Definición de rutas y grupos
│   ├── dtos/           # Objetos de transferencia
│   └── middlewares/    # Auth JWT y Logger
├── docs/               # Documentación Swagger/Scalar
├── docker-compose.yml  # Infraestructura Postgres
└── README.md           # Este archivo
```
