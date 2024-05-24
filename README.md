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
````
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

- Instale as dependências do projeto:

```bash
   go mod tidy
```

- Compile o código:

```bash
   go build -o nome_do_executável
```

- Execute a aplicação:

```bash
   ./nome_do_executável
```

A aplicação estará disponível em http://localhost:8080.


## Capítulo 01 - Instalando e crando a primeira rota com Gin

- Iniciamos o projeto do zero, criando uma pasta e instalando os módulos necessários para criar uma API com Gin;
- Instalamos o gin em nosso projeto com o comando:
```bash
  go get -u github.com/gin-gonic/gin
```
- Criamos um endpoint, que recebe uma requisição GET, retornando um Json exibindo recursos de um aluno.

## Capítulo 02 - Modularizando o código, modelos e banco de dados

- Modularizamos nosso código criando pacotes de controles e rotas, tornando nosso código mais fácil de manter e editar;
- Aprendemos como pegar informações passadas por parâmetros e retornar uma mensagem personalizada;
- Criamos a struct de aluno, que vamos disponibilizar como recurso da nossa API.

## Capítulo 03 - Struct, banco de dados e ORM

- Instalamos o Gorm para não escrever código sql, facilitando a comunicação da aplicação com o banco de dados;
- Conectamos a API com banco de dados e realizamos uma migração com base na struct de aluno;
- Alteramos o controle para exibir os alunos registrados no banco de dados!

## Capítulo 04 - Implementando rotas HTTP

- Alteramos o controller para exibir alunos do banco de dados;
- Criamos um endpoint para exibir alunos por ID;
- Alteramos o comportamento da API para exibir uma mensagem quando o ID do aluno não for encontrado;

## Capítulo 05 - Deletando, editando e buscando alunos

- Adicionamos um endpoint com método Delete para deletar um aluno e removê-lo do banco de dados;
- Adicionamos um endpoint com método Patch para atualizar o cadastro de um aluno;
- Criamos um endpoint para buscar alunos pelo número do CPF;


--------------

# Go: Validações, testes e páginas HTML

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
````
Inicie o contêiner com o seguinte comando:
```bash
   docker-compose up -d
```
## Capítulo 01 - Instalando e criando a primeira rota com Gin

- Carregamos o projeto base e criamos a imagem do banco de dados no Docker;
- Criamos nossas validações na struct de Aluno, garantindo que um campo não fique em branco e tenha uma quantidade específica de caracteres;
- Aplicamos essa validação no controller no momento que criamos ou editamos um aluno.

## Capítulo 02 - Testes

- Realizamos um teste no Postman que verifica o statusCode de uma resposta;
Criamos nosso primeiro teste em Go, o TestFalhador;
- Escrevemos um teste que verifica o endpoint de Saudação da API;
- Instalando o assert e alteramos o código verificando o corpo da resposta.


## Contribuição
Se você encontrar problemas ou tiver sugestões para melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença
Este projeto está licenciado sob a Licença MIT.