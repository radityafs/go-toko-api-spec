### Sales

        Sales table is used to store the sales data.

#### Table name : `sales`

#### Table Column :

| Column       | Type      | Modifiers                      |
| ------------ | --------- | ------------------------------ |
| id           | integer   | not null default primary key   |
| order_id     | varchar   | not null unique                |
| shop_id      | bigint    | foreign key: `shops.id`        |
| cashier_id   | bigint    | foreign key: `users.id`        |
| status       | enum      | enum: `unpaid`, `paid`, `void` |
| payment_type | enum      | enum: `cash`, `qris`           |
| total_bill   | bigint    |                                |
| total_paid   | bigint    |                                |
| total_item   | integer   |                                |
| created_at   | timestamp |                                |
| updated_at   | timestamp |                                |
