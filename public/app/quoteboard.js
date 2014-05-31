var app = angular.module("quoteboard", [
  "firebase", 
  "ui.router",
  "quoteboard.services",
  "quoteboard.directives",
  "quoteboard.controllers",
  "angularjs-gravatardirective"
  ]);
  
app.run(function($rootScope, $firebaseSimpleLogin, messageService) {
  var ref = new Firebase(FBURL);
  $rootScope.$on("$firebaseSimpleLogin:login", function(e, user) {
    $rootScope.user = user;
  });
  $rootScope.$on("$firebaseSimpleLogin:logout", function(e) {
    $rootScope.user = null;
  });
  $rootScope.$on("$firebaseSimpleLogin:error", function(e, error) {
    errorMessage = error.message.replace("FirebaseSimpleLogin: FirebaseSimpleLogin: ", "");
    messageService.showError(errorMessage);
  });

  $rootScope.auth = $firebaseSimpleLogin(ref);
});

app.config(function($stateProvider, $urlRouterProvider) {
  $urlRouterProvider
    .when("", "/");

  $stateProvider
    .state('index', {
      url: "/",
      templateUrl: 'views/home.html'
    }).state('quote', {
      url: '/quotes/:quoteId',
      templateUrl: 'views/quote.html',
      controller: 'QuoteController'
    }).state('notfound', {
      url: '{path:.*}',
      templateUrl: 'views/404.html'
    });
});