### Get Top Products or Categories

    this endpoint is used to get top products or categories based on transaction

#### Endpoint

`GET` `/api/v1/transactions/top`

#### Headers

```http
Authorization : Bearer <token>
```

#### Parameter/Query

    this parameter is optional and can be used to filter the top products or categories

- default value for page is 1
- default value for limit is 10

```http
page : <number>
limit : <number>
type : <string> [product, category]
start_date : <date> [DD-MM-YYYY]
end_date : <date> [DD-MM-YYYY]
```

#### Response [200] - OK

```json
{
  "success": true, // Boolean
  "message": "Successfuly get top products",
  "data": [
    {
      "id": 1,
      "name": "Product 1",
      "sold": 10
    }
    // Product Data
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
  "message": "Unauthorized"
}
```

#### Response [400] - Bad Request

```json
{
  "success": false, // Boolean
  "message": "Invalid query parameter"
}
```
