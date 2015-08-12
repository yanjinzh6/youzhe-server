'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers').

controller('LoginController', ['$scope', 'sessionService', function($scope, sessionService){
	$scope.user = {};
	$scope.login = function() {
		console.log($scope);
		if ($scope.sessionCtrl.curUser == null) {$scope.sessionCtrl.curUser = {}};
		$scope.sessionCtrl.curUser.islogin = true;
		$scope.sessionCtrl.curUser.name = 'yanjin';
		$scope.sessionCtrl.curUser.token =  '1212323343';
		$scope.sessionCtrl.curUser.userMenus = [{name: '菜单1', ico: 'icon-home', href: '/menu/11'}, {name: '菜单2', ico: 'icon-flag', href: '/menu/12'}, {name: '菜单3', ico: 'icon-earth', href: '/menu/13'}, {name: '菜单4', ico: 'icon-bookmarks', href: '/menu/14'}];
		$scope.sessionCtrl.curUser.adminMenus = [{name: '管理菜单1', ico: 'icon-cog', href: '/menu/21'}, {name: '管理菜单2', ico: 'icon-flag', href: '/menu/22'}, {name: '管理菜单3', ico: 'icon-pie-chart', href: '/menu/23'}, {name: '管理菜单4', ico: 'icon-users', href: '/menu/24'}, {name: '管理菜单5', ico: 'icon-clipboard', href: '/menu/25'}];
		sessionService.setItem('user', angular.toJson($scope.sessionCtrl.curUser));
	}
	$scope.logout = function() {
		$scope.sessionCtrl.curUser.islogin = false;
		sessionService.setItem('user', angular.toJson($scope.sessionCtrl.curUser));
	}
}]);