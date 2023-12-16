# paper.id - disbursement

## Getting Started
### Prerequisites
- Go installed
- Docker installed

### Installation

- clone postgresql from docker images
```bash
docker pull postgres
```

- run postgresql image with user and password : root, database name : disbursement
```bash
docker run --name my_postgres_container -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=disbursement -p 5432:5432 -d postgres
```

- run application ( without docker )
```bash
go run cmd/main.go
```

- run application ( with docker )
```bash
docker compose up -d
```

### Postman

you might test localhost with example body payload on this shared Postman :

[CLICK POSTMAN LINK ](https://documenter.getpostman.com/view/5266147/2s9Ykn8hPE)