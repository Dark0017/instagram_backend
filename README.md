# Instagram Backend API

- **GOLANG APIs** to handle `USER` and `POST` creation
- These APIs only use go standard packages
- MongoDB is used for storage

## Prerequisites

- MongoDB connection string / URI (Local or Remote)
- [GOLANG](https://golang.org/)

## Setup

- Extract the repo into a folder, I'll be calling this `Repo`
- Add this directory to `GOPATH` environment variable ([Follow this for help](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/01.2.html))
- Open the terminal of your choice and navigate to `Repo/src/insta_backend`
- Run the below snippet to download the dependencies

```go
 go mod download
```

- Make sure you have the following environment variables setup:

```go
MONGO_URI = <mongoDB connection string> // example: mongodb+srv://<user_name>:<password>@cluster0.spv7l.mongodb.net/Cluster0
MONGO_DBNAME = <name of the mongoDB database> //example: MONGO_DBNAME = "MyDB"
ENCRYPT_KEY = <32 BIT LONG ENCRYPTION KEY> //example: ENCRYPT_KEY = "thisis32bitlongpassphraseimusing"
PORT = <the port where the server will run>// example: PORT = :3000
```

- Run the following snippet to start the server at the `PORT`

```go
go run main.go
```

## Usage

There are 5 endpoints relating to `users` and `posts`

> **`GET /users/:id`**

This endpoint returns user object with the following structure:

    id       //mongoObjectId
    name     //string
    email    //string
    password //string

> **`GET /posts/:id`**

This endpoint returns post object with the following structure:

    id        //mongoObjectId
    userId    //mongoObjectId
    caption   //string
    imageUrl  //string
    postTime  //ISOSTRING

> **`GET /posts/users/:id`**

This endpoint returns an array of post objects with the following structures:

    id        //mongoObjectId
    userId    //mongoObjectId
    caption   //string
    imageUrl  //string
    postTime  //ISOSTRING

> **`POST /users/`**

This endpoint returns objectId to of the user created, the request body should contain:

    name     //string
    email    //string
    password //string

> **`POST /posts/`**

This endpoint returns objectId to of the post created, the request body should contain:

    userId    //mongoObjectId
    caption   //string
    imageUrl  //string
    postTime  //ISOSTRING

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
