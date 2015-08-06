'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('SessionController', ['$scope', '$location', '$routeParams', 'sessionService', function($scope, $location, $routeParams, sessionService){
	// console.log($routeParams.item);
	$scope.menuItem = $routeParams.item;
	$scope.curUser = {};
	$scope.curMenu = '菜单1';
	$scope.curUser.userName = sessionService.getItem('userName');
	$scope.showLogin = function () {
		$location.path('/login');
	}
	$scope.userMenus = [{name: '菜单1', ico: 'icon-home', href: '#/menu/11'}, {name: '菜单2', ico: 'icon-flag', href: '#/menu/12'}, {name: '菜单3', ico: 'icon-earth', href: '#/menu/13'}, {name: '菜单4', ico: 'icon-bookmarks', href: '#/menu/14'}];
	$scope.adminMenus = [{name: '管理菜单1', ico: 'icon-cog', href: '#/menu/21'}, {name: '管理菜单2', ico: 'icon-flag', href: '#/menu/22'}, {name: '管理菜单3', ico: 'icon-pie-chart', href: '#/menu/23'}, {name: '管理菜单4', ico: 'icon-users', href: '#/menu/24'}, {name: '管理菜单5', ico: 'icon-clipboard', href: '#/menu/25'}];
}]);