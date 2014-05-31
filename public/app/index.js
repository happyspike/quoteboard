var app = angular.module("quoteboard", ["firebase", "ui.router", "controllers"]);
app.config(function($stateProvider, $urlRouterProvider) {
  $urlRouterProvider
    .when("", "/");

  $stateProvider
    .state('index', {
      url: "/",
      templateUrl: 'views/index.html',
      controller: 'IndexController'
    });
});