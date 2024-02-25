import buildStateChanges from './stateChanges';
import buildOperations from './operations';
import { LOGIN,EXPENDITURE } from '../constants';

const buildStateManager = (state, setState) => {

	const stateChanges = buildStateChanges(state, setState);
	const operations = buildOperations(state, stateChanges);
	return {
		state,
		ops: operations,
		changes: stateChanges, // the ideal might be to not expose this at all
	};
};

const StateStore =  {
	INITIAL_STATE: {
		[LOGIN]: {
			email: '',
			password: '',
			isLoggedIn: false,
			user: '',
			userDisplayName: '',
		},
		[EXPENDITURE]: {
			expenditures: [],
			month: new Date().getMonth(), // zero-indexed month integer
		}
	},
	buildStateManager,
};

export default StateStore;

