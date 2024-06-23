import jsonRequest from '../../fetching';
import { LOGIN, NAVIGATION, HOME, EXPENDITURES, CATEGORIZE } from '../../constants';

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

	if (hostname === "localhost") {
		return `http://${hostname}:8080${path}`;
	}
	return origin + path;
}

const buildNavigate = (stateChanges) => {
	return (location) => {
		stateChanges[NAVIGATION].setLocation(location);
	};
};

const buildOperations = (state, stateChanges) => {

	const navigate = buildNavigate(stateChanges);

	return {
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

				const config = {
					...DEFAULT_REQUEST_CONFIG,
					method: 'GET',
				};

				const fullURL = getURL("/auth");
				console.log(`fetching to ${fullURL}`);

				return jsonRequest(fullURL, config, handleLoggedIn, handleNotLoggedIn);
			},
			logout: () => {
				console.log('logout coming soon?');
				// const {
				// 	[LOGIN]: {
				// 		handleLogoutSuccess,
				// 		handleLogoutFailure, 
				// 	} 
				// } = stateChanges;

				// const config = {
				// 	...DEFAULT_REQUEST_CONFIG,
				// };

				// const fullURL = getURL("/auth");
				// console.log(`fetching to ${fullURL}`);

				// return jsonRequest(fullURL, config, handleLogoutSuccess, handleLogoutFailure);
			},
		},
		[NAVIGATION]: {
			navigate,
		},
		[CATEGORIZE]: {
			getUncategorizedExpenditures: () => {

			},
		},
		[EXPENDITURES]: {
			getExpenditures: () => {

				const {
					[EXPENDITURES]: {
						handleExpenditureNavigationSuccess: ok,
						handleExpenditureNavigationFailure: fail, 
					}
				} = stateChanges;

				const {
					[LOGIN]: {
						user: userId
					},
					[EXPENDITURES]: {
						month,
					}
				} = state;

				const config = DEFAULT_REQUEST_CONFIG;

				const fullURL = getURL(`/expenditure?userId=${userId}&month=${month}`);

				return jsonRequest(fullURL, config, ok, fail);
			},
		},
	};
};

export default buildOperations;
