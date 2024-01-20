import jsonRequest from '../../fetching';
import { LOGIN } from '../../constants';

const DEFAULT_REQUEST_CONFIG = {
	headers: {
		'Content-Type': 'application/json',
		'Accept': 'application/json',
	},
	method: 'GET',
	mode: 'cors',
	credentials: 'include',
};

const getURL = (path) => {
	const {
		location: {
			hostname,
			port,
			origin,
		}
	} = window;

	if (hostname == "localhost") {
		return `http://${hostname}:8080${path}`;
	}
	return origin + path;
}

const buildOperations = (state, stateChanges) => ({
	[LOGIN]: {
		handleLoginSubmit: (e) => {

			const {
				[LOGIN]: {
					handleLoginSuccess,
					handleLoginFailure, 
				} 
			} = stateChanges;

			const {
				[LOGIN]: {
					email,
					password,
				}
			} = state;

			const config = {
				...DEFAULT_REQUEST_CONFIG,
				body: { email, password },
				method: 'POST',
			};
			console.log('config', JSON.stringify(config));
			window.CONFIG = config;

			const fullURL = getURL("/login");
			console.log(`fetching to ${fullURL}`);

			return jsonRequest(fullURL, config, handleLoginSuccess, handleLoginFailure);
		},
	}
});

export default buildOperations;
