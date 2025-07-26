const { exec } = require('child_process');

const runCmd = (command = '') => new Promise((resolve, reject) => {
	exec(command, {}, (error, stdout, stderr) => {
		const out = `${stdout}.toString()${stderr.toString()}`;
		if (error) {
			return reject('error: ' + error.message + " | " + out);
		}
		return resolve(out);
	});
});

const range = n => new Array(n).fill().map((_, i) => String(i + 1));

const reducer = (accumulator, value) => accumulator.concat(value);

const getCommand = (month, year, isCapOne) => {

	const label = isCapOne ? 'capone' : 'amex';
	const file = `${label}_1_${month}_${year}.csv`;

	return `make upload month=${month} year=${year} file=${file} capone=${Boolean(isCapOne)}`;
};

const main = () => {

	const year = '2025';

	const promises = range(6).map((i) => {
		return [
			runCmd(getCommand(i, year, true)),
			runCmd(getCommand(i, year, false)),
		];
	}).reduce(reducer);

	console.log('promises');
	
	Promise.all(promises).then((out) => {
		console.log('done ->');
		console.log(out);
	}).catch(console.error);
};

main();
