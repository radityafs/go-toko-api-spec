### Sales

    This table is used to store the payment from 3rd party payment gateway.

#### Table name : `sales_payment`

#### Table Column:

| Column       | Type          | Modifiers                      |
| ------------ | ------------- | ------------------------------ |
| id           | integer       | not null default primary key   |
| sales_id     | bigint        | foreign key: `sales.id`        |
| payment_ref  | varchar (255) | not null unique                |
| payment_type | enum          | enum: `qris`                   |
| amount       | bigint        | not null                       |
| status       | enum          | enum: `unpaid`, `paid`, `void` |
| created_at   | timestamp     |                                |
| updated_at   | timestamp     |                                |

Description:

1. payment_ref : is the reference or the id from the 3rd party payment gateway.
2. payment_type : is the type of payment from the 3rd party payment gateway, for example: `qris`.
3. amount : is the amount of the payment.
4. status : is the status of the payment, for example: `unpaid`, `paid`, `void`.
