'use strict';

/**
* yozh Module
*
* Description
*/
angular.module('yozh', ['ngRoute', 'yozh.services', 'yozh.directives', 'yozh.controllers']).

config(['$routeProvider', '$interpolateProvider', function($routeProvider, $interpolateProvider) {
	// console.log($routeProvider);
	$interpolateProvider.startSymbol('[[');
	$interpolateProvider.endSymbol(']]');
	$routeProvider.when('/', {
		controller: 'HomeController',
		templateUrl: '../views/home.html'
	});
	$routeProvider.when('/login', {
		controller: 'LoginController',
		templateUrl: '../views/login.html'
	});
	$routeProvider.when('/menu/:item', {
		controller: 'SessionController',
		template: '[[menuItem]]'
	});
	$routeProvider.otherwise({redirectTo: '/'});
}]);