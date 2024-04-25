# payment-gateway-go
Payment Gateway in Golang





# Instruciones para ejecucion de API

Estando en el folder `/investment_api`

Ejecutar en una terminal:

- `make`
- `make run`

Si no tiene soporte para make:

- `go run main.go`

Si desea eliminar el historial guardado en SQLite, eliminar el file `gorm.db` que se genera

Para la generacion de mock (`brew install mockery`), ultilice:

- `make mock_gen` o `mockery`

# Add Homebrew to your PATH

eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"


cat .env.dev > .env
source .env

brew install golang-migrate


sudo apt-get install -y postgresql-client
psql --version


https://www.sandbox.paypal.com/checkoutnow?token=