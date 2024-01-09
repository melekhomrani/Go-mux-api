# Movie API

This is a simple movie API implemented in Go using the Gorilla Mux router. The API allows you to perform basic CRUD operations on a collection of movies, each associated with a director.

## Getting Started

To get started with this project, follow the instructions below.

### Prerequisites

Make sure you have Go installed on your machine. If not, you can download it from [here](https://golang.org/dl/).

### Installing

1. Clone the repository:

```bash
  git clone https://github.com/melekhomrani/Go-mux-api.git
  cd Go-mux-api
```

1. Build and run the project:

```bash
  go build
  ./Go-mux-api.exe
```

The server will start at localhost:8080.

# API Endpoints

## GET /movies

Returns a list of all movies from a static list.

```http
  GET /movies
```	

### Response

```json
[
  {
    "id": "2",
    "title": "The Godfather",
    "director": {
      "firstname": "Francis",
      "lastname": "Ford Coppola"
    }
  },
  {
    "id": "3",
    "title": "The Godfather: Part II",
    "director": {
      "firstname": "Francis",
      "lastname": "Ford Coppola"
    }
  },
  {
    "id": "4",
    "title": "The Dark Knight",
    "director": {
      "firstname": "Christopher",
      "lastname": "Nolan"
    }
  }
]
```
## GET /movies/{id}

Returns a movie with the specified id.

```http
  GET /movies/2
```

### Response

```json
{
  "id": "2",
  "title": "The Godfather",
  "director": {
    "firstname": "Francis",
    "lastname": "Ford Coppola"
  }
}
```

## POST /movies

Creates a new movie.

```http
  POST /movies
```

### Request Body

```json
{
  "title": "The Godfather: Part III",
  "director": {
    "firstname": "Francis",
    "lastname": "Ford Coppola"
  }
}
```

### Response

```json
{
  "id": "5",
  "title": "The Godfather: Part III",
  "director": {
    "firstname": "Francis",
    "lastname": "Ford Coppola"
  }
}
```

## PUT /movies/{id}

Updates a movie with the specified id.

```http
  PUT /movies/5
```

### Request Body

```json
{
  "title": "The Godfather: Part III",
  "director": {
    "firstname": "Francis",
    "lastname": "Coppola"
  }
}
```

### Response

```json
{
  "id": "5",
  "title": "The Godfather: Part III",
  "director": {
    "firstname": "Francis",
    "lastname": "Coppola"
  }
}
```

## DELETE /movies/{id}

Deletes a movie with the specified id.

```http
  DELETE /movies/5
```

### Response

```json
{
  "id": "5",
  "title": "The Godfather: Part III",
  "director": {
    "firstname": "Francis",
    "lastname": "Coppola"
  }
}
```

## GET /movies/{id}/director

Returns the director of the movie with the specified id.

```http
  GET /movies/2/director
```

### Response

```json
{
  "firstname": "Francis",
  "lastname": "Ford Coppola"
}
```

# Built With

- [Go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

- [Gorilla Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for golang.

# Authors

- **Melek Homrani** - [melekhomrani](https://github.com/melekhomrani)

# License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details