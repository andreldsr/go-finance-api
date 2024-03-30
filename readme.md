
# Finance Manager API

API para criação e controle de transações financeiras, acompanhamento de gastos e visualização de estatísticas.




## Setup

Para iniciar o uso da aplicação é necessário que o banco de dados esteja sendo executado e atualizado. Para isso primeiro utilize o Docker Compose para subir as dependências da aplicação.

```bash
  docker compose up -d
```

Esse comando vai executar o banco de dados, mas ele ainda estará limpo, se for a primeira vez executando esse comando.

> **Atenção** <br/>
> Caso o banco de dados já esteja preenchido não é necessário executar os próximos passos

Após isso é necessário preparar o ambiente para a execução das migrações com Goose

```bash
export GOOSE_DRIVER=postgres;
export GOOSE_DBSTRING="host=localhost port=5432 user=finance-manager dbname=finance-manager password=finance-manager sslmode=disable";
export GOOSE_MIGRATION_DIR=./database/migrations;
```
Para executar as migrações basta executar o comando
```bash
goose up
```
