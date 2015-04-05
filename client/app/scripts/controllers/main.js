'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the clientApp
 */
angular.module('clientApp')
  .controller('MainCtrl', function ($scope,$http) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
    $scope.getStatus = function(url) {
      $http({
	method: 'GET',
	url: url
      })
      .then(function(data){
	console.log('getStatus success:',data);
      },function(data){
	console.log('getStatus error:',data);
      });
    };
    $scope.getStatus('http://192.168.0.15/status');
    $scope.getStatus('http://192.168.0.15/switch');
    $scope.getStatus('http://192.168.0.15/info');
  });
