---
layout: ../../../layouts/GoProjects.astro
title: db-mcp - Servidor MCP para consultas multi-base de datos
description: Documentación de db-mcp, un servidor MCP para consultas multi-base de datos.
keywords: [db-mcp, postgres, mysql, mssql, mcp, servidor, solo lectura, base de datos]
goImport: loadept.com/db-mcp git https://github.com/loadept/db-mcp
goSource: loadept.com/db-mcp https://github.com/loadept/db-mcp https://github.com/loadept/db-mcp/tree/master{/dir} https://github.com/loadept/db-mcp/blob/master{/dir}/{file}#L{line}
---

[![Database MCP](https://github.com/loadept/db-mcp/actions/workflows/release_workflow.yml/badge.svg)](https://github.com/loadept/db-mcp/actions/workflows/release_workflow.yml)
[![Database MCP](https://github.com/loadept/db-mcp/actions/workflows/docker_latest_workflow.yml/badge.svg)](https://github.com/loadept/db-mcp/actions/workflows/docker_latest_workflow.yml)

**Lenguaje:** Go 1.25+  
**Tipo:** Ejecutable binario

## ⚠️ ADVERTENCIA IMPORTANTE
**NO DEBE EJECUTARSE EN UNA BASE DE DATOS DE PRODUCCIÓN**

Este servidor está diseñado para entornos de desarrollo y prueba. No se recomienda su uso en bases de datos de producción sin revisión exhaustiva de seguridad.

## Descripción
Servidor MCP para consultas multi-base de datos. Soporta PostgreSQL, MySQL y SQL Server (MSSQL) mediante interfaz pluggable. 

**Nota sobre MSSQL:** SQL Server no soporta transacciones de solo lectura a nivel del driver, por lo que las consultas se ejecutan sin encapsulación transaccional. Esto significa que los cambios se aplican inmediatamente.

## Instalación

### Última Versión (Latest - Último Commit)
Para obtener las últimas características y mejoras (puede contener cambios no probados):

#### Con Go
```bash
# Instala el binario en $GOBIN (o $GOPATH/bin)
go install loadept.com/db-mcp@latest

# También es posible ejecutar directamente sin instalar:
go run loadept.com/db-mcp@latest
```

#### Con Docker
```bash
docker run --rm -i loadept/db-mcp:latest -u "postgres://usuario:contraseña@host:puerto/basedatos?sslmode=disable" -e postgres
docker run --rm -i loadept/db-mcp:latest -u "usuario:contraseña@tcp(host:puerto)/basedatos" -e mysql
docker run --rm -i loadept/db-mcp:latest -u "sqlserver://usuario:contraseña@host:puerto" -e mssql
```

### Versión Estable (Releases Tagueadas)
Para versiones estables, probadas y seguras:

#### Binarios Precompilados
Descarga el ejecutable precompilado para tu sistema operativo desde las [releases](https://github.com/loadept/db-mcp/releases).

#### Con Docker
```bash
docker run --rm -i loadept/db-mcp:v0.1.0 -u "postgres://usuario:contraseña@host:puerto/basedatos?sslmode=disable" -e postgres
docker run --rm -i loadept/db-mcp:v0.1.0 -u "usuario:contraseña@tcp(host:puerto)/basedatos" -e mysql
docker run --rm -i loadept/db-mcp:v0.1.0 -u "sqlserver://usuario:contraseña@host:puerto" -e mssql
```

## Ejecución

### PostgreSQL
```bash
# Linux/macOS
./db-mcp -u "postgres://usuario:contraseña@localhost:5432/basedatos?sslmode=disable" -e postgres

# Windows
db-mcp.exe -u "postgres://usuario:contraseña@localhost:5432/basedatos?sslmode=disable" -e postgres
```

### MySQL
```bash
# Linux/macOS
./db-mcp -u "root:godylody@tcp(raspberry.local:3306)/mcp_tests" -e mysql

# Windows
db-mcp.exe -u "root:godylody@tcp(raspberry.local:3306)/mcp_tests" -e mysql
```

### SQL Server (MSSQL)
```bash
# Linux/macOS
./db-mcp -u "sqlserver://sa:Password123@localhost:1433?database=mcp_tests" -e mssql

# Windows
db-mcp.exe -u "sqlserver://sa:Password123@localhost:1433?database=mcp_tests" -e mssql
```

## Opciones de Línea de Comandos
```bash
-u        URI de conexión a la base de datos (requerido, SIEMPRE ENTRE COMILLAS)
          PostgreSQL: postgres://usuario:contraseña@host:puerto/basedatos?sslmode=disable
          MySQL:      usuario:contraseña@tcp(host:puerto)/basedatos
          MSSQL:      sqlserver://usuario:contraseña@host:puerto?database=basedatos

-e        Motor de base de datos (requerido)
          Valores: postgres, mysql, mssql

-version  Muestra la versión de la aplicación
```

## Identificación de Versiones
Para saber qué tipo de versión tienes instalada, ejecuta `db-mcp -version`:

- **`dev`**: Última versión (latest) - Contiene el último commit, puede ser inestable
- **`v0.1.0`** (o similar): Versión estable - Release tagueada, probada y segura

## Herramientas
- `execute_query`: Ejecuta consultas SELECT en el motor de BD configurado (máx. 50 filas)

## Compilación desde el código fuente
```bash
go build -o db-mcp .
```
