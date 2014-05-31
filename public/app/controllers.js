var module = angular.module("controllers", ["firebase"]);

module.controller('IndexController', ['$scope', '$firebase', function($scope, $firebase) {
  var ref = new Firebase("https://quoteboard.firebaseio.com/");
  $scope.newQuote = {};
  $scope.quotes = $firebase(ref);
  
  $scope.addQuote = function() {
    $scope.quotes.$add($scope.newQuote);
    $scope.newQuote = {};
    document.getElementById("newQuoteContent").focus();
  }
}]);