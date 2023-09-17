import reducers from './reducers';

const { buildReducers } = reducers;

const buildStateManager = (state, useState) => {
	return {
		...buildReducers(state, useState),
	};
};

const state =  {
	INITIAL_STATE: {},
	buildStateManager,
};

export default state;

