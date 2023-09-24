import React from 'react';

const Login = ({ stateManager }) => {
	
	const {
		handleSubmitLogin,
		handleEmailType,
		handlePasswordType,
	} = stateManager;

	return (
		<div>
			<form onSubmit={handleSubmitLogin}>
				<input type="email" onChange={handleEmailType} />
				<input type="password" onChange={handlePasswordType} />
				<input type="submit" />
			</form>
		</div>
	);
};

export default Login;
