# Aplicação Web com Go

## Rotas da Aplicação

| Método | Rota            | Handler      | Descrição          |
|--------|-----------------|--------------|---------------------|
| POST   | /notes/create   | noteCreate   | Cria uma anotação   |

## Modelo do Banco de Dados

### Tabela: NOTES

COMANDO PARA GERAR AS TABLES: 
- migrate create -ext sql -dir db/migrations -seq create_users_table
- migrate -database postgres://postgres:root@localhost:5432/postgres?sslmode=disable -path db/migrations up

| Campo      | Tipo       | Restrições        |
|------------|------------|-------------------|
| id         | BIGSERIAL  | PK, NOT NULL      |
| title      | TEXT       | NOT NULL          |
| content    | TEXT       | -                 |
| color      | TEXT       | NOT NULL          |
| created_at | TIMESTAMP  | -                 |
| updated_at | TIMESTAMP  | -                 |
