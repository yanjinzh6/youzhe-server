'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('SessionController', ['$routeParams', '$location', 'sessionService', function($routeParams, $location, sessionService){
	var vm = this;
	// console.log($routeParams.item);
	// vm.menuItem = $routeParams.item;
	vm.curMenu = '菜单1';
	vm.curUser = angular.fromJson(sessionService.getItem('user'));
	console.log(vm.curUser);
	if (typeof(vm.curUser) == 'undefined' || vm.curUser == null) {
		vm.curUser = {};
		vm.curUser.name = '未登录';
		vm.curUser.islogin = false;
		vm.curUser.userMenus = [{name: '菜单1', ico: 'icon-home', href: '/menu/11'}, {name: '菜单2', ico: 'icon-flag', href: '/menu/12'}, {name: '菜单3', ico: 'icon-earth', href: '/menu/13'}, {name: '菜单4', ico: 'icon-bookmarks', href: '/menu/14'}];
	}
	vm.login = function () {
		$location.path('/login');
	}
	return vm;
}]);