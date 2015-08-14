'use strict';

/**
* yozh Module
*
* Description
*/
angular.module('yozh', ['ngRoute', 'ngResource', 'yozh.services', 'yozh.directives', 'yozh.controllers']).

config(['$routeProvider', '$interpolateProvider', '$httpProvider', function($routeProvider, $interpolateProvider, $httpProvider) {
	// console.log($routeProvider);
	$interpolateProvider.startSymbol('[[');
	$interpolateProvider.endSymbol(']]');

	$httpProvider.interceptors.push('testInterceptor');

	$routeProvider.when('/', {
		controller: 'HomeController',
		templateUrl: '../views/home.html'
	});
	$routeProvider.when('/login', {
		controller: 'LoginController',
		templateUrl: '../views/login.html'
	});
	$routeProvider.when('/register', {
		controller: 'LoginController',
		templateUrl: '../views/register.html'
	});
	$routeProvider.when('/menu/:item', {
		controller: 'SessionController as sessionCtrl',
		template: '[[sessionCtrl.menuItem]]'
	});
	$routeProvider.otherwise({redirectTo: '/'});
}]);