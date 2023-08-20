angular.module('cfshop')
    .controller('userProfileController', ['$scope', '$http', 'dataService', function ($scope, $http, dataService) {

        //link page /home/user/
        $scope.linkUserProfile = "/home/user/profile"
        $scope.linkUserCart = "/commingsoon"
        $scope.linkUserBuy = "/home/user/listproduct"
        $scope.linkUserOrdered = "/commingsoon"
        $scope.linkUserTracking = "/commingsoon"

        $scope.getUserInfo = function (userID) {
            $http.get('/users?id=' + userID)
                .then(function (response) {
                    console.log(response);
                    $scope.userInfo = response.data;
                    //$scope.updatedUser = angular.copy($scope.updateUser);
                })
                .catch(function (error) {
                    console.error('Error retrieving products:', error);
                });
        }
        $scope.getUserInfo(dataService.getData('userID'));
    }])
