import { LOGIN, NAVIGATION, HOME, EXPENDITURES, CATEGORIES, HEADER, REPORTS } from '../../constants';

const buildStateChanges = (state, setState, getInitialState) => {

	const update = (newState, callback) => {
		if (callback) {
			callback(newState);
		}
		return setState(newState);
	};

	return {
		[LOGIN]: {
			handleEmailType: (e) => {
				const { target: { value } } = e;
				const newState = { ...state, [LOGIN]: { ...state[LOGIN], email: value } };
				update(newState);
			},
			handlePasswordType: (e) => {
				const { target: { value } } = e;
				const newState = { ...state, [LOGIN]: { ...state[LOGIN], password: value } };
				update(newState);
			},
			handleLoginSuccess: (successPayload) => {
				console.log('handleLoginSuccess successPayload:', successPayload);
				const { data: { userId } } = successPayload;
				const newState = {
					...state,
					[LOGIN]: {
						...state[LOGIN],
						isLoggedIn: true,
						email: '',
						password: '',
						user: userId,
					},
					[NAVIGATION]: {
						current: HOME,
					},
				};
				update(newState);
			},
			handleLoginFailure: (failurePayload) => {
				const newState = { ...state, [LOGIN]: { ...state[LOGIN], isLoggedIn: false } };
				update(newState);
			},
			handleLoggedIn: (successPayload) => {
				console.log('handleLoggedIn successPayload', successPayload);
				const { userId } = successPayload;
				const newState = {
					...state,
					[LOGIN]: {
						...state[LOGIN],
						isLoggedIn: true,
						user: userId,
					},
					[NAVIGATION]: {
						current: HOME
					},
				};
				console.log('newState:', newState);
				update(newState);
			},
			handleNotLoggedIn: (failurePayload) => {
				console.log('handleNotLoggedIn failurePayload', failurePayload);
				const newState = {
					...state,
					[LOGIN]: {
						...state[LOGIN],
						isLoggedIn: false
					},
					[NAVIGATION]: {
						current: LOGIN,
					},
				};
				update(newState);
			},
			handleLogoutSuccess: (successPayload) => {
				console.log('handleLogoutSuccess successPayload', successPayload);
				const newState = {
					...state,
					[LOGIN]: {
						...getInitialState()[LOGIN],
						isLoggedIn: false,
					},
					[NAVIGATION]: {
						current: LOGIN,
					},
				};
				update(newState);
			},
			handleLogoutFailure: (failurePayload) => {
				console.log('failurePayload', failurePayload);
				const newState = {
					...state,
					[LOGIN]: {
						...getInitialState()[LOGIN],
						isLoggedIn: false,
					},
					[NAVIGATION]: {
						current: LOGIN,
					},
				};
				update(newState);
			},
		},
		[EXPENDITURES]: {
			handleExpenditureNavigationSuccess: (successPayload) => {
				const { data: expenditures } = successPayload;
				const newState = {
					...state,
					[EXPENDITURES]: {
						...state[EXPENDITURES],
						expenditures,
					},
					[NAVIGATION]: {
						current: EXPENDITURES,
					},
				};
				update(newState);				
			},
			handleExpenditureNavigationFailure: (failurePayload) => {
				const { error } = failurePayload;
				const newState = {
					...state,
					[EXPENDITURES]: {
						...state[EXPENDITURES],
						getExpenditureError: error,
					},
					[NAVIGATION]: {
						current: EXPENDITURES,
					},
				};
				update(newState);
			},
			setMonth: (month) => {
				const newState = {
					...state,
					[EXPENDITURES]: {
						...state[EXPENDITURES],
						month,
					},
				};
				update(newState);
			},
		},
		[NAVIGATION]: {
			setLocation: (location) => {
				const newState = {
					...state,
					[NAVIGATION]: {
						current: location,
					},
				};
				update(newState);	
			}
		},
		[CATEGORIES]: {
			updatePreAssignmentFullNameSelection: (expenditureId, selectedCategoryName) => {
				
				const preAssignmentMap = { ...state[CATEGORIES].preAssignmentMap }
				preAssignmentMap[expenditureId] = {
					name: selectedCategoryName,
					fullName: true,
				};

				const newState = {
					...state,
					[CATEGORIES]: {
						...state[CATEGORIES],
						preAssignmentMap,
					},
				};
				update(newState);			
			},
			updatePreAssignmentPattern: (expenditureId, pattern) => {
				
				const preAssignmentMap = { ...state[CATEGORIES].preAssignmentMap }
				preAssignmentMap[expenditureId] = {
					name: null,
					fullName: false,
					pattern,
				};

				const newState = {
					...state,
					[CATEGORIES]: {
						...state[CATEGORIES],
						preAssignmentMap,
					},
				};
				update(newState);			
			},
			handleGetCategoriesSuccess: (successPayload) => {
				const { data: categories } = successPayload;
				const newState = {
					...state,
					[CATEGORIES]: {
						...state[CATEGORIES],
						categories,
					},
				};
				setState(newState);
				// update(newState);
			},
			handleGetCategoriesFailure: (failurePayload) => {
				console.error('ruh roh error payload in handleGetCategoriesFailure:', failurePayload);
			},
			setSelectedCategory: (expenditureId, categoryName) => {

				const selectionsMap = { ...state[CATEGORIES].selectionsMap };

				selectionsMap[expenditureId] = categoryName;

				const newState = {
					...state,
					[CATEGORIES]: {
						...state[CATEGORIES],
						selectionsMap,
					},
				};
				update(newState);				
			},
			handleUpdateExpenditureFailure: (failurePayload) => {
				console.error('OH NOES BROES:', failurePayload);
				const newState = {
					...state,
					[CATEGORIES]: {
						...state[CATEGORIES],
						error: 'failed to update category for expenditure',
					},
				};
				update(newState);					
			},
			handleGetUncategorizedExpendituresSuccess: (successPayload) => {
				const { data: expenditureNames } = successPayload;
				const newState = {
					...state,
					[CATEGORIES]: {
						...state[CATEGORIES],
						expenditureNames,
					},
				};
				const print = (newState) => {
					console.log(`handleGetUncategorizedExpendituresSuccess has list size: ${expenditureNames.length}`);
					console.log('newState will be:', newState)
				}
				print(newState);
				setState(newState);
				//update(newState, print);
			},
			handleGetUncategorizedExpendituresFailure: (failurePayload) => {
				const { error } = failurePayload;
				const newState = {
					...state,
					[CATEGORIES]: {
						...state[CATEGORIES],
						error,
					},
				};
				update(newState);
			},
			handleApplyCategoryItemSuccess: (successPayload) => {
				console.log('successPayload for handleApplyCategoryItemSuccess', successPayload);
				// update state?
			},
			handleApplyCategoryItemFailure: (failurePayload) => {
				console.log('failurePayload for handleApplyCategoryItemFailure', failurePayload);
			},
		},
		[HEADER]: {
			handleYearSelect: (e) => {
				const { target: { value } } = e;
				const newState = {
					...state,
					[HEADER]: {
						...state[HEADER],
						selectedYear: value,
					},
				};
				update(newState);
			},
		},
		[REPORTS]: {
			handleGetReportSuccess: ({ data: report }) => {
				const newState = {
					...state,
					[REPORTS]: {
						...state[REPORTS],
						report,
					},
				};
				console.log('handleGetReportSuccess setting new state:', newState[REPORTS]);
				setState(newState);				
			},
			handleGetReportFailure: (failurePayload) => {
				console.log('failurePayload for handleGetReportFailure', failurePayload);
			},
			handleGetPriorReportSuccess: ({ data: priorReport }) => {
				const newState = {
					...state,
					[REPORTS]: {
						...state[REPORTS],
						priorReport,
					},
				};
				console.log('handleGetPriorReportSuccess setting new state:', newState[REPORTS]);
				setState(newState);
			},
			handleGetPriorReportFailure: (failurePayload) => {
				console.log('failurePayload for handleGetPriorReportFailure', failurePayload);
			},
		}
	};
};

export default buildStateChanges;
