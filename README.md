# Simple API in Go with Fiber and Docker

This repository is part of my article on Medium:  
[Build a Simple API in Go with Fiber and Docker](https://levelup.gitconnected.com/api-in-go-with-fiber-and-docker-5de04651463a)

## Docker
Clone this repository and run:
```
docker-compose up
```

You can then hit the following endpoints:

| Method | Route         | Body                                           |
| ------ | ------------- | ---------------------------------------------- |
| GET    | /products     |                                                |
| GET    | /products/:id |                                                |
| POST   | /products     | `{"name": "iPhone", "stock": 10, "price": 99 }`|
| DELETE | /products/:id |                                                |
| PUT    | /products/:id | `{"name": "iPhone", "stock": 10, "price": 99 }`|

