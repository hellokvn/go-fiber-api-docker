# Simple API in Go with Fiber and Docker

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

