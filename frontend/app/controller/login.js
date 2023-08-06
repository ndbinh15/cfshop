angular.module('cfshop')
    .controller('loginController', ['$scope', '$http', '$window','$timeout', function ($scope, $http, $window, $timeout) {
        $scope.login = function () {
            $scope.login = false;
            let username = $scope.username;
            let password = $scope.password;

            // let hashedPassword = hashPassword(password);
            let hashedPassword = password;

            $http
                .post('/login/login', { username: username, password: hashedPassword })
                .then(function (response) {
                    console.log(response);
                    if (response.status == 200) {
                        $scope.success = 'Login Success';
                        $scope.error = null;
                        $scope.login = true;
                        $timeout(function () {
                            $window.location.href = '/home'
                        }, 1000);

                    } else {
                        $scope.error = 'Invalid Login';
                        $timeout(function () {
                            $window.location.href = '/login'
                        }, 1000);
                    }
                })
                .catch(function (error) {
                    // Display error message on login page
                    $scope.error = 'Invalid Login';
                    $timeout(function () {
                        $window.location.href = '/login'
                    }, 1000);
                });
        };
        function hashPassword(password) {
            const bcrypt = dcodeIO.bcrypt;
            const saltRounds = 10;
            const salt = bcrypt.genSaltSync(saltRounds);
            const hashedPassword = bcrypt.hashSync(password, salt);
            console.log('hashedPassword', hashedPassword)
            return hashedPassword;
        }
    }])