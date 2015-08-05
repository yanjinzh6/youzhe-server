'use strict';

/**
* yozh.controllers Module
*
* Description
*/
angular.module('yozh.controllers', []).

controller('HeaderController', ['$scope', function($scope){
	console.log($scope);
	// $scope.name = 'header';
	$scope.slide = {};
	$scope.adminMenu = {};
	$scope.slide.isOpen = false;
	$scope.adminMenu.isOpen = false;

	$scope.slide.openSlide = function () {
		// body...
		$scope.slide.isOpen = true;
	};
	$scope.slide.closeSlide = function () {
		// body...
		$scope.slide.isOpen = false;
	};

	$scope.adminMenu.changeMenu = function () {
		// body...
		$scope.adminMenu.isOpen = !$scope.adminMenu.isOpen;
	};
}]);