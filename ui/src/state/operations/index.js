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

			e.preventDefault();

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
				body: JSON.stringify({ email, password }),
				method: 'POST',
			};

			const fullURL = getURL("/login");
			console.log(`fetching to ${fullURL}`);

			return jsonRequest(fullURL, config, handleLoginSuccess, handleLoginFailure);
		},
		checkIfLoggedIn: () => {
			const {
				[LOGIN]: {
					handleLoggedIn,
					handleNotLoggedIn, 
				} 
			} = stateChanges;

			// const {
			// 	[LOGIN]: {
			// 		email,
			// 		password,
			// 	}
			// } = state;

			const config = {
				...DEFAULT_REQUEST_CONFIG,
				method: 'GET',
			};

			const fullURL = getURL("/auth");
			console.log(`fetching to ${fullURL}`);

			return jsonRequest(fullURL, config, handleLoggedIn, handleNotLoggedIn);
		},
	}
});

export default buildOperations;
