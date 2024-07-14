import jsonRequest from '../../fetching';
import { LOGIN, NAVIGATION, HOME, EXPENDITURES, CATEGORIES } from '../../constants';

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

	const getExpenditures = () => {

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

		if (!userId) {
			console.error('no user id in getExpenditures!');
			return;
		} else {
			console.log('userId in getExpenditures:', userId);
		}

		const config = DEFAULT_REQUEST_CONFIG;

		const fullURL = getURL(`/expenditure?userId=${userId}&month=${month}`);

		return jsonRequest(fullURL, config, ok, fail);
	};

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
		[CATEGORIES]: {
			getUncategorizedExpenditures: () => {
				const {
					[LOGIN]: {
						user: userId
					}
				} = state;
				const {
					[CATEGORIES]: {
						handleGetUncategorizedExpendituresSuccess: ok,
						handleGetUncategorizedExpendituresFailure: fail,
					}
				} = stateChanges;
				const fullURL = getURL("/categories/expenditures/names");
				return jsonRequest(fullURL, DEFAULT_REQUEST_CONFIG, ok, fail);
			},
			getCategories: () => {
				const {
					[LOGIN]: {
						user: userId
					}
				} = state;
				const {
					[CATEGORIES]: {
						handleGetCategoriesSuccess,
						handleGetCategoriesFailure,
					}
				} = stateChanges;
				const fullURL = getURL(`/categories?userId=${userId}`);
				const ok = handleGetCategoriesSuccess;
				const fail = handleGetCategoriesFailure;
				return jsonRequest(fullURL, DEFAULT_REQUEST_CONFIG, ok, fail);
			},
			setSelectedCategory: (expenditureId, categoryName) => {
				return stateChanges[CATEGORIES].setSelectedCategory(expenditureId, categoryName)
			},
			updateCategoryForExpenditure: (expenditureDescription, categoryName, expenditureId) => {
				const {
					[CATEGORIES]: {
						// handleUpdateExpenditureSuccess,
						handleUpdateExpenditureFailure,
					}
				} = stateChanges;

				const {
					[CATEGORIES]: {
						categories,
						selectionsMap,
					}
				} = state;

				const { id: categoryId } = {} = categories.find((c) => {
					return c['display_name'] === selectionsMap[expenditureId];
				});

				if (!categoryId) {
					console.error(`could not find id for ${categoryName}!`);
					return;
				}

				console.log('categoryId:', categoryId);
				console.log('expenditureId:', expenditureId)

				const fullURL = getURL(`/expenditure`);

				const fail = handleUpdateExpenditureFailure;
				const config = {
					...DEFAULT_REQUEST_CONFIG,
					method: 'PUT',
					body: JSON.stringify({ categoryId, expenditureId }),
				}

				console.log('updateCategory boutta make API call');

				return jsonRequest(fullURL, config, getExpenditures, fail);
			},
		},
		[EXPENDITURES]: {
			getExpenditures,
			setMonth: (monthIndex) => {
				const { [EXPENDITURES]: { setMonth } } = stateChanges;
				setMonth(monthIndex);
			},
		},
	};
};

export default buildOperations;
