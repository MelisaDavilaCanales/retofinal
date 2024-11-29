## Para ejecutar la migracion de los datos hacia la BD

Esta herramienta migra datos demográficos a la base de datos PostgreSQL, haciendo uso de Workers (GoRutines), puedes definir la cantidad de Workers y el buffer de estos en ./shared/const.go dependiendo de la capacidad de tu computador.

### Rquisito

Tener el servicio de PostgreSQL levantado previamente:

```
El archivo docker-compose.ylm que levanta el servicio esta en la raiz del proyecto.
```

## Ejecución

1. Clonar el repositorio.
2. Navegar al directorio `migrator`.
3. Ejecutar `go mod download` para instalar las dependencias.
4. Ejecutar `go run .` para iniciar el migrador.
