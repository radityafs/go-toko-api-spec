### Shops

    shops table is used to store the shop data.

#### table name : `shops`

#### table column :

| Column      | Type          | Modifiers                    |
| ----------- | ------------- | ---------------------------- |
| id          | integer       | not null default primary key |
| name        | varchar (255) | not null unique              |
| description | varchar (255) | nullable                     |
| user_id     | bigint        | foreign key: `users.id`      |
| province    | varchar (255) | nullable                     |
| regency     | varchar (255) | nullable                     |
| district    | varchar (255) | nullable                     |
| village     | varchar (255) | nullable                     |
| address     | varchar (255) |                              |
| is_active   | boolean       | not null default true        |
| created_at  | timestamp     |                              |
| updated_at  | timestamp     |                              |
