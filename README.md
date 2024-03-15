# GoBrax

GoBrax é uma aplicação Go que gerencia cadastro e vinculação entre motoristas e veículos, fornecendo um CRUD completo para ambas as entidades.

## Pré-Requisitos
Docker
Docker Compose
Go (para desenvolvimento local sem Docker)

## Configuração
```
git clone https://github.com/deduardolima/gobrax.git
cd gobrax
```

## Instalação e Execução com Docker
Construa e inicie os containers:
```
docker-compose up --build
```

isso irá construir a imagem do aplicativo e iniciar os serviços definidos no docker-compose.yml, incluindo o banco de dados e o aplicativo.

A aplicação estará acessível em http://localhost:8080.

## Desenvolvimento Local sem Docker
Se você preferir rodar a aplicação diretamente em sua máquina para desenvolvimento, siga estas etapas:

Garanta que você tenha o Go instalado e configurado corretamente em sua máquina.

Instale as dependências do Go:
```
go mod tidy
```
```
go run cmd/server/main.go
```
## Endpoints da API
A aplicação oferece vários endpoints para gerenciar motoristas e veículos. Aqui estão alguns exemplos:

- POST /drivers: Cria um novo motorista.
- GET /drivers: Obtém a lista de  motorista.
- GET /drivers/{id}: Obtém um motorista por ID.
- PUT /drivers/{id}: Atualiza os dados de um motorista com ID.
- DELETE /drivers/{id}: Deleta um motorista com ID.

- POST /vehicles: Cria um novo veículo.
- GET /vehicles: Obtém a lista de veiculos e de motoristas vinculados respectivamente.
- GET /vehicles/{id}: Obtém um veículo e do motorista vinculado por ID.
- PUT /vehicles/{id}: Atualiza dados de um veículo.
- DELETE /vehicles/{id}: Deleta um veiculo com ID.



