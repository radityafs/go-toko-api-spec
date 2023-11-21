### Refresh Token

    this endpoint is used to refresh the token to the application if the token is expired

#### Endpoint

`POST` `/api/v1/auth/refresh-token`

#### Request

```json
{
  "refreshToken": "string"
}
```

#### Response [200] - OK

```json
{
  "success": true, // Boolean
  "message": "Successfuly generate new token",
  "token": "ey***"
}
```

#### Response [400] - Bad Request

```json
{
  "success": false, // Boolean
  "message": "Refresh token is required or invalid"
}
```
