(function() {

  var module = angular.module("quoteboard.directives", ['quoteboard.controllers']);
  module.directive('sessionElement', function () {
    return {
      restrict: 'E',
      templateUrl: 'views/partials/session.html',
      controller: 'SessionController'
    };
  })

  module.directive('quoteformElement', function () {
    return {
      restrict: 'E',
      templateUrl: 'views/partials/quote.form.html',
      controller: 'QuoteFormController'
    };
  })

  module.directive('quotelistElement', function () {
    return {
      restrict: 'E',
      templateUrl: 'views/partials/quote.list.html',
      controller: 'QuoteListController'
    };
  })

  module.directive('errorsElement', function ($rootScope) {
    $rootScope.errors = [];
    return {
      restrict: 'E',
      templateUrl: 'views/partials/errors.html'
    };
  })

})();
