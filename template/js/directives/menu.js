'use strict';

/**
* yozh.directives Module
*
* Description
*/
angular.module('yozh.directives', []).

directive('commonMenu', function(){
	// Runs during compile
	return {
		// name: '',
		// priority: 1,
		// terminal: true,
		// scope: {
		// 	menus: '=nMenus'
		// }, // {} = isolate, true = child, false/undefined = no change
		controller: function($scope, $element, $attrs, $transclude) {
			// console.log($scope);
			var items = [];
			this.addItem = function (item) {
				items.push(item);
			}
			this.changeActive = function (curItem) {
				angular.forEach(items, function (item) {
					if (curItem != item) {
						item.isActive = false;
					}
				});
			}
		},
		// require: 'ngModel', // Array = multiple requires, ? = optional, ^ = check parent elements
		restrict: 'A', // E = Element, A = Attribute, C = Class, M = Comment
		// template: '<ul></ul>',
		// templateUrl: '../../views/menu.html',
		// replace: true,
		// transclude: true,
		// compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
		// link: function($scope, iElm, iAttrs, controller) {}
	};
}).

directive('menuItem', function(){
	// Runs during compile
	return {
		// name: '',
		// priority: 1,
		// terminal: true,
		// scope: {}, // {} = isolate, true = child, false/undefined = no change
		// controller: function($scope, $element, $attrs, $transclude) {},
		require: '^?commonMenu', // Array = multiple requires, ? = optional, ^ = check parent elements
		restrict: 'A', // E = Element, A = Attribute, C = Class, M = Comment
		template: '<li ng-click="gotoView()" ng-class="{active: isActive}"><a href="[[menu.href]]" class=""><span class="[[menu.ico]]"></span>[[menu.name]]</a></li>',
		// templateUrl: '',
		replace: true,
		transclude: true,
		// compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
		link: function($scope, iElm, iAttrs, commonMenuController) {
			$scope.isActive = false;
			commonMenuController.addItem($scope);
			$scope.gotoView = function () {
				if (!$scope.isActive) {
					$scope.isActive = true;
					commonMenuController.changeActive($scope);
				}
			}
		}
	};
});