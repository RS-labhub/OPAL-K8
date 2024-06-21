CREATE DATABASE recipes_db;

USE recipes_db;

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  karma INT DEFAULT 0,
  location VARCHAR(255) NOT NULL
);

CREATE TABLE recipes (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  required_karma INT DEFAULT 0,
  location VARCHAR(255)
);

INSERT INTO users (username, karma, location) VALUES ('rohan', 100, 'NY');
INSERT INTO recipes (title, description, required_karma, location) VALUES ('Recipe 1', 'Description 1', 50, 'NY');