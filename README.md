# Plutus

The Plutus banking api supports the following operations

| Resource  | Operations |
| ------------- | ------------- |
| User          | CRUD          |
| Account       | CR            |
| Transaction   | CR            |

### Pre-requisites

- `docker (tested with 25.0.3)`
- `docker-compose (tested with 2.23.3)`

### Running the server

Use `make up` to run the api server and `make down` to shutdown it down.

The api will become available at `http://localhost:3000`. 

You can access the webpage [http://localhost:3000/swagger](http://localhost:3000/swagger) to access the banking api swagger docs where you can perform requests.

The database is initialized with some seed entries accessible at `/banking-api/seeds/seed.go`. The system is ready to accept 2 currencies `USD` and `EUR`.

### Improvements

- **Create integration tests**: currently the most crucial logic is on the repository layer which does not contain any tests coverage. All tests were done in service layer which does not contain any special logic other than CRUD.

- **Create api e2e tests**: creating api e2e tests which simulate an incoming request and expect a *correct* response would also give us some guarantees of reliability

- **Database consistency**: some database fields should be reconsidered namely the transaction type field which right now is only a string but should be migrated to an enum type (`CREDIT / DEBIT`) to better safeguard the data

- **Microservice Possibility**: if the system complexity and growth is expected it could make sense to split the banking-api service into smaller microservices such as user-api, for user specific operation, and account-api, for operations and transaction management.

- **Error messages**: considering the addition of standard error messages between layers (`rest/services/repositories`) would be benificial. It could also enable error message translation depending on the user location