angular.module('cfshop')
    .controller('productController', ['$scope', '$http', function ($scope, $http) {
        // add product
        $scope.product = {};

        $scope.createProduct = function () {
            $http
                .post('/home/admin/products/add', $scope.product)
                .then(function (response) {
                    console.log(response)
                    if (response.status == 201) {
                        console.log("create success")
                        $scope.product = {};
                    } else {
                        console.log("fail")
                    }
                })
                .catch(function (error) {
                    console.error('Error creating product:', error);
                });
        };

        // show product
        // $scope.readProducts = [];

        // $scope.getProducts = function () {
        //     $http
        //         .get('/home/admin/products/show')
        //         .then(function (response) {
        //             console.log(response.data)
        //             $scope.readProducts = response.data;
        //         })
        //         .catch(function (error) {
        //             console.error('Error retrieving products:', error);
        //         });
        // };

        //
        $scope.readProducts = [];
        $scope.displayedProducts = [];
        $scope.itemsPerPage = 5;
        $scope.currentPage = 1;
        $scope.totalPages = 0;

        $scope.getProducts = function () {
            $http.get('/home/admin/products/show')
                .then(function (response) {
                    console.log(response)
                    $scope.readProducts = response.data;
                    $scope.totalPages = Math.ceil($scope.readProducts.length / $scope.itemsPerPage);
                    // Assign product numbers to each product
                    $scope.readProducts.forEach(function (product, index) {
                        product.productNumber = index + 1;
                    });
                    $scope.updateDisplayedProducts();
                })
                .catch(function (error) {
                    console.error('Error retrieving products:', error);
                });
        };

        $scope.updateDisplayedProducts = function () {
            const startIndex = math.multiply(math.subtract($scope.currentPage, 1), $scope.itemsPerPage);
            console.log(startIndex)
            const endIndex = startIndex + $scope.itemsPerPage;
            $scope.displayedProducts = $scope.readProducts.slice(startIndex, endIndex);
        };

        $scope.previousPage = function () {
            if ($scope.currentPage > 1) {
                $scope.currentPage--;
                $scope.updateDisplayedProducts();
            }
        };

        $scope.nextPage = function () {
            if ($scope.currentPage < $scope.totalPages) {
                $scope.currentPage++;
                $scope.updateDisplayedProducts();
            }
        };

        //nav bar
        $scope.openNav = function () {
            document.getElementById("mySidenav").style.width = "250px";
            document.getElementById("main").style.marginLeft = "250px";
            document.body.style.backgroundColor = "rgba(0,0,0,0.4)";
        };

        $scope.closeNav = function () {
            document.getElementById("mySidenav").style.width = "0";
            document.getElementById("main").style.marginLeft = "0";
            document.body.style.backgroundColor = "white";
        }

        //on click action
        $scope.isHiddenShow = true;
        $scope.countProduct = 0;
        $scope.openShow = function () {
            $scope.getProducts();
            $scope.isHiddenShow = !$scope.isHiddenShow;
            $scope.isHiddenCreate = true;
            $scope.isHiddenUpdate = true;
            $scope.isHiddenDelete = true;
        }
        $scope.isHiddenCreate = true;
        $scope.openCreate = function () {
            $scope.isHiddenCreate = !$scope.isHiddenCreate;
            $scope.isHiddenShow = true;
            $scope.isHiddenUpdate = true;
            $scope.isHiddenDelete = true;
        }
        $scope.isHiddenUpdate = true;
        $scope.openUpdate = function () {
            $scope.isHiddenUpdate = !$scope.isHiddenUpdate;
            $scope.isHiddenShow = true;
            $scope.isHiddenCreate = true;
            $scope.isHiddenDelete = true;
        }
        $scope.isHiddenDelete = true;
        $scope.openDelete = function () {
            $scope.isHiddenDelete = !$scope.isHiddenDelete;
            $scope.isHiddenShow = true;
            $scope.isHiddenCreate = true;
            $scope.isHiddenUpdate = true;
        }
    }]);