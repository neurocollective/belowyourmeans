// config contents -> https://developer.mozilla.org/en-US/docs/Web/API/fetch
const jsonRequest = (url, config, success, failure) => fetch(url, config)
	.then((response) => {
		return response.json();
	}).then((jsonObject) => {
		return success(jsonObject);
	}).catch((error) => {
		return failure(error);
	});

export default jsonRequest;
