angular.module('cfshop')
    .controller('userController', ['$scope', '$http', function ($scope, $http) {

        //link page /home/user/
        $scope.linkUserProfile = "/commingsoon"
        $scope.linkUserCart = "/commingsoon"
        $scope.linkUserBuy = "/home/user/buyproduct"
        $scope.linkUserOrdered = "/commingsoon"
        $scope.linkUserTracking = "/commingsoon"
    }])
