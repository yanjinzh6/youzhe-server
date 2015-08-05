'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('LoginController', ['$scope', 'sessionService', function($scope, sessionService){
	$scope.curUser = {};
	$scope.curUser.userName = sessionService.getItem('userName');
	$scope.curUser.showLogin = function () {
		console.log('need login');
	}
}]);