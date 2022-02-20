# API DOCS

# Base URL
http://127.0.0.1:8080/api/

# Other resources 

 
# Headers

Authorization: key your token

Accept : application/json


Locale : ar|en  

# API 

| Route                        | Request Method | Parameters | Response  |
| -----------                  | -----------    |----------- |---------- |
| seller/register            | POST           |  [Register Parameters](#Register)|[Response](#Response)|
| seller/login | POST           |[Login Parameters](#Login)|  [Response](#Response)         |
|buyer/register         | POST           |  [Login Parameters](#Login) |[Response](#Response)         |
|buyer/login      |POST           |  [Register Parameters](#Register)|[Response](#Response)     |
|buy              |POST           |  [Buy Parameters](#Buy)|[Response](#Response)     |





# <a name="Buy"> </a> Buy 

```json
{
    "product_id": "int",
    "amount" : "int"
} 
```



# <a name="Register"> </a> Register new User 

```json
{
    "username" : "String",
    "password" : "String",
    "deposit"  : "Float"
} 
```


# <a name="Login"> </a> Login User 

```json
{
    "username" : "String",
    "password" : "String",
} 
```

# <a name="Response"> </a> Responses 

## Unauthorized error

__*Response code : 401*__
```json 
{
    "message" : "Unauthenticated"
}
```

## Validation error 
__*Response code : 422*__

```json 
{
    "errors" {
        "Key" : "Error message"
    }
}
```
## Success  
__*Response code : 201*__
```json
{
    "code": 201,
    "errors": null,
    "message": "created",
    "payload": {
    }
}

```
