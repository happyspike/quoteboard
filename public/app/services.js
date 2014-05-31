(function() {
  var module = angular.module("quoteboard.services", []);

  module.factory('messageService',function($rootScope, $timeout) {
    return {
      showError: function(errorMessage) {
        $rootScope.errors.push(errorMessage);
        $timeout(function () {
          $rootScope.errors.shift();
        }, 3000)

      }
    };
  });

})();