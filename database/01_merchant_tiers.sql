CREATE TABLE IF NOT EXISTS tasks
(
    id SMALLSERIAL PRIMARY KEY, 
    task_name             VARCHAR(255) NOT NULL,
    task_owner             VARCHAR(255) NOT NULL,
    description             VARCHAR(255) NOT NULL,
    status             BOOLEAN
);