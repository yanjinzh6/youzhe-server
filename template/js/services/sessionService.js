'use strict';
/**
* yozh.services Module
*
* Description
*/
angular.module('yozh.services', []).

factory('sessionService', ['$window', function($window){
	var sessionService = {};
	sessionService.setItem = function (key, value) {
		return $window.sessionStorage.setItem(key, value);
	}
	sessionService.getItem = function (key) {
		return $window.sessionStorage.getItem(key);
	}
	sessionService.size = function () {
		return $window.sessionStorage.length;
	}
	sessionService.removeItem = function (key) {
		return $window.sessionStorage.removeItem(key);
	}
	sessionService.clear = function () {
		return $window.sessionStorage.clear();
	}
	return sessionService;
}]);