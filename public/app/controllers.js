(function() {
  var module = angular.module("quoteboard.controllers", ["firebase", "ui.router", "quoteboard.services"]);

  module.controller('QuoteFormController', function($rootScope, $scope, $firebase, messageService) {
    var ref = new Firebase(FBURL + '/quotes');
    $scope.addQuote = function() {
      if (!$rootScope.user) { 
        messageService.showError('Please login first.'); 
        return; 
      }

      if (!($scope.newQuote && $scope.newQuote.content)) { 
        messageService.showError('Quote is empty.'); 
        return; 
      }

      $scope.newQuote.user = $rootScope.user.email;
      $scope.newQuote.added = Firebase.ServerValue.TIMESTAMP;
      $scope.quotes.$add($scope.newQuote);
  
      $scope.newQuote = {};
      document.getElementById("newQuoteContent").focus();
    }
  });

  module.controller('QuoteListController', function($rootScope, $scope, $firebase) {
    var ref = new Firebase(FBURL + '/quotes');
    $scope.quotes = $firebase(ref);
  });

  module.controller('QuoteController', function($scope, $firebase, $stateParams) {
    var ref = new Firebase(FBURL + '/quotes/' + $stateParams.quoteId);
    $scope.quote = $firebase(ref);
  });

  module.controller('SessionController', function($scope, $rootScope, $firebaseSimpleLogin, messageService) {
    $scope.state = 'login';
    $scope.showRegister = function() {
      $scope.state = 'register';
    };
    $scope.showLogin = function() {
      $scope.state = 'login';
    };
    $scope.register = function() {
      if (!$scope.register.email) { 
        messageService.showError('Email is empty.');
        return; 
      }
      if (!$scope.register.password) { 
        messageService.showError('Password is empty.');
        return; 
      }
      if ($scope.register.password != $scope.register.confirmation) { 
        messageService.showError('Passwords do not match.');
        return; 
      }

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
    $scope.logout = function() {
      $rootScope.auth.$logout();
    }
    $scope.login = function() {
      if (!$scope.login.email) { 
        messageService.showError('Email is empty.');
        return; 
      }
      if (!$scope.login.password) { 
        messageService.showError('Password is empty.');
        return; 
      }

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