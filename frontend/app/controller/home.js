angular.module('cfshop')
    .controller('homeController', ['$rootScope', '$scope', '$http', '$window', 'dataService',
        function ($rootScope, $scope, $http, $window, dataService) {
            $scope.linkHomeAdmin = "/home/admin"
            // $scope.linkHomeUser = "/home/user"
            $scope.linkHomeUser = "/home/user"

            $scope.isAuth = false;
            if (dataService.getData('userID') == undefined || dataService.getData('userID') == null) {
                $window.location.href = '/login'
            } else {
                $scope.isAuth = true;
            }
            //$scope.user = {};
            //$scope.checkRole = function () {
            //    var userId = $scope.user.id;

            //    $http.get("/check-role?id=" + userId)
            //        .then(function (response) {
            //            if (response.data?.isAdmin == true) {
            //                $scope.isAdmin = true;
            //            } else if (response.data?.isUser == true) {
            //                $scope.isUser = false;
            //            } else {
            //                $scope.isAdmin = false;
            //                $scope.isUser = false;
            //            }
            //        }, function (error) {
            //            alert(error.message);
            //        });
            //};
        }])
