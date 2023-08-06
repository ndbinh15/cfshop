angular.module('cfshop')
    .controller('welcomeController', ['$scope', '$http', function ($scope, $http) {

        //link page /home/user/
        $scope.linkSignUp = "/signup"
        $scope.linkSignIn = "/login"
    }])
