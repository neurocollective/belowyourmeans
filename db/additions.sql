insert into starter_budget_category (display_name, description) values ('Storage', 'just dont let it become a STORAGE WAR');

insert into budget_category (
	select nextval('budget_category_id_seq'),
	1,
	'Storage',
	'just dont let it become a STORAGE WAR',
	FALSE,
	now(),
	now()
);

insert into starter_budget_category (display_name, description) values ('Spouse', 'smoochie smoochie');

insert into budget_category (
	select nextval('budget_category_id_seq'),
	1,
	'Spouse',
	'smoochie smoochie',
	FALSE,
	now(),
	now()
);

insert into starter_budget_category (display_name, description) values ('Activism', 'i am the change i want to see');

insert into budget_category (
	select nextval('budget_category_id_seq'),
	1,
	'Activism',
	'i am the change i want to see',
	FALSE,
	now(),
	now()
);
