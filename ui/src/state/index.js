import buildStateChanges from './stateChanges';
import buildOperations from './operations';
import { HEADER, LOADING, LOGIN, EXPENDITURES, NAVIGATION, HOME, CATEGORIES, REPORTS } from '../constants';

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
		month: new Date().getMonth(), // zero-indexed month integer
	},
	[NAVIGATION]: {
		current: LOADING,
		default: HOME,
	},
	[HEADER]: {
		selectedYear: '2024',
		years: ['2024','2025', '2026', '2027', '2028'],
	},
	[REPORTS]: {
		report: null,
		priorReport: null,
	}
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

