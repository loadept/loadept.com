---
layout: ../../../layouts/GoProjects.astro
title: pg-mcp - Servidor MCP para consultas PostgreSQL de solo lectura
description: Documentación de pg-mcp, un servidor MCP para consultas PostgreSQL de solo lectura.
keywords: [pg-mcp, postgres, mcp, servidor, solo lectura, base de datos]
goImport: loadept.com/pg-mcp git https://github.com/loadept/pg-mcp
goSource: loadept.com/pg-mcp https://github.com/loadept/pg-mcp https://github.com/loadept/pg-mcp/tree/master{/dir} https://github.com/loadept/pg-mcp/blob/master{/dir}/{file}#L{line}
---

[![Postgres MCP](https://github.com/loadept/pg-mcp/actions/workflows/pg_workflow.yml/badge.svg)](https://github.com/loadept/pg-mcp/actions/workflows/pg_workflow.yml)
[![Postgres MCP](https://github.com/loadept/pg-mcp/actions/workflows/docker_latest_workflow.yml/badge.svg)](https://github.com/loadept/pg-mcp/actions/workflows/docker_latest_workflow.yml)

**Lenguaje:** Go 1.25+  
**Tipo:** Ejecutable binario

## Descripción
Servidor MCP para consultas PostgreSQL de solo lectura.

## Instalación

### Última Versión (Latest - Último Commit)
Para obtener las últimas características y mejoras (puede contener cambios no probados):

#### Con Go
```bash
# Instala el binario en $GOBIN (o $GOPATH/bin)
go install loadept.com/pg-mcp@latest

# También es posible ejecutar directamente sin instalar:
go run loadept.com/pg-mcp@latest
```

#### Con Docker
```bash
docker run --rm -i loadept/pg-mcp:latest -u "postgres://usuario:contraseña@host:puerto/basedatos?sslmode=disable"
```

### Versión Estable (Releases Tagueadas)
Para versiones estables, probadas y seguras:

#### Binarios Precompilados
Descarga el ejecutable precompilado para tu sistema operativo desde las [releases](https://github.com/loadept/pg-mcp/releases).

#### Con Docker
```bash
docker run --rm -i loadept/pg-mcp:v0.3.1 -u "postgres://usuario:contraseña@host:puerto/basedatos?sslmode=disable"
```

## Ejecución
```bash
# Linux/macOS
./pg-mcp -u "postgres://usuario:contraseña@localhost:5432/basedatos?sslmode=disable"

# Windows
pg-mcp.exe -u "postgres://usuario:contraseña@localhost:5432/basedatos?sslmode=disable"
```

## Opciones de Línea de Comandos
```bash
-u    URI de conexión a PostgreSQL (requerido)
      Formato: postgres://usuario:contraseña@host:puerto/basedatos?sslmode=disable

-version    Muestra la versión de la aplicación
```

## Identificación de Versiones
Para saber qué tipo de versión tienes instalada, ejecuta `pg-mcp -version`:

- **`dev`**: Última versión (latest) - Contiene el último commit, puede ser inestable
- **`v0.3.1`** (o similar): Versión estable - Release tagueada, probada y segura

## Herramientas
- `execute_query`: Ejecuta consultas SELECT (máx. 50 filas)
- `get_table_info`: Información de estructura de tabla
- `list_tables`: Lista tablas por esquema (paginado)

## Compilación desde el código fuente
```bash
go build -o pg-mcp .
```
