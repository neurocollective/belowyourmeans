import React from 'react';
import Nav from '../Nav';

const Header = ({ stateManager }) => {
	return (
		<header className="App-header">
			<section>
				<h1>Welcome To BelowYourMeans</h1>
				<Nav stateManager={stateManager} />
			</section>
		</header>
	);
};

export default Header;
