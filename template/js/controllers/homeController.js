'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('HomeController', ['$scope', 'homeService', 'authService', function($scope, homeService, authService){
	$scope.home = 'home';
	var p = homeService.queryInfo();
	p.then(function(resp){
		console.log(resp);
	}, function(resp){
		console.log(resp);
	});
	p.success(function(data, status, headers, config){
		console.log(data, status, headers, config);
		$scope.home = data;
	});
	p.error(function(data, status, headers, config){
		console.log(data, status, headers, config);
	});

	var user = authService;
	user.get({userId: 123}, function(resp) {
		console.log(resp);
	}, function(err) {
		console.log(err);
	});
}]);