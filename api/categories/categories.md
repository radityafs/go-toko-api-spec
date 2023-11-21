### Get All Categories

    this endpoint is used to get all categories related to the user

#### Endpoint

`GET` `/api/v1/categories`

#### Headers

```http
Authorization : Bearer <token>
```

#### Response [200] - OK

```json
{
  "success": true, // Boolean
  "message": "Successfuly get all categories",
  "data": [
    {
      "id": 1,
      "name": "Category 1"
    }
    // Category Data
  ]
}
```

#### Response [401] - Unauthorized

```json
{
  "success": false, // Boolean
  "message": "Unauthorized"
}
```
