### User Profile

    user_profile table is used to store the user profile data.

#### Table name : `user_profile`

#### Table Column :

| Column     | Type         | Modifiers                    |
| ---------- | ------------ | ---------------------------- |
| id         | integer      | not null default primary key |
| user_id    | integer      | foreign key: `users.id`      |
| picture    | varchar(255) | nullable                     |
| phone      | varchar(255) | nullable                     |
| first_name | varchar(255) |                              |
| last_name  | varchar(255) |                              |
| nickname   | varchar(255) | nullable                     |
| gender     | enum         | enum: `male`, `female`       |
| birthdate  | date         | nullable                     |
| address    | varchar(255) | nullable                     |
| created_at | timestamp    | nullable                     |
| updated_at | timestamp    | nullable                     |
