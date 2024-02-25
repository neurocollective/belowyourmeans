import { LOGIN } from '../../constants';

const buildStateChanges = (state, setState) => {

	const update = newState => setState(oldState => ({ ...oldState, ...newState }));

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
			},
			handleLoginFailure: (failurePayload) => {
				const newState = { ...state, login: { ...state.login, isLoggedIn: false } };
				update(newState);
			},
			handleLoggedIn: (successPayload) => {
				console.log('successPayload', successPayload);
				const newState = { ...state, login: { ...state.login, isLoggedIn: true } };
				update(newState);
			},
			handleNotLoggedIn: (failurePayload) => {
				console.log('failurePayload', failurePayload);
				const newState = { ...state, login: { ...state.login, isLoggedIn: false } };
				update(newState);	
			}, 
		}
	};
};

export default buildStateChanges;
