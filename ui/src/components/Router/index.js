import React from 'react';
import { NAVIGATION } from '../../constants';

const Router = ({ children, stateManager } ) => {
	const { state: { [NAVIGATION]: { current } } } = stateManager;
	return children.find(({ props: { navigation } }) => navigation === current) ?? null;
};

export default Router;
