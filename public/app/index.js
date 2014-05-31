var app = angular.module("quoteboard", ["firebase", "ui.router", "controllers"]);
app.config(function($stateProvider, $urlRouterProvider) {
  $urlRouterProvider
    .when("", "/");

  $stateProvider
    .state('index', {
      url: "/",
      templateUrl: 'views/index.html',
      controller: 'IndexController'
    }).state('quote', {
      url: '/quotes/:quoteId',
      templateUrl: 'views/quote.html',
      controller: 'QuoteController'
    }).state('notfound', {
      url: '{path:.*}',
      templateUrl: 'views/404.html'
    });
});