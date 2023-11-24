### Get Transaction Detail

    this endpoint is used to get transaction detail

#### Endpoint

`GET` `/api/v1/transactions/:id`

#### Headers

```http
Authorization : Bearer <token>
```

#### Response [200] - OK

```json
{
  "success": true, // Boolean
  "message": "Successfuly get transaction detail",
  "data": {
    "id": 1,
    "invoice": "INV-20210901-0001", // [INV-<YYYYMMDD>-<number>]
    "payment_method": "CASH", // [CASH, QRIS]
    "status": "PAID", // [PAID, UNPAID, VOID]
    "total": 100000,
    "total_payment": 100000,
    "change": 0,
    "products": [
      {
        "id": 1,
        "name": "Product 1",
        "price": 10000,
        "qty": 10,
        "subtotal": 100000
      }
      // Product Data
    ],
    "created_at": "2021-09-01T00:00:00.000Z",
    "updated_at": "2021-09-01T00:00:00.000Z"
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

#### Response [404] - Not Found

```json
{
  "success": false, // Boolean
  "message": "Transaction not found"
}
```
