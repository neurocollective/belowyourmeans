CREATE TABLE app_user (
	id serial PRIMARY KEY,
	username text,
	hashed_password text
);

CREATE TABLE user_settings (
	id serial PRIMARY KEY,
	use_auto_login boolean,
	update_frequency int
);

CREATE TABLE budget_category (
	id serial PRIMARY KEY,
	user_id int REFERENCES app_user(id),
	name text NOT NULL,
	description text
);

CREATE TABLE budget_allocation (
	id serial PRIMARY KEY,
	user_id int REFERENCES app_user(id),
	category_id REFERENCES budget_category(id),
	amount real NOT NULL,
);

CREATE TABLE exepnditure (
	id serial PRIMARY KEY,
	category_id REFERENCES budget_category(id),
	month int NOT NULL,
	year int NOT NULL,
);
