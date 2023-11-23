### Products

    products table is used to store the product data.

#### Table name : `products`

#### Table Column :

| Column      | Type         | Modifiers                              |
| ----------- | ------------ | -------------------------------------- |
| id          | integer      | not null default primary key           |
| category_id | integer      | foreign key: `categories.id`, nullable |
| brand_id    | integer      | foreign key: `brands.id`, nullable     |
| shop_id     | integer      | foreign key: `shops.id`                |
| name        | varchar(255) |                                        |
| description | varchar(255) | nullable                               |
| sku         | varchar(255) | nullable                               |
| quantity    | integer      | not null default 0                     |
| price_buy   | integer      | not null default 0                     |
| price_sell  | integer      | not null default 0                     |
| images      | varchar(255) | nullable                               |
| isActive    | boolean      | not null default true                  |
| created_at  | timestamp    | nullable                               |
| updated_at  | timestamp    | nullable                               |
