CREATE TABLE sub (
	channel_id INT IDENTITY (1,1) NOT NULL,
	price INT,
	description VARCHAR(255),
	user_id INT NOT NULL,
	link VARCHAR(255),
	image VARCHAR(255),

);
