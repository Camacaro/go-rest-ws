# ==== Primera etapa: crear la imagen ====

# Variable con la version de GO
ARG GO_VERSION=1.18

# Descargar la version de GO y llamarlo builder
FROM golang:${GO_VERSION}-alpine as builder

# Variables de entorno en Go dentro del contenedor
# Ve directo al modulo de go a descargar las dependencias
RUN go env -w GOPROXY=direct

# Instalar git
RUN apk add --no-cache git

# Certificado de seguridad para el repositorio de Go
RUN apk --no-cache add ca-certificates && update-ca-certificates

# Crearemos un directorio de trabajo para el proyecto
# donde ejecutaremos todos estos comandos
WORKDIR /src

# Copiar los archivo de go.mod y go.sum y copiarlos al directorio de trabajo
COPY ./go.sum ./go.mod ./

# Instalar las deoencias de go
RUN go mod download

# Copiar todos los directorios de trabajo a la raiz del contenedor
COPY ./ ./

# Construir la aplicacion
# CGO_ENABLED es para qeu go pueda usar la libreria C para 
# compilar pero en este caso no , ya que esta version de go que estamos descargando no lo soporta
# installsuffix es necesraio para que el ejecutable funcione en el contenedor
# Finalmmente le ponemos la ruta o el nombre del ejecutable
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /go-rest-ws

# ==== Segunda etapa:  ====

# Estas copias son de la imagen del builder a la imagen del runner

# Esta es la encargada de ejecutar nuestro servidor o aplicacion
FROM scratch AS runner

# Copiar los certificados descargados al contenedor
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copiar el archivo de entorno a la ruta principal
COPY .env ./

# Copiar el ejeccutable a la ruta principal
COPY --from=builder /go-rest-ws /go-rest-ws

# Puerto de escucha, expuesto
EXPOSE 5050

# Comando a ejecutar
ENTRYPOINT [ "/go-rest-ws" ]