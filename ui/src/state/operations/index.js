import jsonRequest from '../../fetching';
import { HEADER, LOGIN, NAVIGATION, HOME, EXPENDITURES, CATEGORIES, REPORTS } from '../../constants';

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
		return `http://${hostname}:8080/api${path}`;
	}
	return `${origin}/api${path}`;
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
			},
			[HEADER]: {
				selectedYear: year,
			},
		} = state;

		if (!userId) {
			console.error('no user id in getExpenditures!');
			return;
		} else {
			console.log('userId in getExpenditures:', userId);
		}

		console.log('year before API call is:', year)

		const fullURL = getURL(`/expenditure?userId=${userId}&month=${month}&year=${year}`);

		return jsonRequest(fullURL, DEFAULT_REQUEST_CONFIG, ok, fail);
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
			applyCategoryItem: (expenditureDescription, categoryName, expenditureId) => {
				const {
					[LOGIN]: {
						user: userId
					},
					[CATEGORIES]: {
						categories,
					}
				} = state;
				const {
					[CATEGORIES]: {
						handleApplyCategoryItemSuccess: ok,
						handleApplyCategoryItemFailure: fail,
					}
				} = stateChanges;

				const { id: categoryId } = categories.find(({ ['display_name']: name }) => {
					return name === categoryName;
				});

				const fullURL = getURL("/category-preassignment/apply");
				const config = {
					...DEFAULT_REQUEST_CONFIG,
					method: 'POST',
					body: JSON.stringify({
						expenditureDescription, categoryName, expenditureId, categoryId,
					}),
				};
				return jsonRequest(fullURL, config, ok, fail);
			},
			setSelectedBroadCategory: (exenditureId, selectedCategoryName) => {
				const {
					[CATEGORIES]: {
						updateBroadSelection,
					}
				} = stateChanges;
				updateBroadSelection(exenditureId, selectedCategoryName);
			},
			getUncategorizedExpenditures: () => {
				const {
					// [LOGIN]: {
					// 	user: userId
					// },
					[HEADER]: {
						selectedYear: year,
					}
				} = state;
				const {
					[CATEGORIES]: {
						handleGetUncategorizedExpendituresSuccess: ok,
						handleGetUncategorizedExpendituresFailure: fail,
					}
				} = stateChanges;
				const fullURL = getURL(`/categories/expenditures/names?year=${year}`);
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
				console.log('expenditureId:', expenditureId);

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
		[REPORTS]: {
			getReports: () => {
				const {
					// [LOGIN]: {
					// 	user: userId
					// },
					[HEADER]: {
						selectedYear: year,
						years,
					},
					// [REPORTS]: {
					// 	categories,
					// }
				} = state;
				const {
					[REPORTS]: {
						handleGetReportSuccess: ok,
						handleGetReportFailure: fail,
						handleGetPriorReportSuccess: priorOk,
						handleGetPriorReportFailure: priorFail,
					}
				} = stateChanges;

				let months = year === '2025' ? '6' : '12';

				const fullURL = getURL(`/report/annualized?year=${year}&months=${months}`);
				const config = DEFAULT_REQUEST_CONFIG;

				const one = jsonRequest(fullURL, config, ok, fail).then(() => {

					const currentYearIndex = years.findIndex((yearValue) => yearValue === year);
					const priorYear = years[currentYearIndex - 1]

					if (priorYear) {

						months = priorYear === '2025' ? '6' : '12';

						const fullURLTwo = getURL(`/report/annualized?year=${priorYear}&months=${months}`);
						const config = DEFAULT_REQUEST_CONFIG;

						const two = jsonRequest(fullURLTwo, config, priorOk, priorFail);
					}
				});
			},
		},
	};
};

export default buildOperations;
