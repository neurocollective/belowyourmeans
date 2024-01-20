import React from 'react';
import { LOGIN } from '../../constants';

const Login = ({ stateManager }) => {
	
	const {
		state: {
			[LOGIN]: {
				email,
				password,
			}
		},
		ops: {
			[LOGIN]: {
				handleLoginSubmit,
			},
		},
		changes: {
			[LOGIN]: {
				handleEmailType,
				handlePasswordType,
			}
		},
	} = stateManager;

	console.log("email:", email);
	console.log("password:", password);

	return (
		<div>
			Login To Check Your Expenses
			<form onSubmit={handleLoginSubmit}>
				<input type="text" placeholder="username" onChange={handleEmailType} value={email} />
				<input type="password" placeholder="password" onChange={handlePasswordType} value={password} />
				<input type="submit" value="Login" />
			</form>
		</div>
	);
};

export default Login;
