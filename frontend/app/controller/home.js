angular.module('cfshop')
    .controller('homeController', ['$scope', '$http', function ($scope, $http) {
        $scope.linkHomeAdmin = "/home/admin"
        // $scope.linkHomeUser = "/home/user"
        $scope.linkHomeUser = "/home/user"
    }])
