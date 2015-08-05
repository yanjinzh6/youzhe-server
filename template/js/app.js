'use strict';

/**
* yozh Module
*
* Description
*/
angular.module('yozh', ['ngRoute', 'yozh.services', 'yozh.controllers']).

config(['$routeProvider', '$interpolateProvider', function($routeProvider, $interpolateProvider) {
	// console.log($routeProvider);
	// $routeProvider.otherwise({redirectTo: '/'});
	$interpolateProvider.startSymbol('[[');
	$interpolateProvider.endSymbol(']]');
}]);