### Product Category

    product category table is used to store the product category data.

#### Table name : `product_category`

#### Table Column :

| Column      | Type          | Modifiers                    |
| ----------- | ------------- | ---------------------------- |
| id          | integer       | not null default primary key |
| shop_id     | integer       | foreign key: `shops.id`      |
| name        | varchar (255) | not null                     |
| code        | varchar (255) | not null                     |
| description | varchar (255) | nullable                     |
| images      | varchar (255) | nullable                     |
| isActive    | boolean       | not null default true        |
| isParent    | boolean       | not null default false       |
| created_at  | timestamp     | nullable                     |
| updated_at  | timestamp     | nullable                     |
