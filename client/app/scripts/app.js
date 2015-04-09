'use strict';

/**
 * @ngdoc overview
 * @name clientApp
 * @description
 * # clientApp
 *
 * Main module of the application.
 */
angular
  .module('clientApp', [
    'ngAnimate',
    'ngRoute'
  ])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/about', {
        templateUrl: 'views/about.html',
        controller: 'AboutCtrl'
      })
      .when('/device', {
        templateUrl: 'views/device.html',
        controller: 'DeviceCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  })
  .controller('AppCtrl', function ($scope,EXI) {

    $scope.exi = {
      devices: null,
    };
    
    EXI.getDevices()
    .then(function (data) {
      $scope.exi.devices = data.data;
      console.log('devices:', $scope.exi.devices);
    });
  })
  ;
