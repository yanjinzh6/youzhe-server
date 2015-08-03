'use strict';

/**
* yozh Module
*
* Description
*/
angular.module('yozh', ['ngRoute', 'yozh.header']).

config(['$routeProvider', function($routeProvider) {
	console.log($routeProvider);
	// $routeProvider.otherwise({redirectTo: '/'});
}]);