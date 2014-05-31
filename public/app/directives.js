(function() {

  var module = angular.module("directives", ['controllers']);
  module.directive('sessionElement', function () {
    return {
      restrict: 'E',
      templateUrl: 'views/partials/session.html',
      controller: 'SessionController'
    };
  })

})();
