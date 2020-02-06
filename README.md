# go-gin-todo
![Go](https://github.com/danrusu/go-gin-todo/workflows/Go/badge.svg?branch=master&event=push)

Simple TODO REST API using GO GIN 

### 0. Setup
 - [Go](https://golang.org/dl/)
 - IDE - [VSCode](https://code.visualstudio.com/download) / [GoLand](https://www.jetbrains.com/go/download)
 - [Postman](https://www.postman.com/downloads/)

### 1. Build and start Todo REST API  

```bash
git clone https://github.com/danrusu/go-gin-todo.git
cd go-gin-todo
go build
./go-gin-todo.exe
```

### 2. Test todo API via Postman
 - base URL: http://localhost:3000/api/todo/
 - [test collection](GO_GIN_TODO.postman_collection.json)


### 3. API DOC
Todo REST Api Endpoints
Create a new todo: POST /todos
Create a new todo by posting its details to /todos. You receive back its details in JSON format.

*Request POST /todos*
```
Content-Type: application/json
{   
    "name": "Software testing course", 
    "description": "Java Selenium Module", 
    "project" : "SDA",
    "date":  "01/03/2020",
    "time": "09:00 AM",
    "priority": 1
}
```

*Response*
```
HTTP/1.0 200 OK
Content-Type: application/json
{
    "data": {
        "id": 0,
        "name": "Software testing course",
        "description": "Java Selenium Module",
        "project": "SDA",
        "date": "01/03/2020",
        "time": "09:00 AM",
        "priority": 1
    },
    "message": "Todo successfully created",
    "status": 200
}
```


