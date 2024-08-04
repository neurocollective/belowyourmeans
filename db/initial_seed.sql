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
	('Utilities', 'keep the lights on'),
	('IGNORED', 'expenditures that should not be included in totals');

insert into budget_category (
	select nextval('budget_category_id_seq'),
	1,
	sbc.display_name,
	sbc.description,
	FALSE,
	now(),
	now() from starter_budget_category sbc
);

-- use this when a new user is created
-- insert into budget_category (select nextval('budget_category_id_seq'), $1, sbc.display_name, sbc.description, now(), now() from starter_budget_category sbc);

insert into expenditure values 
	(nextval('expenditure_id_seq'), 1, null, 20.99, 'Digital Card Purchase - NETFLIX COM LOS GATOS CA', to_timestamp(1699920860387), now(), now()),
	(nextval('expenditure_id_seq'), 1, null, 800.00, 'NOT cocaine', to_timestamp(1699920848387), now(), now()),
	(nextval('expenditure_id_seq'), 1, null, 20000.00, 'Darkweb Gambling', to_timestamp(1699920836387), now(), now()
);

insert into budget_category_preassignment (category_id, description) values
	(1, 'Digital Card Purchase - NETFLIX COM LOS GATOS CA'),
	(2,'Debit: Withdrawal from AMEX EPAYMENT ACH PMT'),
	(15,'Debit: Digital Card Purchase - DIGITALOCEAN COM NEW YORK CIT NY'),
	(15,'Debit: Debit Card Purchase - GOOGLE GOOGLE STORAGE 650 253 0000 CA'),
	(11,'Debit: Preauthorized Withdrawal to AMERICAN EXPRESS NATIONAL BANK savings account XXXXXXXX3012'),
	(1,'Debit: Digital Card Purchase - HELP MAX COM NEW YORK NY'),
	(15,'Debit: Digital Card Purchase - MEETUP ORG SUB 1M NEW YORK NY'),
	(19,'Debit: Digital Card Purchase - 1PASSWORD TORONTO ON'),
	(14,'Debit: Digital Card Purchase - BUBBLESANDSUDSLAUNDRO BROOKLYN NY'),
	(1,'Debit: Digital Card Purchase - NETFLIX COM LOS GATOS CA'),
	(22,'Debit: Bill payment to Narrows Bayview LLC'),
	(11,'Debit: Withdrawal from VENMO PAYMENT'),
	(15,'Debit: Debit Card Purchase - GOOGLE GSUITE NEUROCO 650 253 0000 CA'),
	(24,'Debit: Withdrawal from CON ED OF NY XXXXXXXXXX'),
	(24,'Debit: Debit Card Purchase - VZWRLSS APOCC VISN 800 922 0204 FL'),
	(22,'Debit: Paper Payment to Narrows Bayview LLC'),
	(5,'Debit: Debit Card Purchase - TST J P GIFFORD MARK KENT CT'),
	(14,'Debit: Debit Card Purchase - BUBBLESANDSUDSLAUNDRO BROOKLYN NY'),
	(11,'Debit: Withdrawal to Capital One Bank  XXXXXX1903'),
	(15,'Debit: Debit Card Purchase - GOOGLE GSUITE NEUROCO MOUNTAIN VIE CA'),
	(11,'Debit: Check #0 Cashed'),
	(11,'Debit: ATM Withdrawal - 000000000206341 TX032510 NEWARK  NJ'),
	(11,'Debit: ATM Withdrawal - WALGREENS # -XE1 AXE10344 BROOKLYN  NY'),
	(11,'Debit: ATM Withdrawal - WALGREENS #1-000 A0004950 CANAAN  CT'),
	(11,'Debit: ATM Withdrawal - FORT HAMILTO-111272 P111272 BROOKLYN   NY'),
	(11,'Debit: ATM Withdrawal - CUMBERLAND F-U560054 CU560054 BENNINGTON  VT'),
	(11,'Debit: ATM Withdrawal - CU560827 CU560827 WILLIAMSTOWN  MA'),
	(11,'Debit: ATM Withdrawal - CITIBAN0020293 00202093 BROOKLYN  NY'),
	(11,'Debit: ATM Withdrawal - 7ELEVEN-FCTI 7E003049 BROOKLYN  NY'),
	(11,'Debit: ATM Withdrawal - 502-512 86TH ST 00978091 BRKLYN  NY'),
	(11,'Debit: Check #338 Cashed'),
	(11,'Debit: Check #340 Cashed'),
	(11,'Debit: Check #341 Cashed'),
	(11,'Debit: Check #343 Cashed'),
	(11,'Debit: Check #344 Cashed'),
	(11,'Debit: Check #345 Cashed'),
	(11,'Debit: Check #346 Cashed'),
	(11,'Debit: Check #348 Cashed'),
	(11,'Debit: Check #350 Cashed'),
	(11,'Debit: Check #352 Cashed'),
	(15,'Debit: Debit Card Purchase - CISCO SYSTEMS INC 9193922254 CA');