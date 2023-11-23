### Sales

    This table is used to store the payment from 3rd party payment gateway.

#### Table `sales_payment`

| Column       | Type          | Modifiers                      |
| ------------ | ------------- | ------------------------------ |
| id           | integer       | not null default primary key   |
| sales_id     | bigint        | foreign key: `sales.id`        |
| payment_ref  | varchar (255) | not null unique                |
| payment_type | varchar (255) | not null                       |
| amount       | bigint        | not null                       |
| status       | enum          | enum: `unpaid`, `paid`, `void` |
| created_at   | timestamp     |                                |
| updated_at   | timestamp     |                                |

Description:

1. payment_ref : is the reference or the id from the 3rd party payment gateway.
2. payment_type : is the type of payment, for example: `qris`, `gopay`, `ovo`, `dana`, `linkaja`, `bca`, `etc`.
3. amount : is the amount of the payment.
4. status : is the status of the payment, for example: `unpaid`, `paid`, `void`.
