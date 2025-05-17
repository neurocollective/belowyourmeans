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
	('Healthcare', 'keeping your body and mind healthy'),
	('Groceries', 'food for home'),
	('Transportation', 'cars, maintenance, mass transit'),
	('Restaurants', 'sitdown, food truck, cart, stall, whatever'),
	('Coffee', 'rich, wakey goodness'),
	('Beverages', 'suite suite Refreshment'),
	('Alcohol', 'stirred, not shaken'),
	('Cigarettes', 'now film me in black & white'),
	('Clothing', 'utility or fashion, maybe both?'),
	('Other', 'whatever doesn''t fit anywhere else'),
	('Kids', 'they need a lot, every day'),
	('Spouse', 'smoochie smoochie'),
	('Home Maintenance', 'keep the castle'),
	('Business', 'hustlin'),
	('Pets', 'furry/feathered/scaled children'),
	('Recreation', 'hobbies, glee, enjoyment'),
	('Taxes', 'when ur rich u wont pay em'),
	('Fees', 'death, taxes, and fees'),
	('Fitness and Wellness', 'sharpen the axe'),
	('Investment', 'someday i want to be that guy'),
	('Rent', 'gotta sleep somewhere'),
	('Mortgage', 'gotta sleep somewhere - with the bank''s permission'),
	('Utilities', 'keep the lights on'),
	('IGNORED', 'expenditures that should not be included in totals'),
	('Streaming Video', 'remember DVDs?'),
	('Laundry', 'i am going to pretend i did''t see this stain'),
	('Streaming Video', 'remember DVDs?'),
	('Savings', 'inflation hates me but i don''t care'),
	('Discretionary Purchase', 'i saw it, i wanted it'),
	('Self-care', 'haircuts, massages, nails, whatever your well-being requires'),
	('Education', 'lurnin'),
	('Hotels & Lodging', 'no one will miss this bathrobe...'),
	('Insurance', 'my stuff can resurrect into money'),
	('Cars', 'vroom');

insert into budget_category (
	select nextval('budget_category_id_seq'),
	1,
	sbc.display_name,
	sbc.description,
	FALSE,
	now(),
	now() from starter_budget_category sbc
);
