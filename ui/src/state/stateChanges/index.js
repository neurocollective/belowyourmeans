const buildStateChanges = (state, setState) => {

	const update = newState => setState(oldState => ({ ...oldState, ...newState }));

	return {

		handleSubmitLogin: (e) => {
			e.preventDefault();
			console.log('SUBMITTING BRAH');
			const {
				login: {
					email, password
				}
			} = state;
			console.log('email', email);
			console.log('password', password);
			// const newState = {
			// 	...state,

			// }
			// update(newState);
		},
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
	};
};

const stateChanges = {
	buildStateChanges,
};

export default stateChanges;
