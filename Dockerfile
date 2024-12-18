# Importando a imagem oficial do golang
FROM golang:1.23-alpine

# Definindo o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiando o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod ./

# Baixando as dependências
RUN go mod download

# Copiando o código-fonte para o diretório de trabalho
COPY . .

# Compilando a aplicação
RUN go build -o /app/main ./cmd/main.go

# Verificando o conteúdo do diretório /app
RUN ls -la /app

# Definindo o comando de entrada para rodar a aplicação
CMD ["/app/main"]