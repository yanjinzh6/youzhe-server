'use strict';
/**
* yozh.services Module
*
* Description
*/
angular.module('yozh.services').

factory('userService', ['$resource', function($resource){
	var users = $resource('/users/:userId', {userId: '@userId'}, {
		checkEmail: function() {
			method: 'GET'
		}
	});
	return users;
}])