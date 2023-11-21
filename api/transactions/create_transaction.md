### Create Transaction

    this endpoint is used to create transaction

#### Endpoint

`POST` `/api/v1/transactions`

#### Headers

```http
Authorization : Bearer <token>
```

#### Body

```json
{
  "payment_method": "CASH", // [CASH, QRIS]
  "products": [
    {
      "id": 1,
      "qty": 10
    }
    // Product Data
  ],
  "total_payment": 100000,
  "change": 0
}
```

#### Response [201] - Created [CASH]

```json
{
  "success": true, // Boolean
  "message": "Successfuly create transaction",
  "data": {
    "id": 1,
    "invoice": "INV-20210901-0001",
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
    ],
    "created_at": "2021-09-01T00:00:00.000Z",
    "updated_at": "2021-09-01T00:00:00.000Z"
  }
}
```

#### Response [201] - Created [QRIS]

```json
{
  "success": true, // Boolean
  "message": "Successfuly create transaction",
  "data": {
    "id": 1,
    "invoice": "INV-20210901-0001",
    "payment_method": "QRIS", // [CASH, QRIS]
    "qris_id": "qris-1234567890",
    "status": "UNPAID", // [PAID, UNPAID, VOID]
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
    ],
    "created_at": "2021-09-01T00:00:00.000Z",
    "updated_at": "2021-09-01T00:00:00.000Z"
  }
}
```

#### Response [400] - Bad Request

```json
{
  "success": false, // Boolean
  "message": "Total payment must be greater than total"
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
  "message": "Product not found"
}
```

#### Response [500] - Internal Server Error

```json
{
  "success": false, // Boolean
  "message": "Internal Server Error"
}
```
