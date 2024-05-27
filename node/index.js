const express = require('express');
const pg = require('pg');

const {
	env: {
		PORT = 3001,
	}
} = process;

const app = express();

const { Client } = pg;
const clientConfig = {
	connectionString: 'postgres://postgres:postgres@localhost:5432/postgres',
};
const client = new Client(clientConfig);

// pg.types.setTypeParser(pg.types.builtins.DECIMAL, value => parseFloat(value));

const boot = async () => {

	await client.connect();

	app.get('/', async (req, res) => {

		const dbRes = await client.query('SELECT $1::text as message', ['Hello world!']);

		const { rows: [{ message }]} = dbRes;

		return res.json({ status: message });
	});

	app.get('/expenditure', async (req, res) => {

		const dbRes = await client.query('SELECT * from expenditure;', []);

		const { rows } = dbRes;

		console.log('rows', rows);

		return res.json(rows);
	});


	app.post('/query', express.json(), async (req, res) => {

		let { query, parameters, specialRule } = (req.body ?? {});

		console.log("req.body", req.body);
		console.log("query:", query);
		console.log("parameters:", parameters);

		if (typeof query !== "string") {
			return res.status(400).json({ error: "query is not a string" });
		}
		if (!Array.isArray(parameters)) {
			parameters = [];
		}

		const dbRes = await client.query(query, parameters);

		let { rows = [] } = dbRes;

		if (specialRule === "mapExpenditures") {
			rows = rows.map((row) => {
				const newRow = { ...row };
				try {
					newRow.value = parseFloat(row.value);
				} catch (err) {
					console.error(err);
					return row;
				}
				return newRow;
			})
		}

		console.log('rows[0:9]', rows.slice(0,9));

		return res.json(rows);
	});

	app.listen(PORT, () => {
		console.log(`listening on ${PORT}`);
	});
};

boot().catch((error) => {
	console.error(error);
});
