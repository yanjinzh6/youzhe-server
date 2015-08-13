'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('HomeController', ['$scope', 'homeService', function($scope, homeService){
	$scope.home = 'home';
	homeService.queryInfo().then(function(resp){
		console.log(resp);
	}, function(resp){
		console.log(resp);
	});
	homeService.queryInfo().success(function(data, status, headers, config){
		console.log(data, status, headers, config);
		$scope.home = data;
	});
	homeService.queryInfo().error(function(data, status, headers, config){
		console.log(data, status, headers, config);
	});
}]);