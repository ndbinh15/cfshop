angular.module('cfshop')
    .controller('productController', ['$scope', '$http', '$timeout', '$window', function ($scope, $http, $timeout, $window) {
        $scope.linkAdminProducts = "/home/admin/products";
        $scope.linkAdminUsers = "/home/admin/users";
        $scope.linkAdminOrders = "/home/admin/orders";
        $scope.linkAdminHome = "/home/admin";

        //refresh
        $scope.refreshPage = function () {
            if ($scope.successNoti === true) {
                $timeout(function () {
                    $window.location.reload();
                }, 3000);

            }
        }
        $scope.refreshOriginPage = function () {
            $timeout(function () {
                $window.location.reload();
            }, 1000);
        }

        // count product
        $scope.countProducts = function () {
            $http.get('/products/count')
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
            $http.post('/products/add', $scope.product)
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
                    console.log(error)
                    $scope.errorNoti = true;
                    $scope.errorMess = 'Error creating product:' + error;
                });
            $timeout(function () {
                $scope.successNoti = false;
                $scope.errorNoti = false;
            }, 3000);
        };

        // delete product
        $scope.successNoti = false;
        $scope.errorNoti = false;
        $scope.product = {};
        $scope.deleteProduct = function (productName) {
            $http.post('/products/delete', { name: productName })
                .then(function (response) {
                    console.log(response);
                    if (response.data && response.data.success) {
                        $scope.successMess = "Product deleted successfully";
                        $scope.successNoti = true;
                    } else {
                        $scope.errorNoti = true;
                        $scope.errorMess = "Failed to delete product: " + response.data.message;
                    }
                })
                .catch(function (error) {
                    console.log(error);
                    $scope.errorNoti = true;
                    $scope.errorMess = 'Error deleting product: ' + error;
                })
            $timeout(function () {
                $scope.successNoti = false;
                $scope.errorNoti = false;
                $scope.getProducts();
            }, 1000);
        };

        // update product quantity
        $scope.successNoti2 = false;
        $scope.errorNoti2 = false;
        $scope.product2 = {};
        $scope.updateQuantity = function () {
            $http.post('/products/addQuantity', $scope.product2)
                .then(function (response) {
                    console.log(response);
                    if (response.data && response.data.success) {
                        $scope.successMess2 = "Product created successfully";
                        $scope.product2 = {}; // Clear the form fields
                        $scope.successNoti2 = true;
                    } else {
                        $scope.errorNoti2 = true;
                        $scope.errorMess2 = "Failed to create product:" + response.data.message;
                    }
                })
                .catch(function (error) {
                    $scope.errorNoti2 = true;
                    $scope.errorMess2 = 'Error creating product:' + error;
                });
        };


        //get category
        $scope.getCategories = function () {
            $http
                .get('/products/categories/get')
                .then(function (response) {
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
            $http.get('/products/get')
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
        $scope.isHiddenCreate = true;
        $scope.isHiddenUpdate = true;
        $scope.isHiddenDelete = true;
        $scope.isHiddenAdd = true;
        $scope.countProduct = 0;
        $scope.openShow = function () {
            $scope.getProducts();
            $scope.isHiddenShow = !$scope.isHiddenShow;
            $scope.isHiddenCreate = true;
            $scope.isHiddenUpdate = true;
            $scope.isHiddenDelete = true;
            $scope.isHiddenAdd = true;
        }
        $scope.openAdd = function () {
            $scope.getProducts();
            $scope.isHiddenAdd = !$scope.isHiddenAdd;
            $scope.isHiddenShow = true;
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
            $scope.isHiddenAdd = true;
            $scope.getCategories();
        }
        $scope.isHiddenUpdate = true;
        $scope.openUpdate = function () {
            $scope.isHiddenUpdate = !$scope.isHiddenUpdate;
            $scope.isHiddenShow = true;
            $scope.isHiddenCreate = true;
            $scope.isHiddenDelete = true;
            $scope.isHiddenAdd = true;
        }
        $scope.isHiddenDelete = true;
        $scope.openDelete = function () {
            $scope.isHiddenDelete = !$scope.isHiddenDelete;
            $scope.isHiddenShow = true;
            $scope.isHiddenCreate = true;
            $scope.isHiddenUpdate = true;
            $scope.isHiddenAdd = true;
        }
    }]);