FROM golang:1.19.3

WORKDIR /backEnd

# Copia solo los archivos necesarios para descargar módulos y construir
COPY go.mod go.sum ./

# Descarga los módulos
RUN go mod download

# Copia el resto de los archivos
COPY . .

# Construye la aplicación
RUN go build -o main .

EXPOSE 8090

CMD [ "/backEnd/main" ]

