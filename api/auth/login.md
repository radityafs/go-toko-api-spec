### Login

    this endpoint is used to login to the application

#### Endpoint

`POST` `/api/v1/auth/login`

#### Request

```json
{
  "email": "string",
  "password": "string"
}
```

#### Response [200] - OK

```json
{
  "success": true, // Boolean
  "message": "Successfuly logged in",
  "data": {
    // User Data
  },
  "token": "ey***",
  "refreshToken": "ey***"
}
```

#### Response [400] - Bad Request

```json
{
  "success": false, // Boolean
  "message": "Email or password is required"
}
```

#### Response [401] - Unauthorized

```json
{
  "success": false, // Boolean
  "message": "Invalid email or password"
}
```
