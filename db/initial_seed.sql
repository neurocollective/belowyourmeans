insert into budget_user values (
	nextval('budget_user_id_seq'),
	'david',
	'ashe',
	'david@neurocollective.io',
	'$2a$10$8.lPRFurMUrF8Piv6hrhiOyyfwzTa1JafiF2mMLP.YnZefwTp/Qgu', -- make it the hash of 'password' for now
	now(),
	now()
);

insert into starter_budget_category (display_name, description) values 
	('Entertainment', 'anything and everything fun'),
	('Healthcare', 'keeping your self healthy'),
	('Groceries', 'food for home'),
	('Transportation', 'cars, maintenance, mass transit'),
	('Restaurants', 'anything and everything fun'),
	('Coffee', 'Rich, wakey goodness'),
	('Beverages', 'Suite Suite Refreshment'),
	('Alcohol', 'I need a night off'),
	('Cigarettes', 'I need a day off right now'),
	('Clothing', 'utility or fashion, maybe both?'),
	('Other', 'whatever doesn''t fit anywhere else'),
	('Kids', 'they need a lot, every day'),
	('Spouse', 'Smoochie smoochie'),
	('Home Maintenance', 'keep the castle'),
	('Business', 'hustlin'),
	('Pets', 'furry/feathered/scaled children'),
	('Hobby', 'just for fun'),
	('Taxes', 'when ur rich u wont pay em'),
	('Fees', 'death, taxes, and fees'),
	('Fitness and Wellness', 'sharpen the axe'),
	('Investment', 'someday i want to be that guy'),
	('Rent', 'gotta sleep somewhere'),
	('Mortgage', 'gotta sleep somewhere - with the bank''s permission'),
	('Utilities', 'keep the lights on');

insert into budget_category (select nextval('budget_category_id_seq'), 1, sbc.display_name, sbc.description, now(), now() from starter_budget_category sbc);

-- use this when a new user is created
-- insert into budget_category (select nextval('budget_category_id_seq'), $1, sbc.display_name, sbc.description, now(), now() from starter_budget_category sbc);

insert into expenditure values 
	(nextval('expenditure_id_seq'), 1, null, 20.99, 'Digital Card Purchase - NETFLIX COM LOS GATOS CA', to_timestamp(1699920860387), now(), now()),
	(nextval('expenditure_id_seq'), 1, null, 800.00, 'NOT cocaine', to_timestamp(1699920848387), now(), now()),
	(nextval('expenditure_id_seq'), 1, null, 20000.00, 'Darkweb Gambling', to_timestamp(1699920836387), now(), now()
);

insert into budget_category_items values (
	1, 'Digital Card Purchase - NETFLIX COM LOS GATOS CA', now(), now()
);

-- send in user_id, description
-- insert into budget_category_items values (
-- 	nextval('budget_category_items_id_seq'), $1, $2, now(), now()
-- );
