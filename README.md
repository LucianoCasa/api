# Desafio: Listagem de Orders REST/gRPC/GraphQL

## Como executar

```bash
docker compose up --build
```

## Endpoints

- REST: http://localhost:8080/order
- gRPC: porta 50051 
- GraphQL: http://localhost:8081/graphql

## Banco de dados

- MySQL 8
- Host interno: `db:3306`
- Usuário: `user`
- Senha: `pass`
- Banco: `ordersdb`


## Gerar código da estrutura
```cmd
# gRPC
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative internal/api/grpcapi/pb/order.proto

# GraphQL
gqlgen generate
```

## Testes

### Testar REST
Use o arquivo `api.http` para testar criação e listagem via REST.


### Testar gRPC
```cmd
evans -r repl --host localhost --port 50051 --package order
show service
call ListOrders
```

### Testar GraphQL
http://localhost:8081/playground

```grapql
mutation createOrder {
  createOrder(
    customer_name: "Zélia"
    total: 200.50
    created_at: "2025-06-22T10:00:00"
  ) {
    customer_name
    total
  }
}
```

```grapql
query { 
  listOrders { 
    id 
    customer_name 
  } 
}
```
