insert into budget_category (name) values ('entertainment'),('healthcare'),
	('grocery'),('transportation'),('eating out'),('clothing'),('other'),
	('home maintenance'),('business'),('petcare'),('hobby'),('taxes'),('fees'),
	('fitness & wellness'),('investment'),('rent'),('mortgage'),('utility'),
	('UNASSIGNED');

insert into budget_user values (
	nextval('budget_user_id_seq'),
	'david',
	'ashe',
	'david@neurocollective.io',
	'$5$hWvy8n69/SiNtC3d$rxYVgSbIlZC46tmMwZ/rE7fgz8iDemF7dIs8MCNkZB7',
	now(),
	now()
);

insert into budget_concept values (nextval('budget_concept_id_seq'),1,null,null,'test model budget');
insert into applied_budget values (nextval('budget_ideation_id_seq'),null,null,null,'test model budget');
insert into budget_concept values (nextval('budget_concept_id_seq'),1,11,2017,'november test budget concept');
insert into applied_budget values (nextval('budget_ideation_id_seq'),1,11,2017,'november test budget');

insert into applied_budget (budget_id, budget_category_id, dollar_allotment) 
	select 2 as budget_id, id as budget_category_id, 5000 as dollar_allotment
		from budget_category;

insert into expenditure (budget_category_assignment_id, value, description,create_date)
	values (1,200,'test value',now()),(2,200,'test value',now()),(3,200,'test value',now()),
		(4,200,'test value',now()),(5,200,'test value',now()),(6,200,'test value',now()),
		(7,200,'test value',now()),(8,200,'test value',now()),(9,200,'test value',now()),
		(10,300,'test value',now()),(11,560,'test value',now()),(12,208,'test value',now()),
		(13,260,'test value',now()),(14,600,'test value',now()),(15,2000,'test value',now()),
		(16,2500,'test value',now()),(17,100,'test value',now()),(18,700,'test value',now());

insert into expenditure values (20,18,50,'another test value',now());

insert into tag (user_id, name) values (1, 'alcohol'),(1, 'mta'),(1, 'lyft'),(1, 'halal');

insert into expenditure_tag (tag_id, expenditure_id) values (1,13);		