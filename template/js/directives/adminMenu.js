'use strict';

/**
* yozh.directives Module
*
* Description
*/
angular.module('yozh.directives', []).

directive('adminMenu', function(){
	// Runs during compile
	return {
		// name: '',
		// priority: 1,
		// terminal: true,
		// scope: {}, // {} = isolate, true = child, false/undefined = no change
		// controller: function($scope, $element, $attrs, $transclude) {},
		// require: 'ngModel', // Array = multiple requires, ? = optional, ^ = check parent elements
		restrict: 'A', // E = Element, A = Attribute, C = Class, M = Comment
		// template: '',
		templateUrl: '../../views/admin_menu.html',
		// replace: true,
		// transclude: true,
		// compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
		link: function($scope, iElm, iAttrs, controller) {
			console.log($scope, iElm, iAttrs, controller);
			$scope.menus = [{name: '管理菜单1', ico: 'icon-flag', href: '/###'}, {name: '管理菜单2', ico: 'icon-flag', href: '/###'}, {name: '管理菜单3', ico: 'icon-flag', href: '/###'}, {name: '管理菜单4', ico: 'icon-flag', href: '/###'}, {name: '管理菜单5', ico: 'icon-flag', href: '/###'}];
		}
	};
});