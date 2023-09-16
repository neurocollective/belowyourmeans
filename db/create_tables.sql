CREATE TABLE budget_user (
	id serial PRIMARY KEY,
	first_name text, 
	last_name text,
	email text,
	hashed_password text,
	create_date timestamp,
	last_modified_date timestamp
);

-- EAV table to keep track of any user preferences
-- PLANNED -> use this to track whether user wants to allow a budget to be changeable 
-- after month has started
CREATE TABLE user_preference_choice (
	id serial PRIMARY KEY,
	user_id int REFERENCES budget_user(id) NOT NULL,
	choice_key text,
	choice_value text
);

-- a budget that will be referenced by budget_category_assignment & budget_category.
CREATE TABLE budget (
	id serial PRIMARY KEY,
	user_id int REFERENCES budget_user(id),
	month_id int, -- this budget is a 'model' if month & year are null, actual if given a month
	year int,
	name text
);

-- a category that will get an associated dollar value. Can exist for either a model budget or a 
-- monthly budget
CREATE TABLE budget_category (
	id serial PRIMARY KEY,
	name text
);

-- a category & dollar value pair that can be assigned to either a model_budget 
-- or a monthly_budget
CREATE TABLE budget_category_assignment (
	id serial PRIMARY KEY,
	budget_id int REFERENCES budget(id),
	budget_category_id int REFERENCES budget_category(id) NOT NULL,
	dollar_allotment int NOT NULL
);

CREATE TABLE tag (
	id serial PRIMARY KEY,
	user_id int REFERENCES budget_user(id) NOT NULL,
	name text
);

CREATE TABLE expenditure (
	id serial PRIMARY KEY,
	budget_category_assignment_id int REFERENCES budget_category_assignment(id) NOT NULL,
	value numeric NOT NULL,
	description text,
	note text,
	create_date timestamp
);

CREATE TABLE expenditure_tag (
	id serial PRIMARY KEY,
	tag_id int REFERENCES tag(id),
	expenditure_id int REFERENCES expenditure(id) NOT NULL
);

CREATE TABLE token (
	id serial PRIMARY KEY,
	user_id int REFERENCES budget_user(id) NOT NULL,
	value text,
	create_date timestamp,
	expiration_date timestamp
);