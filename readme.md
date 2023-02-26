# 🔖 Go todo app
This is a todo application just for learning purpose.

# 🚀 Getting Started
## Backend
### Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- Copy `.env.example` to `.env` and change values accordingly
- Change to backend directory

```shell
$ cd backend
```

#### Without Docker compose
- Make sure your mysql database is started

- Run below command
```shell
$ make run
```

#### With Docker compose
- Run below command
```shell
$ docker-compose up -d
```

## Frontend
- Change to frontend directory
```shell
$ cd frontend
```
- Open [index.html](frontend/index.html) in your favorite browser

### Tech stack
- Go for backend server
- MySQL for storing todos
- Docker for containerisation
- Frontend by [Max Sandelin](https://instagram.com/themaxsandelin)


### Self note (TODO)
1. Use ansible for configuration (If needed)
2. Terraform for infrastructure provisioning
3. Deploy on AWS
4. Automate complete workflow
