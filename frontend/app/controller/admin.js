angular.module('cfshop')
    .controller('adminController', ['$scope', '$http', function ($scope, $http) {

        //link page
        $scope.linkAdminProducts = "/home/admin/products"
        $scope.linkAdminUsers = "/home/admin/users"
        $scope.linkAdminOrders = "/home/admin/orders"
        $scope.linkAdminHome = "/home/admin"

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