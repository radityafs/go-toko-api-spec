### Sales Detail

        Sales Detail table is used to store the sales data detail.

#### Table name : `sales_detail`

#### Table Column :

| Column     | Type      | Modifiers                    |
| ---------- | --------- | ---------------------------- |
| id         | integer   | not null default primary key |
| sales_id   | bigint    | foreign key: `sales.id`      |
| product_id | bigint    | foreign key: `products.id`   |
| name       | varchar   | not null                     |
| quantity   | integer   | not null                     |
| category   | varchar   | not null                     |
| price      | bigint    | not null                     |
| total      | bigint    | not null                     |
| created_at | timestamp |                              |
| updated_at | timestamp |                              |
