
# Library System Api

This project is a backend api for a library app. created using Golang's restfull api without a framework.



## Features

- User CRUD (with level)
- Book CRUD
- Loan CRUD
- Penalties CRUD


## Installation

Clone this repo

```bash
  git clone https://github.com/Wrendra57/library-api.git
```


Create Mysql database with name `library`.

Running file `query.sql` in mysql.





## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DATABASE_URL`
`CLOUDINARY_URL`
`SECRET_KEY`

copy file `.env.example` to `.env` set with your configuration

## Run Locally

Go to the project directory

```bash
  cd library-api
```

Start the server

```bash
  go run .
```

## Documentation API

- Import file `Documentation_postmant.json` ke Postman App
- 

## User Seed
- admin

```bash
  email: admin@gmail.com
  password: 1234
``` 

- member
 ```bash
  email: member@gmail.com
  password: 1234
``` 

