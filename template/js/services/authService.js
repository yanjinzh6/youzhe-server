'use strict';
/**
* yozh.services Module
*
* Description
*/
angular.module('yozh.services').

factory('authService', ['$resource', function($resource){
	var user = $resource('/test/:userId', {userId: '@userId'}, {
		myAction: {
			method: 'GET'
		}
	});
	return user;
}])