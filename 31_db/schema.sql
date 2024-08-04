--CREATE DATABASE tasks;

DROP TABLE IF EXISTS tasks_labels, labels, tasks, users;

CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE labels(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE tasks(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	opened BIGINT NOT NULL DEFAULT extract(epoch from now()),
	closed BIGINT DEFAULT 0,
	title TEXT,
	content TEXT,
	author_id INTEGER REFERENCES users(id) NOT NULL DEFAULT 0,
	assigned_id INTEGER REFERENCES users(id) NOT NULL DEFAULT 0
);

CREATE TABLE tasks_labels(
	task_id INTEGER REFERENCES tasks(id) NOT NULL DEFAULT 0,
	label_id INTEGER REFERENCES labels(id) NOT NULL DEFAULT 0
);
