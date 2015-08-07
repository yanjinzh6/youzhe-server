'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('SessionController', ['$routeParams', 'sessionService', function($routeParams, sessionService){
	var vm = this;
	// console.log($routeParams.item);
	vm.menuItem = $routeParams.item;
	vm.curMenu = '菜单1';
	vm.curUser = angular.fromJson(sessionService.getItem('user'));
	console.log(vm.curUser);
	return vm;
}]);