CREATE TABLE sub (
	channel_id INT PRIMARY KEY,
	name VARCHAR(255),
	price INT,
	description VARCHAR(255),
	user_id INT NOT NULL,
	link VARCHAR(255),
	images VARCHAR(255),
	tags VARCHAR(255),
	wallet VARCHAR(1000)
);

CREATE TABLE sub_users (
	channel_id INT NOT NULL,
	user_id INT NOT NULL,
	endtime VARCHAR(255) NOT NULL
);