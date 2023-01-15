# Backend Service for Derma Care Store

## Available Services
### User
- Signup(`/signup`) POST
- Signin(`/signin`) POST
### Product
- Query Products(`/products`) GET
- Query Product By Id(`/products/:id`) GET
- Create Product(`/products`) POST
- Update Product(`/products/:id`) PUT
- Delete Product(`/products/:id`) DELETE
### Category
- Query Categories(`/categories`) GET
- Create Category(`/categories`) POST
- Update Category(`/categories/:id`) PUT
- Delete Category(`/categories/:id`) DELETE
### Order
- Query Orders(`/orders`) GET
- Create Order(`/orders`) POST

## Local Development
This project use [air](https://github.com/cosmtrek/air) to live reload. If you are familiar with `nodemon` the concept is same but `air` is built with Go. 
Rename .air.example.toml to .air.toml and modify `DB_URI` and `DB_NAME` properly.
```
full_bin = "DB_URI=mongodb://localhost:PORT_NUMBER DB_NAME=YOUR_DB_NAME ./tmp/main"
```
Then run 
```
make generate-private-key && make generate-public-key
```
to generate `public` and `private` keys pair. After that, you can just run `air` in the root of your project directory.
## Running inside container
To run the application inside container you need to do these things.
1. copy `.env.example` inside `zarf/docker` to `.env` file.
2. create a docker network called `dermacare` by running this command `docker create network dermacare`.
3. since containers(one for `database` and one for `application`) are running within the same network, they can connect each other by their name. update the `DB_URI` inside `.env` to 
```
DB_URI=mongodb://db:PORT_NUMBER
DB_NAME=YOUR_DB_NAME
``` 
> Note `db` is come from `services` which declare inside `docker-compose.yaml`.
4. run `make start`. This command will build `dermacare-service:tag` image and will start the application.
5. server is running on port `3000`.
# TODO
* Support search(filter) in all query services.
* Indexing on email fields to prevent duplicate email.
* Add other services
  1. Analytic, Summary
  2. Excel Export
* Rotating Secret
* Embed Admin UI to Binary
