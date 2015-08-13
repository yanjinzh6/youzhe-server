'use strict';
/**
* yozh.services Module
*
* Description
*/
angular.module('yozh.services').

factory('homeService', ['$http', function($http){
	return {
		queryInfo: function() {
			return $http({
				method: 'GET',
				url: '/test/123'
			});
		}
	};
}]);