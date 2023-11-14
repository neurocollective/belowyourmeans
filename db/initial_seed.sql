insert into budget_user values (
	nextval('budget_user_id_seq'),
	'david',
	'ashe',
	'david@neurocollective.io',
	'$5$hWvy8n69/SiNtC3d$rxYVgSbIlZC46tmMwZ/rE7fgz8iDemF7dIs8MCNkZB7', -- make it the hash of 'test123' for now
	now(),
	now()
);

insert into expenditure values 
	(nextval('expenditure_id_seq'), 1, 20.99, to_timestamp(1699920860387), now(), now()),
	(nextval('expenditure_id_seq'), 1, 199.98, to_timestamp(1699920848387), now(), now()),
	(nextval('expenditure_id_seq'), 1, 3.99, to_timestamp(1699920836387), now(), now());
