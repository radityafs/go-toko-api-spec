### Get All Transactions

    this endpoint is used to get all transactions related to the user

#### Endpoint

`GET` `/api/v1/transactions`

#### Headers

```http
Authorization : Bearer <token>
```

#### Parameter/Query

    this parameter is optional and can be used to filter the transactions

- default value for page is 1
- default value for limit is 10

```http
page : <number>
limit : <number>
start_date : <date> [DD-MM-YYYY]
end_date : <date> [DD-MM-YYYY]
payment_method : <string> [CASH, QRIS]
```

#### Response [200] - OK

```json
{
  "success": true, // Boolean
  "message": "Successfuly get all transactions",
  "data": [
    {
      "id": 1,
      "invoice": "INV-20210901-0001", // [INV-<YYYYMMDD>-<number>]
      "payment_method": "CASH", // [CASH, QRIS]
      "status": "PAID", // [PAID, UNPAID, VOID]
      "total": 100000,
      "created_at": "2021-09-01T00:00:00.000Z",
      "updated_at": "2021-09-01T00:00:00.000Z"
    }
    // Transaction Data
  ],
  "summary": {
    "total_transaction": 100, // only for PAID transaction
    "total_income": 1000000 // only for PAID transaction
  },
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

#### Response [500] - Internal Server Error

```json
{
  "success": false, // Boolean
  "message": "Internal Server Error"
}
```
