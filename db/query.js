const noNos = [
	';',
	' union select',
	' union insert',
	' union update',
	' union delete',
	' union all',
	'=',
	'$$'
];

const INVALID_INPUT = 'invalid input detected';

const wipeInjections = (strings) => {

	strings.forEach((str, index) => {

		if (typeof str !== 'string') {
			return;
			// throw new Error(INVALID_INPUT + ', not a string, input was:' + strings.join(','));
		}

		noNos.forEach((nono) => {
			if (str.toUpperCase().indexOf(nono.toUpperCase()) > -1) {
				throw new Error(INVALID_INPUT + `: ${nono}`);
			}
		});
	});
};

// const escapeApostrophe = (str) => {
// 		if (str.includes("'")) {
// 			return str.replace(/'/g, "\'");
// 		}
// 		return str;
// };

// TODO - excise this method from file, it is now useless
const escapeApostrophe = str => str;

const queries = {

	getBudgetsForUser : (userId) => {
		wipeInjections([userId]);
		return `select * from budget where user_id = ${userId} or user_id is null;`
	},

	getBudgetsForMonth : (userId, month) => {
		wipeInjections([userId, month]);
		return `select * from budget where month_id = ${month} and 
			(user_id = ${userId} or user_id is null);`
	},

	getAllTags : () => {
		return `select * from tag order by name;`;
	},

	getLatestBudget: () => {
		return `select * from budget where month_id is not null order by year, 
			month_id desc limit 1;`;
	},

	getModelBudgets: (userId) => {
		wipeInjections([userId]);
		return `select id, name from budget where 
			month_id is null and year is null and user_id = ${userId};`;
	},

	getBudgetAmountsForUserAndMonth : (userId, month, year) => {
		wipeInjections([userId, month, year]);

		return `select b.name, b.id, bca.id as assignment_id, bc.name as category_name, 
			bca.dollar_allotment as budgeted
			from budget b
			join budget_category_assignment bca 
			on bca.budget_id = b.id
			join budget_category bc 
			on bca.budget_category_id = bc.id
			where b.month_id = ${month}
			and b.user_id = ${userId}
			and b.year = ${year}
			order by bc.name`;
	},

	getBudgetAmountsForBudgetId : (budgetId) => {
		wipeInjections([budgetId]);

		return `select b.name, b.id as budget_id, bca.id as assignment_id, 
			bc.id as category_id,
			bca.dollar_allotment as budgeted
			from budget b 
			join budget_category_assignment bca 
			on bca.budget_id = b.id 
			join budget_category bc
			on bc.id = bca.budget_category_id
			where b.id = ${budgetId}
			order by bc.name;`;
	},

	getBudgetAmountsAndTotalSpentForUserAndMonth : (userId, month, year) => {
		wipeInjections([userId, month, year]);

		return `select * from (
			select b.name, b.id, bca.id as assignment_id, bc.name as category_name,
			bc.id as category_id, 
			bca.dollar_allotment as budgeted, coalesce(sum(e.value), 0) as spent  
			from budget b
			join budget_category_assignment bca 
			on bca.budget_id = b.id
			join budget_category bc 
			on bca.budget_category_id = bc.id
			left join expenditure e on e.budget_category_assignment_id = bca.id
			where b.month_id = ${month}
			and b.user_id = ${userId}
			and b.year = ${year}
			and lower(bc.name) != 'other'
			group by b.name, b.id, bca.id, bc.name, bc.id, bca.dollar_allotment
			order by bc.name) sub
			union all
			select b.name, b.id, bca.id as assignment_id, bc.name as category_name,
			bc.id as category_id, 
			bca.dollar_allotment as budgeted, coalesce(sum(e.value), 0) as spent  
			from budget b
			join budget_category_assignment bca 
			on bca.budget_id = b.id
			join budget_category bc 
			on bca.budget_category_id = bc.id
			left join expenditure e on e.budget_category_assignment_id = bca.id
			where b.month_id = ${month}
			and b.user_id = ${userId}
			and b.year = ${year}
			and lower(bc.name) = 'other'
			group by b.name, b.id, bca.id, bc.name, bc.id, bca.dollar_allotment;`
	},

	insertExpenditure : (catAssignId, value, description) => {
		wipeInjections([catAssignId, value, description]);
		value = escapeApostrophe(value);

		catAssignId = parseInt(catAssignId);
		value = parseFloat(value);

		const q = `insert into expenditure 
			(id, budget_category_assignment_id, value, description, create_date) values 
					(nextval('expenditure_id_seq'), ${catAssignId}, ${value}, 
						$$${description}$$, now()) RETURNING id;`						

		return q;
	},

	insertTag : (userId, name) => {
		wipeInjections([userId, name]);
		name = escapeApostrophe(name);

		userId = parseInt(userId);

		const q = `insert into tag values 
			(nextval('tag_id_seq'), ${userId}, $$${name}$$);`			

		return q;
	},

	insertExpenditureTag : (tagId, expenditureId) => {
		wipeInjections([tagId, expenditureId]);

		tagId = parseInt(tagId);
		expenditureId = parseFloat(expenditureId);

		const q = `insert into expenditure_tag values 
			(nextval('expenditure_tag_id_seq'), ${tagId}, ${expenditureId});`			

		return q;
	},

	updateExpenditure : (id) => {
		wipeInjections([id]);
		return `update expenditure where value = ${value} where id = ${id}`;
	},

	getAllExpendituresForCategoryAssignmentId : (budgetCatAssignId) => {
		wipeInjections([budgetCatAssignId]);
		return `select * from expenditure where 
			budget_category_assignment_id = ${budgetCatAssignId}`;
	},

	checkForValidToken : (token) => {
		wipeInjections([token]);
		return `select * from token where value = $$${token}$$ and 
			(expiration_date is null or expiration_date > now());`;
	},

	getUserForCredentials : (email, password) => {
		wipeInjections([email, password]);
		const query = `select * from budget_user where email=$$${email}$$ 
			and hashed_password=$$${password}$$`;
		return query;
	},

	insertToken : (userId, token) => {
		return `insert into token values (nextval('token_id_seq'),
			${userId}, $$${token}$$, now(), now() + interval '10 year')`;
	},

	wipeTokensForUserId: (userId) => {
		return `delete from token where user_id = ${userId} and (expiration_date 
			is not null and expiration_date < now())`;
	},

	getCurrentBudget: (userId, month, year) => {
		// wipeInjections([userId, month, year]);
		return `select * from budget b where b.user_id = $$${userId}$$ and 
			b.month_id = $$${month}$$ and b.year = '${year}'`;
	},

	getExpendituresForBudget: () => {
		return `select * from expenditure where `;
	},

	insertBudget: (userId, monthId, yearId, text) => {
		wipeInjections([userId, monthId, yearId, text]);
		text = escapeApostrophe(text);
		return `insert into budget values (nextval('budget_id_seq'), 
			${userId}, ${monthId}, ${yearId}, $$${text}$$);`;
	},

	copyBudgetCatAssignmentsFromOldBudget: (newBudgetId, oldBudgetId) => {
		wipeInjections([newBudgetId, oldBudgetId]);
		return `insert into budget_category_assignment 
			select nextval('budget_category_assignment_id_seq') as id, 
			${newBudgetId} as budget_id, budget_category_id, dollar_allotment 
			from budget_category_assignment where budget_id = ${oldBudgetId};`;
	},

	getAllBudgetCategories : () => {
		return `(select * from budget_category where name != 'other' order by name) 
			union all 
			select * from budget_category where name = 'other';`;
	},

	getBudget: (userId, month, year) => {
		wipeInjections([userId, month, year]);
		return `select * from budget where month_id = ${month} and year = ${year}
			and user_id = ${userId}`;
	},

	insertCategoryAssignment: (budgetId, budgetCategoryId, amount) => {
		wipeInjections([budgetId, budgetCategoryId, amount]);
		return `insert into budget_category_assignment values
			(nextval('budget_category_assignment_id_seq'), ${budgetId}, 
			${budgetCategoryId}, ${amount});`;
	},

	getUserForEmail: (email) => {
		wipeInjections([email]);
		// email = escapeApostrophe(email);

		return`select hashed_password, id from budget_user bu 
			where bu.email = '${email}'`;
	}

};

module.exports = queries
