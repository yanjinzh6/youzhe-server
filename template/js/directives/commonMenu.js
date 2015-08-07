'use strict';

/**
* yozh.directives Module
*
* Description
*/
angular.module('yozh.directives', []).

directive('menuItem', ['$location', function($location){
	// Runs during compile
	return {
		// name: '',
		// priority: 1,
		// terminal: true,
		// scope: {}, // {} = isolate, true = child, false/undefined = no change
		// controller: function($scope, $element, $attrs, $transclude) {},
		// require: 'ngModel', // Array = multiple requires, ? = optional, ^ = check parent elements
		restrict: 'A', // E = Element, A = Attribute, C = Class, M = Comment
		template: '<li ng-click="gotoView()" ng-class="{active: isActive}"><a ng-href="#[[menu.href]]" class=""><span class="[[menu.ico]]"></span>[[menu.name]]</a></li>',
		// templateUrl: '',
		replace: true,
		transclude: true,
		// compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
		link: function($scope, iElm, iAttrs, controller) {
			$scope.isActive = false;
			console.log($scope);
			// console.log($location.path(), $scope.menu.href);
			if ($location.path() == $scope.menu.href) {
				$scope.isActive = true;
			}
			$scope.$watch(function() {
				return $location.path();
			}, function(value) {
				console.log(value);
				if (value == $scope.menu.href) {
					$scope.isActive = true;
				} else {
					$scope.isActive = false;
				}
			});
			// $scope.$on('$locationChangeSuccess', function(e) {
			// 	console.log(e);
			// 	if ($location.path() == $scope.menu.href) {
			// 		$scope.isActive = true;
			// 	} else {
			// 		$scope.isActive = false;
			// 	}
			// });
		}
	};
}]);