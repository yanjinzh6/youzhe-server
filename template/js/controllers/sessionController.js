'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('SessionController', ['$scope', '$location', 'sessionService', function($scope, $location, sessionService){
	$scope.curUser = {};
	$scope.curUser.userName = sessionService.getItem('userName');
	$scope.showLogin = function () {
		$location.path('/login');
	}
	$scope.userMenus = [{name: '菜单1', ico: 'icon-home', href: '/###'}, {name: '菜单2', ico: 'icon-home', href: '/###'}, {name: '菜单3', ico: 'icon-home', href: '/###'}, {name: '菜单4', ico: 'icon-home', href: '/###'}];
	$scope.adminMenus = [{name: '管理菜单1', ico: 'icon-flag', href: '/###'}, {name: '管理菜单2', ico: 'icon-flag', href: '/###'}, {name: '管理菜单3', ico: 'icon-flag', href: '/###'}, {name: '管理菜单4', ico: 'icon-flag', href: '/###'}, {name: '管理菜单5', ico: 'icon-flag', href: '/###'}];
}]);