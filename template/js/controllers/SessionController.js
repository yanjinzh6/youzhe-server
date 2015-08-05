'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('SessionController', ['$scope', '$location', 'sessionService', function($scope, $location, sessionService){
	$scope.curUser = {};
	$scope.curUser.userName = sessionService.getItem('userName');
	$scope.showLogin = function () {
		$location.path('/login');
	}
}]);