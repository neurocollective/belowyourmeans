const buildStateChanges = (state, setState) => {

	const update = (newState) => {
		setState(oldState => ({ ...oldState, ...newState}));
	};

	return {

		handleSubmitLogin: (e) => {
			const {
				login: {
					
				}
			} = state;
			const newState = {
				...state,

			}
			update(newState);
		},
		handleEmailType: (e) => {

		},
		handlePasswordType: (e) => {
			
		},
	};
};

const stateChanges = {
	buildStateChanges,
};

export default stateChanges;
