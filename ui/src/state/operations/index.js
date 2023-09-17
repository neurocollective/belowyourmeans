import fetching from '../../fetching';

const { jsonRequest } = fetching;

const DEFAULT_REQUEST_CONFIG = {
	headers: {
		'Content-Type': 'application/json'
	},
	method: 'GET',
	mode: 'no-cors',
	credentials: 'same-origin',
};

const buildOperations = (reducers) => ({
	handleLoginSubmit: () => {
		const { handleLoginSuccess, handleLoginFailure } = reducers;
		jsonRequest("/login", DEFAULT_REQUEST_CONFIG, handleLoginSuccess, handleLoginFailure);
	},
});

export default buildOperations;
