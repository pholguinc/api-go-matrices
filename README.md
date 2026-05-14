# Matrix Factorization API (Go)

API profesional desarrollada en **Go** utilizando el framework **Fiber v3** para realizar la factorización QR de matrices rectangulares.

## Tecnologías
- [Go](https://go.dev/) (v1.26+)
- [Fiber v3](https://docs.gofiber.io/v3/)
- [Swaggo](https://github.com/swaggo/swag)
- [Scalar](https://scalar.com/)

---

## Configuración (Variables de Entorno)
El proyecto utiliza un archivo `.env` para la configuración. Puedes basarte en `.env.template`.

| Variable | Descripción | Por defecto |
| :--- | :--- | :--- |
| `PORT` | Puerto en el que correrá el servidor | `3001` |

---

## Ejecución

### 1. Instalar dependencias
```bash
go mod tidy
```

### 2. Generar documentación (Swagger)
Si realizas cambios en las anotaciones de la API, ejecuta:
```bash
swag init -g cmd/api/main.go
```

### 3. Iniciar el servidor
```bash
go run cmd/api/main.go
```

---

## Documentación de la API
Una vez iniciado el servidor, puedes acceder a la documentación interactiva en:
**[http://localhost:3001/docs](http://localhost:3001/docs)**

---

## Endpoints

### Factorización QR
Realiza la descomposición de una matriz $A$ en una matriz ortogonal $Q$ y una matriz triangular superior $R$.

**URL:** `/matrix/factorize`  
**Método:** `POST`

#### Ejemplo de Payload (Request):
```json
{
  "matrix": [
    [12, -51],
    [6, 167]
  ]
}
```

#### Ejemplo de Respuesta (Success):
```json
{
  "status": 200,
  "message": "Factorización QR completada",
  "data": {
    "q": [
      [0.8944271909999159, -0.4472135954999579],
      [0.4472135954999579, 0.8944271909999159]
    ],
    "r": [
      [13.416407864998739, 29.06915570749726],
      [0, 172.1772342674838]
    ]
  }
}
```

---

## Estructura del Proyecto
```text
.
├── cmd/
│   └── api/            # Punto de entrada de la aplicación
├── internal/
│   ├── controllers/    # Manejadores de peticiones HTTP
│   ├── services/       # Lógica de negocio (Algoritmo QR)
│   ├── routes/         # Definición de rutas
│   ├── dtos/           # Objetos de transferencia de datos
│   ├── constants/      # Constantes y mensajes de error
│   └── middlewares/    # Middlewares (Logger, etc.)
├── docs/               # Documentación generada y archivos estáticos
├── .env                # Configuración local
└── go.mod              # Dependencias del proyecto
```
