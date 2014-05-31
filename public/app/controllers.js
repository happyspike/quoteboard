var module = angular.module("controllers", ["firebase", "ui.router"]);

module.controller('IndexController', function($scope, $firebase) {
  var ref = new Firebase(FBURL + '/quotes');
  
  $scope.quotes = $firebase(ref);
  $scope.addQuote = function() {
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