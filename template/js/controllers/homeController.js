'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('HomeController', ['$scope', 'homeService', function($scope, homeService){
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
}]);