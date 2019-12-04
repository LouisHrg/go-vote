# Awesome Golang API [Project](https://github.com/ritoon/cours/blob/esgi/project/README.md)

[Try the demo version](https://awesome-go-vote.herokuapp.com/)

Developers :
- [MATI Audric](https://github.com/odrik)
- [VARVEROPOULOS Valentin](https://github.com/firenaik)
- [HARANG Louis](https://github.com/louishrg)

## Requirements

go 1.13

## Run the project

Install go vendor :

```make install-go-vendor```

Install all the dependencies :

```make install-dep```

Build and run the project :

```make run```

Simply start a builded version of the project :

```make start```

You can run static analysis of the base code with :

```make analysis```

## Dependencies

-  [gabs](github.com/Jeffail/gabs)
-  [gin-jwt](github.com/appleboy/gin-jwt/v2)
-  [govalidator](github.com/asaskevich/govalidator)
-  [godotenv](github.com/joho/godotenv)
-  [pagination](github.com/biezhi/gorm-paginator/pagination)
-  [gin](github.com/gin-gonic/gin)
-  [gorm](github.com/jinzhu/gorm)
-  [go.uuid](github.com/satori/go.uuid)
-  [crypto](golang.org/x/crypto)

## Documentation

A postman collection is available in the doc/ folder

### POST /login
    {
         "email": "admin@admin.com",
         "password": "admin"
    }

#### Response
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2wiOiIxIiwidGltZSI6MTU2OTg0Njg2NSwidXVpZCI6Ijc2ZDFhZjYzLTU3MWEtNDNkNC04NzNjLTYwOGY1MDc4YmY5NyJ9.p8IPR8sl2jsNr4yU6AvScaoRWhtCkRQ0K4zz2moPGUE"
    }

---

> For following, set header to :

    {
        "authorization": "Bearer: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2wiOiIxIiwidGltZSI6MTU2OTg0Njg2NSwidXVpZCI6Ijc2ZDFhZjYzLTU3MWEtNDNkNC04NzNjLTYwOGY1MDc4YmY5NyJ9.p8IPR8sl2jsNr4yU6AvScaoRWhtCkRQ0K4zz2moPGUE"
    }

---

### GET /users

#### Response
    {
        "total_record": 2,
        "total_page": 1,
        "records": [
            {
                "uuid": "76d1af63-571a-43d4-873c-608f5078bf97",
                "createdAt": "2019-09-30T14:34:15.002587+02:00",
                "updatedAt": "2019-09-30T14:34:15.002587+02:00",
                "email": "admin@admin.com",
                "firstname": "Jean",
                "lastname": "Admin",
                "access_level": 1,
                "date_of_birth": "1996-02-08T00:00:00Z"
            },
            {
                "uuid": "8a8bd7e7-4fd3-4228-b257-24049fd227d9",
                "createdAt": "2019-09-30T14:34:16.211285+02:00",
                "updatedAt": "2019-09-30T14:34:16.211285+02:00",
                "email": "test@test.com",
                "firstname": "Test",
                "lastname": "Test",
                "access_level": 2,
                "date_of_birth": "1996-02-08T00:00:00Z"
            }
        ],
        "offset": 0,
        "limit": 10,
        "page": 1,
        "prev_page": 1,
        "next_page": 1
    }

---

### GET /users/:udid

#### Response 
    {
        "code": 200,
        "data": {
            "uuid": "76d1af63-571a-43d4-873c-608f5078bf97",
            "createdAt": "2019-09-30T14:34:15.002587+02:00",
            "updatedAt": "2019-09-30T14:34:15.002587+02:00",
            "email": "admin@admin.com",
            "firstname": "Jean",
            "lastname": "Admin",
            "access_level": 1,
            "date_of_birth": "1996-02-08T00:00:00Z"
        }
    }

---

### POST /users
    {
        "email": "newuser@esgi.fr",
        "firstname": "New",
        "lastname": "User",
        "access_level": 1,
        "date_of_birth": "1998-01-13T00:00:00Z"
    }

#### Response
    {
        "code": 201,
        "data": {
            "uuid": "8fc8ca6f-2ce9-4eda-a2f8-8f2e51b0caa1",
            "createdAt": "2019-09-30T14:42:45.060986+02:00",
            "updatedAt": "2019-09-30T14:42:45.060986+02:00",
            "email": "newuser@esgi.fr",
            "firstname": "New",
            "lastname": "User",
            "access_level": 1,
            "date_of_birth": "1998-01-13T00:00:00Z"
        },
        "message": "Ressource created"
    }

---

### DELETE /users/:udid

#### Response
    {
        "code": 200,
        "data": {
            "uuid": "",
            "createdAt": "0001-01-01T00:00:00Z",
            "updatedAt": "0001-01-01T00:00:00Z",
            "email": "",
            "firstname": "",
            "lastname": "",
            "access_level": 0,
            "date_of_birth": "0001-01-01T00:00:00Z"
        },
        "message": "Ressource deleted"
    }

---

### PUT /users/:udid
> WIP

---

### GET /surveys

#### Response
    {
        "total_record": 1,
        "total_page": 1,
        "records": [
            {
                "uuid": "134d9b04-e0c9-4c48-83b8-ece3a29a24b2",
                "createdAt": "2019-09-30T14:34:13.814135+02:00",
                "updatedAt": "2019-09-30T14:34:13.817821+02:00",
                "title": "Propreté des trottoirs",
                "desc": "Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires.",
                "startDate": "2018-10-01T00:00:00Z",
                "endDate": "2020-10-01T00:00:00Z",
                "Responses": [
                    {
                        "uuid": "6d2fcc4b-c06a-4a9e-be1e-460884479a79",
                        "createdAt": "2019-09-30T14:34:13.816395+02:00",
                        "updatedAt": "2019-09-30T14:34:13.816395+02:00",
                        "message": "Oui je suis d'accord",
                        "Users": []
                    },
                    {
                        "uuid": "331bc8aa-03e4-4908-8894-640600c90f0d",
                        "createdAt": "2019-09-30T14:34:13.819248+02:00",
                        "updatedAt": "2019-09-30T14:34:13.819248+02:00",
                        "message": "Non je ne suis pas d'accord",
                        "Users": []
                    }
                ]
            }
        ],
        "offset": 0,
        "limit": 10,
        "page": 1,
        "prev_page": 1,
        "next_page": 1
    }

---

### GET /surveys/:udid

#### Response
    {
        "code": 200,
        "data": {
            "uuid": "134d9b04-e0c9-4c48-83b8-ece3a29a24b2",
            "createdAt": "2019-09-30T14:34:13.814135+02:00",
            "updatedAt": "2019-09-30T14:34:13.817821+02:00",
            "title": "Propreté des trottoirs",
            "desc": "Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires.",
            "startDate": "2018-10-01T00:00:00Z",
            "endDate": "2020-10-01T00:00:00Z",
            "Responses": [
                {
                    "uuid": "6d2fcc4b-c06a-4a9e-be1e-460884479a79",
                    "createdAt": "2019-09-30T14:34:13.816395+02:00",
                    "updatedAt": "2019-09-30T14:34:13.816395+02:00",
                    "message": "Oui je suis d'accord",
                    "Users": []
                },
                {
                    "uuid": "331bc8aa-03e4-4908-8894-640600c90f0d",
                    "createdAt": "2019-09-30T14:34:13.819248+02:00",
                    "updatedAt": "2019-09-30T14:34:13.819248+02:00",
                    "message": "Non je ne suis pas d'accord",
                    "Users": []
                }
            ]
        }
    }

---

### POST /surveys
    {
        "title": "Propreté des trottoirs",
        "desc": "Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires.",
        "startDate": "2018-10-01T00:00:00Z",
        "endDate": "2020-10-01T00:00:00Z"
    }

#### Response
    {
        "code": 201,
        "data": {
            "uuid": "65627472-ef4a-4e7d-ad73-2cc00d2a0353",
            "createdAt": "2019-09-30T14:46:07.822161+02:00",
            "updatedAt": "2019-09-30T14:46:07.822161+02:00",
            "title": "Propreté des trottoirs",
            "desc": "Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires.",
            "startDate": "2018-10-01T00:00:00Z",
            "endDate": "2020-10-01T00:00:00Z",
            "Responses": null
        },
        "message": "Ressource created"
    }

---

### PUT /surveys/:udid
> WIP

---

### DELETE /responses/:uuid
> WIP

---

### POST /responses/:uuid
> WIP
