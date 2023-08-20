angular.module('cfshop')
    .controller('listProductController', ['$scope', '$http', function ($scope, $http) {

        $scope.readProducts = [];
        $scope.getProducts = function () {
            $http.get('/products/get')
                .then(function (response) {
                    console.log(response)
                    $scope.readProducts = response.data;
                    //$scope.totalPages = Math.ceil($scope.readProducts.length / $scope.itemsPerPage);
                    // Assign product numbers to each product
                    //$scope.readProducts.forEach(function (product, index) {
                    //    product.productNumber = index + 1;
                    //});
                    //$scope.updateDisplayedProducts();
                })
                .catch(function (error) {
                    console.error('Error retrieving products:', error);
                });
        };
        $scope.getProducts();
    }])
