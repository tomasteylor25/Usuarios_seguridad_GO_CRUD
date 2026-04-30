# Usuarios_seguridad_GO_CRUD

Módulo de seguridad del proyecto **Xchango** una plataforma de trueque e intercambio de bienes y servicios. Este módulo se encarga de todo lo relacionado con los usuarios: registro, autenticación, manejo de contraseñas, sesiones activas y perfil personal.

---

## ¿Qué hace este módulo?

| Funcionalidad | Descripción |
|---|---|
| Usuarios | Registro y gestión de cuentas |
| Credenciales | Control de contraseñas y bloqueos por intentos fallidos |
| Sesiones | Manejo de tokens de sesión activos |
| Recuperación | Flujo de recuperación de contraseña por token o código |
| Historial | Registro de cambios de contraseña |
| Perfil | Datos personales del usuario |
| Intentos de login | Log de accesos exitosos y fallidos |

---

## Tecnologías utilizadas

- **Golang** — lenguaje principal del backend
- **Gorilla Mux** — manejo de rutas HTTP
- **PostgreSQL / PgAdmin 4** — base de datos relacional
- **Postman** — pruebas de endpoints

---

## Estructura del proyecto

```
Usuarios_seguridad_GO_CRUD/
├── main.go                  # Punto de entrada, configuración del servidor
├── go.mod                   # Dependencias del proyecto
├── config/
│   └── db.go                # Conexión a PostgreSQL
├── models/                  # Estructuras de datos (8 tablas)
├── controllers/             # Lógica de cada endpoint
│   └── helper.go            # Función compartida respondJSON
└── routes/                  # Registro de rutas por tabla
```

---

## Variables de entorno

```
HOST     → localhost
PORT     → 5432
USER     → postgres
PASSWORD → postgres
DB       → Xchango_db
SCHEMA   → Usuario_seguridad
PORT     → 8082
```

---

## Endpoints disponibles

Todos los endpoints siguen el patrón REST estándar:

| Método | Ruta | Acción |
|---|---|---|
| GET | `/usuarios` | Obtener todos |
| GET | `/usuarios/{id}` | Obtener uno por ID |
| POST | `/usuarios` | Crear nuevo |
| PUT | `/usuarios/{id}` | Actualizar |
| DELETE | `/usuarios/{id}` | Baja lógica |

---

## Base de datos

- **Motor:** PostgreSQL
- **Base de datos:** `Xchango_db`
- **Esquema:** `Usuario_seguridad`
- **Tablas:** Usuario, Credenciales, Perfil, Sesion, IntentoLogin, HistorialContrasena, RecuperacionContrasena, Crear_contrasena

---
