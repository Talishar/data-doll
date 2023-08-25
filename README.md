# data-doll

Data Doll is a new backend for Talishar. The game engine itself will still be where it is. But all the out-of-game stuff will be hosted in data-doll, this includes things like user info, decklists, password recovery; Patreon, fabdb and fabrary integrations; cosmetics, match history, and so on and so forth.

How to run:

### Using local Go Installation

- install go 1.21.0 https://go.dev/doc/install
- `go mod tidy`
- `go run main.go serve --dev`

### The lazy docker way:

- Have docker installed
- `docker compose up`

and there you go, the docker is configured with air which will automatically reload the app if you make any code changes.

It will not reload if there are any migrations, so be sure to reload the container if you make any data changes.
