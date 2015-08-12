'use strict';

/**
* yozh.directives Module
*
* Description
*/
angular.module('yozh.directives').

directive('userInfo', ['$location', function($location){
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
		templateUrl: '../../views/userInfo.html',
		// replace: true,
		// transclude: true,
		// compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
		link: function($scope, iElm, iAttrs, controller) {
			$scope.openAdminMenu = function () {
				if ($scope.sessionCtrl.curUser.islogin) {
					$scope.adminMenu.changeMenu();
				} else {
					$scope.sessionCtrl.login();
				}
			}
		}
	};
}]);