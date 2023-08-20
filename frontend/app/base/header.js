angular.module('cfshop')
    .directive('headerComponent', function () {
        return {
            templateUrl: function () {
                return window.config.baseLocation + '/header.html';
            },
            scope: {
                login: "="
            },
            controller: ['$scope','$window', 'dataService', function ($scope,$window, dataService) {
                $scope.isNavbarOpen = false;

                $scope.toggleNavbar = function () {
                    $scope.isNavbarOpen = !$scope.isNavbarOpen;
                };
                $scope.logout = function () {
                    dataService.clearData();
                    $window.location.href = '/login';
                }
                $scope.role = dataService.getData('userRole');
                console.log($scope.role)
            }]
        };
    });
