# Awesome Golang API [Project](https://github.com/ritoon/cours/blob/esgi/project/README.md)

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
-  [pagination](github.com/biezhi/gorm-paginator/pagination)
-  [gin](github.com/gin-gonic/gin)
-  [gorm](github.com/jinzhu/gorm)
-  [go.uuid](github.com/satori/go.uuid)
-  [crypto](golang.org/x/crypto)

## Documentation

### GET /users

### GET /users/:udid

### POST /users
    {
        "email": "newuser@esgi.fr",
        "firstname": "New",
        "lastname": "User",
        "accesslevel": 1,
        "dateofbirth": "1998-01-13T00:00:00Z"
    }

### DELETE /users/:udid

### PUT /users/:udid
    {
        "email": "newuser@esgi.fr",
        "firstname": "New",
        "lastname": "User",
        "accesslevel": 1,
        "dateofbirth": "1998-01-13T00:00:00Z"
    }
