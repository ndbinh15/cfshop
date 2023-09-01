angular.module('cfshop')
    .controller('adminController', ['$scope', '$http', '$window', 'dataService', function ($scope, $http, $window, dataService) {

        //link page
        $scope.linkAdminProducts = "/home/admin/products"
        $scope.linkAdminUsers = "/home/admin/users"
        $scope.linkAdminOrders = "/home/admin/orders"
        $scope.linkAdminHome = "/home/admin"


        $scope.isAuth = false;
        if (dataService.getData('userRole') == undefined || dataService.getData('userRole') == null || dataService.getData('userRole') == 'user') {
            $window.location.href = '/home'
        } else {
            $scope.isAuth = true;
        }
        //
        const xValues = ["Italy", "France", "Spain", "USA", "Argentina"];
        const yValues = [55, 49, 44, 24, 55];
        const barColors = ["red", "green", "blue", "orange", "brown"];

        new Chart("myChart", {
            type: "bar",
            data: {
                labels: xValues,
                datasets: [{
                    backgroundColor: barColors,
                    data: yValues
                }]
            },
            options: {
                legend: { display: false },
                title: {
                    display: true,
                    text: "World Wine Production 2018"
                }
            }
        });

        //count product
        $scope.countProducts = 0;
        $scope.countProducts = function () {
            $http.get('/products/count')
                .then(function (response) {
                    $scope.countProducts = response.data;

                })
                .catch(function (error) {
                    console.error('Error retrieving products:', error);
                });
        };
        $scope.countProducts();

        //count user
        $scope.countUser = 0;
        $scope.countUsers = function () {
            $http.get('/users/count')
                .then(function (response) {
                    $scope.countUser = response.data;

                })
                .catch(function (error) {
                    console.error('Error retrieving users:', error);
                });
        };
        $scope.countUsers();

        //count order inprogress
        $scope.countOrderInprogress = 0;
        $scope.countOrderInprogress = function () {
            $http.get('/orders/countInprogress')
                .then(function (response) {
                    $scope.countOrderInprogress = response.data;

                })
                .catch(function (error) {
                    console.error('Error retrieving order:', error);
                });
        };
        $scope.countOrderInprogress();

        //count order completed
        $scope.countOrderCompleted = 0;
        $scope.countOrderCompleted = function () {
            $http.get('/orders/countCompleted')
                .then(function (response) {
                    $scope.countOrderCompleted = response.data;

                })
                .catch(function (error) {
                    console.error('Error retrieving order:', error);
                });
        };
        $scope.countOrderCompleted();

        $scope.navChoose = false;
        $scope.contactNav = function () {
            $scope.navChoose = !$scope.navChoose;
            if ($scope.navChoose === true) {
                document.getElementById("mySidenav").style.width = "250px";
                document.getElementById("main").style.marginLeft = "250px";
                document.body.style.backgroundColor = "rgba(0,0,0,0.4)";
            } else {
                document.getElementById("mySidenav").style.width = "0";
                document.getElementById("main").style.marginLeft = "0";
                document.body.style.backgroundColor = "white";
            }
        };
    }]);