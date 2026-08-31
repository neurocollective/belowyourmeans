// config contents -> https://developer.mozilla.org/en-US/docs/Web/API/fetch
const jsonRequest = (url, config, success, failure) => fetch(url, config)
	.then((response) => {

		console.log('jsonRequest got status code', response.status, 'from url', url);
		if (!response.ok) {
			console.log('response NOT OK');
			return response.json().then(json => Promise.reject(json));
		}

		return response.json();
	}).then((jsonObject) => {
		return success(jsonObject);
	}).catch((error) => {
		console.error(`ERROR during \`fetch("${url}", ...)\``);
		console.error(error);
		return failure(error);
	});

export default jsonRequest;
