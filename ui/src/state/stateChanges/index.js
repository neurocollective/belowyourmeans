import { LOGIN, NAVIGATION, HOME, EXPENDITURES } from '../../constants';

const buildStateChanges = (state, setState, getInitialState) => {

	const update = newState => setState(oldState => ({ ...oldState, ...newState }));

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
				const { userId } = successPayload;
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
		}
	};
};

export default buildStateChanges;
