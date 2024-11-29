# Trucode3 Challenge

Este proyecto es una aplicación completa que incluye una API, una aplicación frontend y un migrador de datos. La aplicación permite visualizar, filtrar, ordenar y mostrar estadísticas a través de gráficos para analizar datos demográficos basados en un censo de 1994, ayudando a entender qué tipo de personas ganaban más o menos de $50,000 USD.

## Estructura del Proyecto

```plaintext
.
├── api/               # Contiene la lógica del backend y la definición de la API
├── app/               # Aplicación frontend para interactuar con los datos
├── docker-compose.yml # Configuración de contenedores para ejecutar el proyecto
├── migrator/          # Herramientas para migrar y transformar datos
└── tools/             # Utilidades y scripts adicionales para soporte del proyecto
```

## Tecnologias utilizadas

- GO
- GROM
- Gin
- VueJS
- Vue Router
- Pinia
- Taileind CSS
- Vue-Chart JS

## Servicios

- **API**: Proporciona endpoints para la gestión de usuarios, autenticación, datos del censo y estadísticas.
- **App**: Interfaz de usuario para visualizar y analizar datos.
- **Migrator**: Herramienta para migrar datos a la base de datos de forma eficiente implementando GoRutines mediante WorkerPools.
- **DB**: Base de datos PostgreSQL que contiene la informacion del censo asi como de los usuarios.

## Requisitos

- Docker
- Docker Compose

## Configuración

1. Clonar el repositorio.
2. Ejecutar `docker-compose up` para levantar los servicios.

## Contacto

Para cualquier consulta, puedes realizarla a melisadavila2001@gmail.com.
