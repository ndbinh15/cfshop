angular.module('cfshop')
    .controller('productController', ['$scope', '$http', '$timeout','$window', function ($scope, $http, $timeout,$window) {
        $scope.linkAdminProducts = "/home/admin/products";
        $scope.linkAdminUsers = "/home/admin/users";
        $scope.linkAdminOrders = "/home/admin/orders";
        $scope.linkAdminHome = "/home/admin";

        //refresh
        $scope.refreshPage = function () {
            $timeout(function () {
                $window.location.reload();
            }, 1000);
          }

        // count product
        $scope.countProducts = function () {
            $http.get('/home/admin/products/count')
                .then(function (response) {
                    $scope.countProducts = response.data;

                })
                .catch(function (error) {
                    console.error('Error count products:', error);
                });
        };
        $timeout(function () {
            $scope.countProducts();
        }, 1000);

        // add product
        $scope.successNoti = false;
        $scope.errorNoti = false;
        $scope.product = {};
        $scope.createProduct = function () {
            $http.post('/home/admin/products/add', $scope.product)
                .then(function (response) {
                    console.log(response);
                    if (response.data && response.data.success) {
                        $scope.successMess = "Product created successfully";
                        $scope.product = {}; // Clear the form fields
                        $scope.successNoti = true;
                    } else {
                        $scope.errorNoti = true;
                        $scope.errorMess = "Failed to create product:" + response.data.message;
                    }
                })
                .catch(function (error) {
                    $scope.errorNoti = true;
                    $scope.errorMess = 'Error creating product:' + error;
                });
            $timeout(function () {
                $scope.successNoti = false;
                $scope.errorNoti = false;
            }, 3000);
        };


        //get category
        $scope.getCategories = function () {
            $http
                .get('/products/categories/get')
                .then(function (response) {
                    console.log(response)
                    $scope.categories = response.data;

                })
                .catch(function (error) {
                    console.error('Error get categories:', error);
                });
        }

        //show 5 products per page
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
        $scope.navChoose = false;
        $scope.contactNav = function () {
            $scope.navChoose = !$scope.navChoose;
            console.log($scope.navChoose)
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
            $scope.getCategories();
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