## Para ejecutar la migracion de los datos hacia la BD

### Tener el servicio de PostgreSQL levantado previamente:

```
El archivo docker-compose.ylm que levanta el servicio esta ubicado en la carpeta tools.
```

### Tener archivo .env con las variables:

```plaintext
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=myuser
   DB_PASSWORD=mypassword
   DB_NAME=mydatabase
   DB_SSLMODE=disable
```

### Ejecutar en la consola:

```bash
    go run .
```
