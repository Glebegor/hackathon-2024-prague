CREATE TABLE sub (
	channel_id INT PRIMARY KEY,
	name VARCHAR(255),
	price INT,
	description VARCHAR(255),
	user_id INT NOT NULL,
	link VARCHAR(255),
	images VARCHAR(255),
	tags VARCHAR(255)
);

CREATE TABLE sub_users (
	channel_id INT PRIMARY KEY,
	user_id INT NOT NULL,
	endtime VARCHAR(255)
);