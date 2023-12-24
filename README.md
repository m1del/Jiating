# Jiating
Website for Jiating Lion and Dragon

## Getting Started
##### Prerequisites
- Node
- Docker & Docker Compose
- Go
- Psql (optional)

### Frontend
Install dependencies and run the development client.
```
cd frontend/
npm i
npm run dev
```
### Backend
Install dependencies, and run the server.
- use make watch for an easy one time run
```
cd backend/
go mod tidy
make run
```
otherwise, use `make watch` to enable hot reload.
see Makefile for more options

#### Database
For this part, have psql 15.5 installed, and the backend Go server running
```
psql -h 127.0.0.1 -p 5432 -U [DB_USER] -d [DB_DATABASE]
```
Congrats, now you can do sql yessir.


