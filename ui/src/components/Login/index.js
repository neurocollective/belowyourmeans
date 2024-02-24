import React from 'react';
import { LOGIN } from '../../constants';

const Login = ({ stateManager }) => {
	
	const {
		state: {
			[LOGIN]: {
				email,
				password,
				isLoggedIn,
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

	if (isLoggedIn) {
		return null;
	}

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
