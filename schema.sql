-- Create the database if it doesn't exist
CREATE DATABASE IF NOT EXISTS todo_db;

-- Use the todo_db database
USE todo_db;

-- Create the todos table
CREATE TABLE IF NOT EXISTS todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE
);
