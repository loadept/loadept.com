# Loadept Core

> 📝 **Nota**: Este es un proyecto de ejemplo/plantilla basado en una aplicación real privada. Se ha creado como referencia arquitectónica, removiendo características privadas y datos sensibles del proyecto original.

API REST backend desarrollada en Go. Proporciona servicios de gestión de artículos, categorías, compresión de PDFs y generación de sitemaps.

## 🚀 Características

- **API RESTful** con arquitectura limpia (Hexagonal Architecture)
- **Gestión de artículos** organizados por categorías
- **Caché distribuido** con Redis para alto rendimiento
- **Compresión de PDFs** mediante proxy reverso
- **Generación automática de sitemap.xml**
- **Health checks** para monitoreo
- **Middleware de logging** y compresión Brotli
- **Base de datos SQLite** con migraciones
- **Dockerizado** con multi-stage builds
- **Tests unitarios** con coverage

## 📋 Requisitos previos

- Go 1.24.0 o superior
- Redis 8.0 o superior
- Docker y Docker Compose (opcional)
- SQLite3

## 🛠️ Instalación

### Configuración local

1. Clonar el repositorio:
```bash
git clone https://github.com/loadept-web/core.git
cd core
```

2. Copiar el archivo de configuración:
```bash
cp .env.example .env
```

3. Configurar las variables de entorno en `.env`:
```bash
DEBUG=true
SECRET_KEY=base64secretkey==
PORT=8080
DB_NAME=db.sqlite3

REDIS_HOST=localhost
REDIS_PORT=6380
REDIS_USER=default
REDIS_PASSWORD=

GITHUB_API=https://api.github.com
GITHUB_TOKEN=your_github_token

PDF_SERVICE_URL=http://localhost:8000
```

4. Instalar dependencias:
```bash
go mod download
```

5. Ejecutar migraciones:
```bash
go run cmd/db_migrate/main.go
```

6. Iniciar el servidor:
```bash
go run cmd/loadept/main.go
```

El servidor estará disponible en `http://localhost:8080`

### Configuración con Docker

1. Configurar variables de entorno en `.env`

2. Construir y ejecutar los contenedores:
```bash
docker-compose up -d
```

Servicios disponibles:
- **web**: API backend (puerto interno)
- **cache**: Redis (puerto 6380)
- **proxy**: Nginx reverse proxy (puertos 80/443)

## 📁 Estructura del proyecto

```
.
├── api/                      # Capa de presentación
│   ├── middleware/          # Middlewares (CORS, logging, encoding)
│   └── v1/                  # Endpoints API v1
│       ├── router.go
│       └── handler/         # Handlers HTTP
├── cmd/                     # Puntos de entrada de aplicaciones
│   ├── loadept/            # Servidor principal
│   └── db_migrate/         # Herramienta de migraciones
├── internal/               # Código interno de la aplicación
│   ├── application/        # Casos de uso / Servicios
│   ├── config/            # Configuración
│   ├── di/                # Inyección de dependencias
│   ├── domain/            # Entidades y repositorios
│   ├── infrastructure/    # Implementaciones (DB, Cache)
│   └── validation/        # Validadores
├── pkg/                   # Paquetes reutilizables
│   ├── logger/           # Sistema de logging
│   └── respond/          # Utilidades de respuesta HTTP
├── migrations/           # Migraciones de base de datos
├── scripts/             # Scripts de utilidad
├── test/               # Tests de integración
├── docker-compose.yml  # Configuración de Docker Compose
├── Dockerfile         # Multi-stage build
└── go.mod            # Dependencias Go
```

## 🔌 API Endpoints

### Health Check
```
GET /api/v1/health
```
Verifica el estado del servidor y sus dependencias.

### Categorías
```
GET /api/v1/categories
```
Obtiene la lista de todas las categorías disponibles.

### Artículos
```
GET /api/v1/articles/{category}
```
Lista todos los artículos de una categoría específica.

```
GET /api/v1/articles/{category}/{name}
```
Obtiene el contenido completo de un artículo específico.

### PDF
```
POST /api/v1/pdf/compress
```
Comprime un archivo PDF mediante proxy reverso al servicio de PDF.

### Sitemap
```
GET /sitemap.xml
```
Genera el sitemap XML del sitio web.

## 🧪 Testing

Ejecutar todos los tests:
```bash
go test ./...
```

Ejecutar tests con coverage:
```bash
go test -cover ./...
```

Ejecutar tests específicos:
```bash
go test ./api/middleware/...
go test ./pkg/logger/...
```

## 🏗️ Build

### Build local
```bash
CGO_ENABLED=1 go build -o loadept.com cmd/loadept/main.go
```

### Build con script
```bash
chmod +x scripts/build.sh
./scripts/build.sh
```

Esto genera un archivo `loadept-core.tar.gz` listo para deployment.

### Build con Docker
```bash
docker build -t loadept-core:latest .
```

## 🔧 Arquitectura

El proyecto implementa **Arquitectura Hexagonal** (Ports & Adapters):

### Conceptos Fundamentales

**Hexágono (Núcleo de la Aplicación)**
- `internal/domain/`: Entidades de negocio
- `internal/application/`: Lógica de negocio y casos de uso
- Completamente independiente de tecnologías externas

**Puertos (Interfaces)**
- Definen los contratos de comunicación
- `internal/domain/*/repository.go`: Puertos de salida (interfaces)
- Los puertos permiten que el núcleo se comunique sin conocer implementaciones

**Adaptadores de Entrada (Driving Adapters)**
- Invocan la lógica de negocio
- `api/v1/handler/`: Adaptadores HTTP (REST API)
- `api/middleware/`: Middlewares que procesan requests

**Adaptadores de Salida (Driven Adapters)**
- Implementan los puertos definidos por el núcleo
- `internal/infrastructure/repository/external/`: Adaptador para GitHub API
- `internal/infrastructure/repository/redis/`: Adaptador para Redis (caché)
- `internal/infrastructure/repository/db/`: Adaptador para SQLite
- `internal/infrastructure/database/`: Conexión a base de datos
- `internal/infrastructure/cache/`: Conexión a Redis

### Estructura Hexagonal

```
        Adaptadores de Entrada
    ┌─────────────────────────────┐
    │   HTTP REST API Handler     │
    │   (api/v1/handler/)         │
    └──────────────┬──────────────┘
                   │
                   ↓
    ┌──────────────────────────────┐
    │      PUERTO DE ENTRADA       │
    │   (Application Services)     │
    │  internal/application/       │
    │                              │
    │   ┌──────────────────────┐   │
    │   │   NÚCLEO/DOMINIO     │   │
    │   │  internal/domain/    │   │
    │   │  (Entidades + Puertos)│  │
    │   └──────────────────────┘   │
    │                              │
    │    PUERTOS DE SALIDA         │
    │   (Repository Interfaces)    │
    └──────────────┬───────────────┘
                   │
                   ↓
        Adaptadores de Salida
    ┌─────────────────────────────┐
    │  External API | Redis | DB  │
    │ (infrastructure/repository) │
    └─────────────────────────────┘
```

### Principios Aplicados

**Inversión de Dependencias**:
- El núcleo define las interfaces (puertos)
- Los adaptadores implementan esas interfaces
- Las dependencias apuntan hacia adentro (hacia el núcleo)

**Múltiples Adaptadores para un Puerto**:

**Múltiples Adaptadores para un Puerto**:
```
        Puerto (Interface)
    ArticleRepository en domain/
                ↑
                │ implementan
    ┌───────────┼───────────┐
    │           │           │
External    Redis       DB
Adapter     Adapter   Adapter
(GitHub)   (Caché)   (SQLite)
```

Un mismo puerto puede tener múltiples adaptadores. La aplicación elige cuál usar en tiempo de ejecución.

**Beneficios de la Arquitectura Hexagonal**:
- ✅ **Testeable**: El núcleo se prueba sin dependencias externas
- ✅ **Flexible**: Cambiar de base de datos o API sin tocar el núcleo
- ✅ **Independiente**: El dominio no conoce HTTP, frameworks ni infraestructura
- ✅ **Mantenible**: Lógica de negocio separada de detalles técnicos

### Inyección de Dependencias

El contenedor DI (`internal/di/container.go`) conecta puertos con adaptadores:
1. Instancia los adaptadores concretos (implementaciones)
2. Inyecta los adaptadores en los servicios del núcleo
3. Conecta los servicios con los adaptadores de entrada (handlers)
4. Todo el cableado ocurre en tiempo de ejecución

Esto permite que el núcleo hexagonal permanezca ignorante de qué adaptadores específicos se están usando.

## 🔐 Seguridad

- CORS habilitado en modo debug
- Compresión Brotli para reducir ancho de banda
- Certificados TLS/SSL configurables
- Redis con autenticación y TLS
- Logging de todas las requests

## 📦 Dependencias principales

- **github.com/mattn/go-sqlite3**: Driver SQLite
- **github.com/redis/go-redis/v9**: Cliente Redis
- **github.com/andybalholm/brotli**: Compresión Brotli
- **github.com/joho/godotenv**: Variables de entorno
- **github.com/stretchr/testify**: Framework de testing

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commit tus cambios (`git commit -am 'Agrega nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto de ejemplo está disponible como referencia arquitectónica. El proyecto original es privado.

## 👥 Autores

- **Loadept Team**

---

**Disclaimer**: Este repositorio es una versión desensibilizada del proyecto original con fines educativos y de demostración arquitectónica. No contiene información sensible, credenciales ni lógica de negocio propietaria.
