angular.module('cfshop')
    .directive('headerComponent', function () {
        return {
            templateUrl: function () {
                return window.config.baseLocation + '/header.html';
            },
            scope: {
                login: "="
            },
            controller: ['$scope', function ($scope) {
                $scope.isNavbarOpen = false;

                $scope.toggleNavbar = function () {
                    $scope.isNavbarOpen = !$scope.isNavbarOpen;
                };
            }]
        };
    });
