> **Task-api**: This API was designed in the intention of handling list of tasks in and out our database, allowing users to create a task and saving into the database 
> of their choice (for now: either in psql or mysql), get all tasks avaliable in the database, get a task using the ID, update contents of a task (whether the
> title, description or both) and delete task from a database by ID.

**Prerequisites** <br>
Before using the API, make sure to either have either mysql or postgres (psql) in your operating system installed. Also make sure to run the schema below either in mysql
or in psql (recommend using shell or bash). <br>
**Schema**: <br>
*For MySQL* - <br>
```sql
CREATE DATABASE IF NOT EXISTS taskdb;
USE taskdb;

CREATE TABLE IF NOT EXISTS tasks (
id CHAR(36) PRIMARY KEY,
title VARCHAR(150) NOT NULL,
description VARCHAR(300) NOT NULL
);
```
*For psql in bash* - 
```sql
CREATE DATABASE IF NOT EXISTS task_api;
\c task_api

CREATE TABLE IF NOT EXISTS tasks (
id CHAR(36) PRIMARY KEY,
title VARCHAR(150) NOT NULL,
description VARCHAR(300) NOT NULL
);
```
For testing, you can add these samples to your database: <br>
```sql
INSERT INTO tasks (id,title,description) VALUES ('1161377b-7c47-45aa-b087-1a9b08423021','test1','testing123');
INSERT INTO tasks (id,title,description) VALUES ('18fd5e4b-d137-477f-a499-06ef11859292','test2','testing456');
INSERT INTO tasks (id,title,description) VALUES ('3ce5ceec-6c5a-11f0-9dac-68ecc56cca7d','test3','testing789');
```
