import buildStateChanges from './stateChanges';
import buildOperations from './operations';
import { LOADING, LOGIN, EXPENDITURES, NAVIGATION, HOME, CATEGORIES } from '../constants';

const INITIAL_STATE = {
	[LOGIN]: {
		email: '',
		password: '',
		isLoggedIn: false,
		user: '',
		userDisplayName: '',
	},
	[CATEGORIES]: {
		categories: [],
		selectionsMap: {}, // categorize specific expenditures
		broadSelectionsMap: {}, // categorize all expenditures w/ a description
		expenditureNames: [],
		expenditureNamePage: 1,
		error: null,
	},
	[EXPENDITURES]: {
		expenditures: [],
		month: new Date().getMonth(), // zero-indexed month integer,
	},
	[NAVIGATION]: {
		current: LOADING,
		default: HOME,
	},
};

const getInitialState = () => INITIAL_STATE;

const buildStateManager = (state, setState) => {

	const stateChanges = buildStateChanges(state, setState, getInitialState);
	const operations = buildOperations(state, stateChanges);
	return {
		state,
		ops: operations,
		changes: stateChanges, // the ideal might be to not expose this at all
		getInitialState,
	};
};

const StateStore =  {
	INITIAL_STATE,
	buildStateManager,
};

export default StateStore;

