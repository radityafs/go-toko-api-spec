### Product Brand

    product brand table is used to store the product brand data.

#### Table `product_brand`

| Column      | Type          | Modifiers                    |
| ----------- | ------------- | ---------------------------- |
| id          | integer       | not null default primary key |
| shop_id     | integer       | foreign key: `shops.id`      |
| name        | varchar (255) | not null unique              |
| description | varchar (255) | nullable                     |
| isActive    | boolean       | not null default true        |
| isParent    | boolean       | not null default false       |
| created_at  | timestamp     | nullable                     |
| updated_at  | timestamp     | nullable                     |
