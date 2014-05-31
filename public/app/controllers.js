(function() {
  var module = angular.module("controllers", ["firebase", "ui.router"]);

  module.controller('IndexController', function($rootScope, $scope, $firebase) {
    var ref = new Firebase(FBURL + '/quotes');
    $scope.quotes = $firebase(ref);
    $scope.addQuote = function() {
      if (!$scope.newQuote) { return; }
      $scope.newQuote.user = $rootScope.user.email;
      $scope.newQuote.added = Firebase.ServerValue.TIMESTAMP;
      $scope.quotes.$add($scope.newQuote);
      $scope.newQuote = {};
      document.getElementById("newQuoteContent").focus();
    }
  });

  module.controller('QuoteController', function($scope, $firebase, $stateParams) {
    var ref = new Firebase(FBURL + '/quotes/' + $stateParams.quoteId);
    $scope.quote = $firebase(ref);
  });

  module.controller('SessionController', function($scope, $rootScope, $firebaseSimpleLogin) {
    $scope.state = 'login';
    $scope.showRegister = function() {
      $scope.state = 'register';
    };
    $scope.showLogin = function() {
      $scope.state = 'login';
    };
    $scope.register = function() {
      if (!$scope.register.email) { return; }
      if (!$scope.register.password) { return; }
      if ($scope.register.password != $scope.register.confirmation) { return; }

      $rootScope.auth.$createUser($scope.register.email, $scope.register.password).then(function(user) {
        if (user) { 
          $rootScope.auth.$login('password', {
            email: $scope.register.email,
            password: $scope.register.password,
            rememberMe: true
          });
        }
      }).then(function () {
        $scope.register = null;
      });
    };
    $scope.login = function() {
      if (!$scope.login.email) { return; }
      if (!$scope.login.password) { return; }

      $rootScope.auth.$login('password', {
        email: $scope.login.email,
        password: $scope.login.password,
        rememberMe: true
      }).then(function () {
        $scope.login = null;
      });
    };
  });
})();