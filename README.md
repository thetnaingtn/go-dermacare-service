# Backend service for derma care store

## Local Development
This project use [air](https://github.com/cosmtrek/air) to live reload. If you are familiar with `nodemon` the concept is same but `air` is built with Go. 
Rename .air.example.toml to .air.toml and modify `DB_URI` and `DB_NAME` properly.
```
full_bin = "DB_URI=YOUR_DB_URI DB_NAME=YOUR_DB_NAME ./tmp/main"
```
Then run 
```
make generate-private-key && make generate-public-key
```
to generate `public` and `private` keys pair. After that, you can just run `air` in the root of your project directory.
## Running inside container
To run the application inside container you need to do these things.
1. copy .env.example inside `zarf/docker` to .env and populate `DB_URI` `DB_NAME` with proper values.
2. create a docker network called `dermacare` by running this command `docker create network dermacare`
3. run `make start`. This command will build `dermacare-service:tag` image and will start the application.
4. server is running on port `3000`.
