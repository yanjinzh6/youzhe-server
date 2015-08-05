'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('LoginController', ['$scope', 'sessionService', function($scope, sessionService){
	$scope.user = {};
	$scope.login = function () {
		console.log(user.username, user.password);
		sessionService.setItem('userName', username);
	}
}]);