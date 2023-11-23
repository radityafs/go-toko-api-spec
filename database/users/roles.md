### Roles

    roles table is used to store the roles data.

#### Table name : `roles`

#### Table Column :

| Column     | Type         | Modifiers                    |
| ---------- | ------------ | ---------------------------- |
| id         | integer      | not null default primary key |
| name       | varchar(255) | not null unique              |
| created_at | timestamp    |                              |
| updated_at | timestamp    |                              |

Description:

1. `id` is the primary key of the table.
   1. 1 is Super Admin => can do anything
   2. 2 is Owner => can manage shop
   3. 3 is Cashier => can manage sales
