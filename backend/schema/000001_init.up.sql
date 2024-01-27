CREATE TABLE sub (
	channel_id INT PRIMARY KEY,
	price INT,
	description VARCHAR(255),
	user_id INT NOT NULL,
	link VARCHAR(255),
	image VARCHAR(255)
);

CREATE TABLE sub_users (
	channel_id INT PRIMARY KEY,
	user_id INT NOT NULL,
	endtime VARCHAR(255)
);