import React from 'react';
import Nav from '../Nav';

const Header = ({ stateManager }) => {
	return (
		<div>
			<h1>Welcome To BelowYourMeans</h1>
			<Nav stateManager={stateManager} />
		</div>
	);
};

export default Header;
