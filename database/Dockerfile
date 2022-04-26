FROM postgres:10.3

# Copiar este archivo y que se ejecute al iniciar el contenedor
# Le asigno 1 porque se ejecutan de manera alfabetica
COPY up.sql /docker-entrypoint-initdb.d/1.sql

# Ejecutar este comando para inicializar la base de datos
CMD ["postgres"]
