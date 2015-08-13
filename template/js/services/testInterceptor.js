'use strict';
/**
* yozh.services Module
*
* Description
*/
angular.module('yozh.services').

factory('testInterceptor', ['$q', function($q){
	var interceptor = {
		'request': function(config) {
			// request success
			console.log(config);
			return config; // or $q.when(config)
		},
		'response': function(response) {
			// response success
			console.log(response);
			return response; 
		},
		'requestError': function(rejection) {
			// request error, if can recovery from error, you can return a new request or promise
			console.log(rejection);
			return rejection; // or new promise;
			// or return a rejection to stop
			// return $q.reject(rejection);
		},
		'responseError': function(rejection) {
			//response error
			console.log(rejection);
			return rejection;
		}
	}
	return interceptor;
}])