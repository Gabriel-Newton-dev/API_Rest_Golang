<h1>API REST GOLANG</h1> 

</div>
<div style="display: inline_block"><br>
  <img align="center" height="30" width="40" img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" />
  <img align="center" height="30" width="40" img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/html5/html5-original.svg" />
  <img align="center" height="30" width="40" img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/css3/css3-original.svg" />
  <img align="center" height="30" width="40" img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original-wordmark.svg" />
</div>

> <p><h2>Status do Projeto: em andamento</h2></p>

### Tópicos 

:small_blue_diamond: [Descrição do projeto](#descrição-do-projeto)

:small_blue_diamond: [Funcionalidades](#funcionalidades)

:small_blue_diamond: [Pré-requisitos](#pré-requisitos)

:small_blue_diamond: [Como rodar a aplicação](#como-rodar-a-aplicação-arrow_forward)


## Descrição do projeto 

<p align="justify">
  Codei essa Api usando apenas Golang.  

Desenvolvi uma API rest do zero, criando rotas, endpoints, modelos etc. Conectei a minha aplicação com um banco de dados Postgres o qual configurei  usando o Docker-compose. Além disso, implementei o CRUD completo(GET, PATCH, POST, DELETE), fiz rotas para busca por ID, CPF e RG.

</p>

## Funcionalidades

:heavy_check_mark: Crud completo no banco de dados; 

## Pré-requisitos

:warning: [Golang](https://go.dev/dl/)
:warning: [Docker](https://www.docker.com/products/docker-desktop/)

## Dependências a serem instaladas no terminal na pasta do projeto:

:heavy_check_mark: go get -u github.com/gin-gonic/gin)

:heavy_check_mark: go get -u gorm.io/gorm

:heavy_check_mark: go get github.com/spf13/viper

## Como rodar a aplicação :arrow_forward:

No terminal, clone o projeto: 

```
$ git clone https://github.com/Gabriel-Newton-dev/API_Rest_Golang.git
```
Ainda no terminal execute:

```
$ go run main.go
```

## Casos de Uso

Para utilizar a mesma além de fazer as instalações necessárias conforme dito acima, se faz necessário criar e apontar as suas variáveis de ambiente privadas, para que você consiga configurar o docker-compose e subir o seu container com o respectivo banco de dados. 

## Iniciando/Configurando banco de dados

# Para iniciar o banco de dados se faz necessário apontar as variaveis de ambiente, após rodar o comando:

```
$ docker-compose up
```

## Linguagens, dependencias e libs utilizadas :books:

- [Golang](https://go.dev/dl/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Viper](https://github.com/spf13/viper)
- [Gorm](https://gorm.io/)

## Resolvendo Problemas :exclamation:

Em [issues](https://github.com/Gabriel-Newton-dev/API_Rest_Golang/issues) foram abertos alguns problemas gerados durante o desenvolvimento desse projeto e como foram resolvidos. 

## Desenvolvedor:

- [Gabriel Newton](https://gabriel-newton-dev.github.io/)

## Licença 

The [MIT License]() (MIT)

Copyright :copyright: Ano - Titulo do Projeto
