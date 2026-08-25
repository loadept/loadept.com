---
title: Herramientas de compresión en Linux
date: 2025-05-26
keywords: [linux, terminal, gzip, xz, 7zip]
---

En Linux, existen varias herramientas de compresión como **gz** o **zip**, algunas más robustas que otras. A continuación, revisaremos las más comunes y mostraremos ejemplos de compresión de archivos y directorios, pero antes, un pequeño cuadro de comparación.

Este cuadro muestra cuál es más robusto que otro:

| Nombre | Algoritmo | Compresion |   Velocidad    |              Usos               |
| :----: | :-------: | :--------: | :------------: | :-----------------------------: |
|  7zip  |   LZMA2   | 🔥 Máxima  |    🐢 Lento    | Archivos o paquetes muy grandes |
|   xz   |   LZMA    |  🔥 Alta   | 🐌 Medio-lento |  Distribuciones Linux, backups  |
|  gzip  |  Deflate  |  ⚡ Media   |   🚀 Rápido    |   Logs, transmiciones en vivo   |
|  zip   |  Deflate  |  🧊 Baja   |   🚀 Rápido    |    Compatibilidad universal     |

# Usos
## Compresión con gz
Comencemos con una de las herramientas más comunes: **gzip**.

**gzip** nos permite comprimir un solo archivo. Veamos cómo comprimir un archivo llamado `backup.sql`:
```bash
gzip backup.sql
```

>Este comando comprimirá el archivo `backup.sql` y generará el archivo `backup.sql.gz`. Cabe destacar que, por defecto, **gzip** elimina el archivo original. Si queremos mantener el archivo original y generar solo el archivo comprimido, podemos usar la opción `-k` o `--keep`

```bash
gzip -k backup.sql
```
De este modo, se generará `backup.sql.gz` sin eliminar el archivo original `backup.sql`.

## Descompresión con gz
Para descomprimir un archivo comprimido con **gzip**, usamos la opción `-d` o el comando `gunzip`:
```bash
gzip -d backup.sql.gz
```
o
```bash
gunzip backup.sql.gz
```

>Esto descomprimirá el archivo pero eliminará el archivo `.gz`. Si queremos conservar tanto el archivo comprimido como el descomprimido, usamos `-k` nuevamente.

```bash
gzip -dk backup.sql.gz
```
o
```bash
gunzip -k backup.sql.gz
```

---

## Compresión con xz
La compresión con **xz** es muy similar a la de **gzip**. Para comprimir un archivo, usamos el siguiente comando:
```bash
xz backup.sql
```

>Esto generará `backup.sql.xz` pero al igual que **gzip** eliminará el archivo original. Si queremos mantener el archivo original, usamos `-k`

```bash
xz -k backup.sql
```

## Descompresión con xz
Para descomprimir un archivo `.xz`, podemos usar la opción `-d` o el comando `unxz`:
```bash
xz -d backup.sql.gz
```
o
```bash
unxz backup.sql.gz
```

Al igual que con **gzip**, si queremos conservar el archivo `.xz` original, agregamos `-k`:
```bash
xz -dk backup.sql.xz
```
o
```bash
unxz -k backup.sql.gz
```

---

## Compresión con 7zip
**7zip** es una herramienta de compresión más robusta. Para usarla, utilizamos el comando `7z`. Para comprimir un archivo, usamos la opción `a` (agregar):
```bash
7z a backup.7z backup.sql
```
Este comando comprimirá el archivo `backup.sql` en `backup.7z`. También podemos comprimir directorios:
```bash
7z a directory.7z ~/.config/nvim
```
Esto habrá comprimido el directorio `nvim`.

>Es importante señalar que, a diferencia de **gzip** y **xz**, **7zip** no elimina el archivo original.

## Descompresión con 7zip
Para descomprimir con **7zip**, utilizamos la opción `x`:
```bash
7z x backup.7z
```
Esto descomprimirá el archivo `backup.7z` sin eliminarlo.

## Usos de tar
Aunque **gzip** y **xz** son útiles para comprimir archivos individuales, no pueden comprimir directorios completos. Para ello, usamos **tar**.

**tar** es una herramienta que permite agrupar varios archivos o directorios en un solo archivo, sin comprimir. **tar** no realiza la compresión por sí sola, solo genera un archivo denominado "**tarball**" con la extensión `.tar`.

Para archivar varios archivos con **tar**, usamos el siguiente comando:
```bash
tar -cvf files.tar file1.txt file2.txt file3.txt
```
Las opciones son las siguientes:
- `-c`: Crea un nuevo archivo.
- `-v`: Muestra información detallada.
- `-f`: Especifica el nombre del archivo de salida (en este caso, `files.tar`).

También podemos archivar un directorio completo:
```bash
tar -cvf files.tar directory/
```
Simplemente indicamos el directorio que deseamos archivar.

## Des-archivar con tar
Para desarchivar un archivo `.tar`, usamos la opción `-x` (extract):
```bash
tar -xvf files.tar
```

>Este comando extraerá los archivos del archivo tar sin eliminarlos.

---

## Compresión con tar
Como **gzip** y **xz** solo comprimen un archivo individual, podemos combinarlas con **tar** para comprimir varios archivos o directorios a la vez.

Primero, archivamos los archivos o directorios con **tar**:
```bash
tar -cvf files.tar file1.txt file2.txt
```
o
```bash
tar -cvf files.tar directory/
```

Luego, comprimimos el archivo `.tar` usando **gzip** o **xz**:
```bash
gzip -k files.tar
```
o 
```bash
xz -k files.tar
```
Esto nos habrá generado `files.tar.gz` o `files.tar.xz` y con `-k` evitamos que se elimine el archivo `.tar` original.

Esta fue la manera más lenta de comprimir varios ficheros o directorios, si preferimos hacerlo todo en un solo paso, usamos la opción `-z` para **gzip** o `-J` para **xz** al momento de crear el archivo **tar**:

Pero tar nos permite acelerar este paso:
```bash
tar -czvf files.tar.gz file1.txt file2.txt
```
o
```bash
tar -czvf files.tar.gz directory/
```
Con la nueva flag `-z` lo que haremos es comprimir el archivo final `tar` (`tarball`) con `gzip`.

También podemos hacerlo con `xz`:
```bash
tar -cJvf files.tar.gz file1.txt file2.txt
```
o
```bash
tar -cJvf files.tar.gz directory/
```
Con la flag `-J` habremos comprimido el archivo `tar` (`tarball`) con `xz`.

## Descompresión con tar
Para descomprimir con `tar` simplemente debemos reemplazar la flag `-c` (create) por `-x` (extract).

Para archivos `gzip`
```bash
tar -xzvf files.tar.gz
```

Para archivos `xz`
```bash
tar -xJvf files.tar.gz
```

>**tar** no elimina los archivos al archivar o desarchivar, por lo que no necesitamos la opción `-k`.
