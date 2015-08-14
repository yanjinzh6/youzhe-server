'use strict';
/**
* yozh.services Module
*
* Description
*/
angular.module('yozh.services').

factory('authService', ['$resource', function($resource){
	var user = $resource('/users/:userId', {userId: '@userId'}, {
		queryUser: {
			method: 'GET'
		},
		register: {
			method: 'POST'
		}
	});
	return user;
}])