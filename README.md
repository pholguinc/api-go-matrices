# Matrix Factorization API (Go)

API profesional desarrollada en **Go** utilizando el framework **Fiber v3** para realizar la factorización QR de matrices rectangulares. Este proyecto está diseñado para ser **Cloud-Ready**, soportando despliegues en contenedores (Docker/ECS) y arquitecturas Serverless (AWS Lambda).

## Tecnologías
- [Go](https://go.dev/) (v1.24+)
- [Fiber v3](https://docs.gofiber.io/v3/)
- [GORM](https://gorm.io/) & [PostgreSQL](https://www.postgresql.org/)
- [JWT](https://jwt.io/) & [Bcrypt](https://en.wikipedia.org/wiki/Bcrypt)
- [Docker](https://www.docker.com/) (Multi-stage)
- [Serverless Framework](https://www.serverless.com/) (AWS Lambda)

---

## Ejecución Local (Docker Compose)

La forma más rápida de iniciar todo el ecosistema (API + DB) es usando Docker Compose:

```bash
docker compose up --build
```

Esto levantará la API en el puerto `3001` y una base de datos PostgreSQL lista para usar. El contenedor de la API está configurado en modo **development**, lo que facilita la depuración.

---

## Configuración (Variables de Env)

| Variable | Descripción | Producción (AWS) |
| :--- | :--- | :--- |
| `PORT` | Puerto del servidor | `3001` |
| `DB_HOST` | Host de la base de datos | RDS Endpoint / Supabase |
| `DB_PORT` | Puerto de Postgres | `5432` |
| `DB_USER` | Usuario de Postgres | `user_admin` |
| `DB_NAME` | Nombre de la base de datos | `matrices_db` |
| `DB_SSLMODE` | Modo SSL de conexión | `require` |
| `JWT_SECRET` | Clave para firmar tokens | Secreto largo y complejo |

---

## Despliegue en Producción

### 1. AWS Lambda (Serverless)
El proyecto incluye un adaptador para correr como una función Lambda detrás de un API Gateway.

**Compilar para AWS (ARM64):**
```bash
GOOS=linux GOARCH=arm64 go build -o bootstrap cmd/api/main.go
```

**Desplegar:**
```bash
serverless deploy
```

### 2. Docker (Producción)
Para generar una imagen optimizada para producción (mínimo peso, máxima seguridad):

```bash
docker build --target production -t matrix-api-go:prod .
```

---

## Endpoints

### Autenticación
- **POST** `/auth/register`: Registro de usuarios (UUID).
- **POST** `/auth/login`: Obtención de Token JWT.

### Operaciones (Protegidas)
- **POST** `/matrix/factorize`: Calcula QR y guarda el historial en DB.
- **GET** `/matrix/history`: Recupera cálculos previos del usuario.

---

## Documentación Interactiva
Accede a la consola de Scalar para probar todos los endpoints y la autenticación:
**[http://localhost:3001/docs](http://localhost:3001/docs)**

---

## Estructura del Proyecto
```text
.
├── cmd/api/            # Punto de entrada (Local & Lambda)
├── internal/
│   ├── controllers/    # Manejadores HTTP
│   ├── services/       # Lógica (Algoritmo QR / Auth / Persistencia)
│   ├── repositories/   # Acceso a DB (GORM)
│   ├── models/         # Entidades (User, MatrixRecord)
│   └── middlewares/    # Auth JWT, Logger y Adaptadores
├── Dockerfile          # Multi-stage (Dev & Prod)
├── docker-compose.yml  # Orquestación Local
├── serverless.yml      # Configuración AWS Lambda
└── README.md           # Este archivo
```
