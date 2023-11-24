### Users

    user table is used to store the user data.

#### Table name : `users`

#### Table Column :

| Column            | Type         | Modifiers                    |
| ----------------- | ------------ | ---------------------------- |
| id                | integer      | not null default primary key |
| email             | varchar(255) | not null unique              |
| email_verified_at | timestamp    | nullable                     |
| password          | varchar(255) | not null                     |
| last_seen         | timestamp    | nullable                     |
| is_subscribe      | boolean      | not null default false       |
| remember_token    | varchar(100) | nullable                     |
| role_id           | integer      | foreign key: `roles.id`      |
| google_id         | varchar(255) | nullable                     |
| created_at        | timestamp    | nullable                     |
| updated_at        | timestamp    | nullable                     |
