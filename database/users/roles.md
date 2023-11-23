### Roles

    roles table is used to store the roles data.

#### Table: `roles`

| Column     | Type         | Modifiers                    |
| ---------- | ------------ | ---------------------------- |
| id         | integer      | not null default primary key |
| name       | varchar(255) | not null unique              |
| created_at | timestamp    |                              |
| updated_at | timestamp    |                              |
