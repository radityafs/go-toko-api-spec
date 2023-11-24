### Get All Products

    this endpoint is used to get all products which stock is greater than 0 related to the user

#### Endpoint

`GET` `/api/v1/products`

#### Headers

```http
Authorization : Bearer <token>
```

#### Parameter/Query

    this parameter is optional and can be used to filter the products

- default value for page is 1
- default value for limit is 10

```http
page : <number>
limit : <number>
query : <string>
id_category : <number>
sort : (optional)
  - name
  - order
  - price
  - created_at
  - updated_at
order :
  - asc
  - desc
```

#### Response [200] - OK

```json
{
  "success": true, // Boolean
  "message": "Successfuly get all products",
  "data": [
    {
      "id": 1,
      "name": "Product 1",
      "price": 10000,
      "stock": 10,
      "photo": "http://localhost:3000/uploads/products/1.jpg",
      "category": {
        "id": 1,
        "name": "Category 1"
      },
      "created_at": "2021-09-01T00:00:00.000Z",
      "updated_at": "2021-09-01T00:00:00.000Z"
    }
  ],
  "meta": {
    "page": 1,
    "limit": 10,
    "totalData": 100,
    "totalPage": 10
  }
}
```

#### Response [401] - Unauthorized

```json
{
  "success": false, // Boolean
  "message": "Invalid token or token expired"
}
```

#### Response [400] - Bad Request

```json
{
  "success": false, // Boolean
  "message": "Invalid query parameter"
}
```
