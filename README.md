# API Rest em Go com Gin, Gorm e Docker

Este repositório contém uma aplicação que desenvolve uma API Rest utilizando o framework Gin na linguagem Go. A aplicação utiliza o Gorm como ORM e o PostgreSQL rodando em um contêiner Docker para gerenciar o banco de dados. Este projeto faz parte da formação em Linguagem Go oferecida pela Alura.

## Tecnologias Utilizadas

- [Go](https://golang.org/) (versão 1.16 ou superior)
- [Gin](https://gin-gonic.com/) - Framework web em Go
- [Gorm](https://gorm.io/) - ORM para Go
- [Docker](https://www.docker.com/) - Plataforma de contêineres
- [PostgreSQL](https://www.postgresql.org/) - Sistema de gerenciamento de banco de dados relacional

## Pré-requisitos

Antes de executar esta aplicação, certifique-se de ter o seguinte instalado em sua máquina:

- Go
- Docker
- Docker Compose

## Configuração do Banco de Dados

A aplicação utiliza o PostgreSQL rodando em um contêiner Docker. Você pode configurar o banco de dados utilizando o Docker Compose.

### Docker Compose

Crie um arquivo `docker-compose.yml` na raiz do projeto com o seguinte conteúdo:

```yaml
version: '3.8'
services:
  db:
    image: postgres:latest
    environment:
      POSTGRES_USER: seu_usuario
      POSTGRES_PASSWORD: sua_senha
      POSTGRES_DB: seu_banco_de_dados
    ports:
      - "5432:5432"

Inicie o contêiner com o seguinte comando:
```bash
   docker-compose up -d
```

## Configuração da Aplicação
Antes de executar a aplicação, você precisará configurar as variáveis de ambiente necessárias. Para isso, crie um arquivo .env na raiz do projeto e defina as seguintes variáveis:

```bash
  DB_HOST=localhost
  DB_PORT=5432
  DB_USER=seu_usuario
  DB_PASSWORD=sua_senha
  DB_NAME=seu_banco_de_dados
```

## Instalando Dependências

Para instalar as dependências do projeto, incluindo o Gin, execute o seguinte comando:

```bash
   go get -u github.com/gin-gonic/gin
```

## Executando a Aplicação

Para executar a aplicação, siga estas etapas:

1 - Instale as dependências do projeto:
