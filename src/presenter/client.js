import HTTP from 'marbles/http';
import JSONMiddleware from 'marbles/http/middleware/serialize_json';

var Client = {
	getSlides: function () {
		return HTTP({
			method: 'GET',
			url: '/api/slides',
			headers: {
				Accept: 'application/json'
			},
			middleware: [JSONMiddleware]
		}).then(function (args) {
			return args[0];
		});
	}
};

export default Client;
