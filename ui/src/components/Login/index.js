import React from 'react';

const Login = ({ stateManager }) => {
	
	const {
		changes: {
			handleSubmitLogin,
			handleEmailType,
			handlePasswordType,
		}
	} = stateManager;

	return (
		<div>
			Login To Check Your Expenses
			<form onSubmit={handleSubmitLogin}>
				<input type="text" placeholder="username" onChange={handleEmailType} value={stateManager.state.login.email} />
				<input type="password" placeholder="password" onChange={handlePasswordType} value={stateManager.state.login.password} />
				<input type="submit" value="Login" />
			</form>
		</div>
	);
};

export default Login;
