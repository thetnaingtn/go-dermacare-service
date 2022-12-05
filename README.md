# Backend service for derma care store

The purpose of this system is to allow the owner to manage stock and order.

## Local Development
This project use [air](https://github.com/cosmtrek/air) to live reload. If you are familiar with `nodemon` the concept is same but `air` is built with Go. 
Rename .air.example.toml to .air.toml and modify `DB_URI` and `DB_NAME` properly.
```
full_bin = "DB_URI=YOUR_DB_URI DB_NAME=YOUR_DB_NAME ./tmp/main"
```
