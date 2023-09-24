import stateChanges from './stateChanges';

const { buildStateChanges } = stateChanges;

const buildStateManager = (state, setState) => {

	const stateChanges = buildStateChanges(state, setState);
	const operations = buildOperations(stateChanges);
	return {
		state,
		ops: operations,
		changes: stateChanges, // the ideal might be to not expose this at all
	};
};

const state =  {
	INITIAL_STATE: {
		login: {
			email: '',
			password: '',
			isLoggedIn: false,
			user: '',
			userDisplayName: '',
		}
	},
	buildStateManager,
};

export default state;

