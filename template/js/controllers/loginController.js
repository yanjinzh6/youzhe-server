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
		console.log($scope.user.username, $scope.user.password);
		sessionService.setItem('userName', $scope.user.username);
	}
}]);