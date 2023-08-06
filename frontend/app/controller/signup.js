angular.module('cfshop')
    .controller('signupController', ['$scope', '$http', '$window', '$timeout', function ($scope, $http, $window, $timeout) {
        // add user
        $scope.inValidPassword = false;
        $scope.successNoti = false;
        $scope.errorNoti = false;
        $scope.user = {};

        $scope.register = function () {
            if ($scope.user.confirmPassword !== $scope.user.password) {
                $scope.inValidPassword = true;
                $scope.user.confirmPassword = "";
            } else {
                $scope.inValidPassword = false;

                $http.post('/user/register', $scope.user)
                    .then(function (response) {
                        console.log(response);
                        if (response.data && response.data.success) {
                            $scope.successMess = "User register successfully";
                            $scope.user = {}; // Clear the form fields
                            $scope.successNoti = true;
                            $timeout(function () {
                                $window.location.href = '/login'
                            }, 2000);

                        } else {
                            $scope.errorNoti = true;
                            $scope.errorMess = "Failed to register User:" + response.data.message;
                        }
                    })
                    .catch(function (error) {
                        console.log(error)
                        $scope.errorNoti = true;
                        $scope.errorMess = 'Error register user:' + error;
                    });
                //$timeout(function () {
                //    $scope.successNoti = false;
                //    $scope.errorNoti = false;
                //}, 3000);
            }

        };

    }])