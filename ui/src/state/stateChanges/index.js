import { LOGIN, NAVIGATION, HOME } from '../../constants';

const buildStateChanges = (state, setState, getInitialState) => {

	const update = newState => setState(oldState => ({ ...oldState, ...newState }));

	const navigate = (location) => {
		const newState = {
			...state,
			[NAVIGATION] : {
				...state[NAVIGATION],
				current: location,
			},
		}
		update(newState);
	};

	return {
		[LOGIN]: {
			handleEmailType: (e) => {
				const { target: { value } } = e;
				const newState = { ...state, login: { ...state.login, email: value } };
				update(newState);
			},
			handlePasswordType: (e) => {
				const { target: { value } } = e;
				const newState = { ...state, login: { ...state.login, password: value } };
				update(newState);
			},
			handleLoginSuccess: (successPayload) => {
				const newState = { ...state, login: { ...state.login, isLoggedIn: true } };
				update(newState);
				navigate(HOME);
			},
			handleLoginFailure: (failurePayload) => {
				const newState = { ...state, login: { ...state.login, isLoggedIn: false } };
				update(newState);
			},
			handleLoggedIn: (successPayload) => {
				console.log('successPayload', successPayload);
				const newState = { ...state, login: { ...state.login, isLoggedIn: true } };
				update(newState);
				navigate(HOME);
			},
			handleNotLoggedIn: (failurePayload) => {
				console.log('failurePayload', failurePayload);
				const newState = { ...state, login: { ...state.login, isLoggedIn: false } };
				update(newState);
				navigate(LOGIN);
			},
			handleLogoutSuccess: (successPayload) => {
				console.log('successPayload', successPayload);
				const newState = {
					...state,
					login: {
						...getInitialState()[LOGIN],
						isLoggedIn: false,
					},
				};
				update(newState);
				navigate(LOGIN);
			},
			handleLogoutFailure: (failurePayload) => {
				console.log('failurePayload', failurePayload);
				const newState = {
					...state,
					login: {
						...getInitialState()[LOGIN],
						isLoggedIn: false,
					},
				};
				update(newState);
				navigate(LOGIN);
			},
		},
		[NAVIGATION]: {
			navigate,
		},
	};
};

export default buildStateChanges;
