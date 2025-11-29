# loadept.com

> Blog personal, portfolio y colección de herramientas open source para desarrolladores.

![loadept banner](https://loadept.com/static/img/loadept.webp)

Este repositorio contiene el código fuente del frontend de [loadept.com](https://loadept.com), construido con **Astro**, **React (Preact)** y **TailwindCSS**.

## 🚀 Stack Tecnológico

- **Framework:** [Astro 5](https://astro.build) (Static Site Generation)
- **UI Library:** [Preact](https://preactjs.com/) (para componentes interactivos)
- **Estilos:** [TailwindCSS v4](https://tailwindcss.com)
- **Contenido:** Markdown & MDX (Content Collections)
- **Iconos:** [Lucide](https://lucide.dev)
- **Tipografía:** Fira Code (Nerd Font)

## 📂 Estructura del Proyecto

El proyecto sigue una arquitectura modular para facilitar la escalabilidad:

```text
src/
├── content/          # Colecciones de contenido (Markdown)
│   ├── posts/        # Artículos del blog
│   └── resources/    # Documentación de herramientas y paquetes
├── modules/          # Módulos funcionales (Feature-based architecture)
│   ├── home/         # Lógica de la página de inicio
│   ├── tools/        # Herramientas interactivas (PDF, Imágenes, etc.)
│   └── shared/       # Componentes y utilidades compartidas
├── pages/            # Rutas de Astro (File-based routing)
│   ├── index.astro   # Home
│   ├── [tool].astro  # Generador de páginas de herramientas
│   └── ...
└── layouts/          # Plantillas base (SEO, Header, Footer)
```

## 🛠️ Características Principales

- **Arquitectura de Islas:** Hidratación parcial con `client:load` solo donde es necesario.
- **Rutas Dinámicas:** Generación automática de páginas para posts y herramientas desde archivos.
- **SEO Optimizado:** Metadatos dinámicos, Open Graph y JSON-LD.
- **View Transitions:** Navegación suave tipo SPA sin sacrificar el rendimiento SSG.
- **Herramientas Integradas:**
  - Compresor de PDF (WASM/API)
  - Documentación de paquetes Go
  - Blog técnico

## 🧞 Comandos

| Comando | Acción |
| :--- | :--- |
| `pnpm install` | Instala las dependencias |
| `pnpm dev` | Inicia el servidor de desarrollo en `localhost:4321` |
| `pnpm build` | Compila el sitio para producción en `./dist/` |
| `pnpm preview` | Previsualiza la build localmente |

## 📄 Licencia

Este proyecto es Open Source y está disponible bajo la licencia [MIT](LICENSE).

---

Hecho con ❤️ y mucho ☕ por [loadept](https://loadept.com/about).
