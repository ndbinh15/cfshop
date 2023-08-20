angular.module('cfshop')
    .controller('userController', ['$scope', '$http', function ($scope, $http) {

        //link page /home/user/
        $scope.linkUserProfile = "/home/user/profile"
        $scope.linkUserCart = "/home/user/cart"
        $scope.linkUserBuy = "/home/user/listproduct"
        $scope.linkUserOrdered = "/home/user/ordered"
        $scope.linkUserTracking = "/commingsoon"

    }])
