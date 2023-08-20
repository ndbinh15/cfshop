angular.module('cfshop')
    .controller('loginController', ['$rootScope', '$scope', '$http', '$window', '$timeout', '$location','dataService',
        function ($rootScope, $scope, $http, $window, $timeout, $location, dataService) {
            $scope.login = function () {
                //$scope.login = false;
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
                            //$scope.login = true;
                            dataService.setData('userID', response.data.ID);
                            dataService.setData('userRole', response.data.RL);

                            $timeout(function () {
                                $window.location.href = '/home';
                                //$location.path('/home')
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

            //$rootScope.isAdmin = function () {
            //    return $scope.user.role == "admin";
            //};

            //$rootScope.isUser = function () {
            //    return $scope.user.role == "user";
            //};
        }])