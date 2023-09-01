angular.module('cfshop')
    .controller('adUserController', ['$scope', '$http', '$timeout', '$window', function ($scope, $http, $timeout, $window) {
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

        //count order inprogress
        $scope.countOrderInprogress = 0;
        $scope.countOrderInprogress = function () {
            $http.get('/orders/countInprogress')
                .then(function (response) {
                    $scope.countOrderInprogress = response.data;

                })
                .catch(function (error) {
                    console.error('Error retrieving users:', error);
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
                    console.error('Error retrieving users:', error);
                });
        };
        $scope.countOrderCompleted();

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

        //count order
        $scope.countOrder = 0;
        $scope.countOrder = function () {
            $http.get('/orders/count')
                .then(function (response) {
                    $scope.countOrder = response.data;

                })
                .catch(function (error) {
                    console.error('Error retrieving users:', error);
                });
        };
        $scope.countOrder();


        //// delete product
        //$scope.successNoti = false;
        //$scope.errorNoti = false;
        //$scope.product = {};
        //$scope.deleteProduct = function (productName) {
        //    $http.post('/products/delete', { name: productName })
        //        .then(function (response) {
        //            console.log(response);
        //            if (response.data && response.data.success) {
        //                $scope.successMess = "Product deleted successfully";
        //                $scope.successNoti = true;
        //            } else {
        //                $scope.errorNoti = true;
        //                $scope.errorMess = "Failed to delete product: " + response.data.message;
        //            }
        //        })
        //        .catch(function (error) {
        //            console.log(error);
        //            $scope.errorNoti = true;
        //            $scope.errorMess = 'Error deleting product: ' + error;
        //        })
        //    $timeout(function () {
        //        $scope.successNoti = false;
        //        $scope.errorNoti = false;
        //        $scope.getProducts();
        //    }, 1000);
        //};

        // update product
        $scope.successNoti = false;
        $scope.errorNoti = false;
        $scope.updateUser = {};

        $scope.updateUserSection = function (userID) {
            $scope.openUpdate();
            $http.get('/users?id=' + userID)
                .then(function (response) {
                    console.log(response);
                    $scope.updateUser = response.data;
                    $scope.updatedUser = angular.copy($scope.updateUser);
                })
                .catch(function (error) {
                    console.error('Error retrieving products:', error);
                });
        }
        $scope.userUpdate = function () {
            $http.put('/users/update', $scope.updatedUser)
                .then(function (response) {
                    console.log(response);
                    if (response.data && response.data.success) {
                        $scope.successMess = "User update successfully";
                        $scope.successNoti = true;
                    } else {
                        $scope.errorNoti = true;
                        $scope.errorMess = "Failed to update user: " + response.data.message;
                    }
                })
                .catch(function (error) {
                    console.log(error);
                    $scope.errorNoti = true;
                    $scope.errorMess = 'Error updating user: ' + error?.data;
                })
            $timeout(function () {
                $scope.successNoti = false;
                $scope.errorNoti = false;
                $scope.openShow();
            }, 2000);
        };

        //show 5 Users per page
        $scope.readUsers = [];
        $scope.displayedUsers = [];
        $scope.itemsPerPage = 5;
        $scope.currentPage = 1;
        $scope.totalPages = 0;

        $scope.getUsers = function () {
            $http.get('/users/get')
                .then(function (response) {
                    console.log(response)
                    $scope.readUsers = response.data;
                    $scope.totalPages = Math.ceil($scope.readUsers.length / $scope.itemsPerPage);
                    // Assign product numbers to each product
                    $scope.readUsers.forEach(function (product, index) {
                        product.productNumber = index + 1;
                    });
                    $scope.updateDisplayedUsers();
                })
                .catch(function (error) {
                    console.error('Error retrieving Users:', error);
                });
        };

        $scope.updateDisplayedUsers = function () {
            const startIndex = math.multiply(math.subtract($scope.currentPage, 1), $scope.itemsPerPage);
            const endIndex = startIndex + $scope.itemsPerPage;
            $scope.displayedUsers = $scope.readUsers.slice(startIndex, endIndex);
        };

        $scope.previousPage = function () {
            if ($scope.currentPage > 1) {
                $scope.currentPage--;
                $scope.updateDisplayedUsers();
            }
        };

        $scope.nextPage = function () {
            if ($scope.currentPage < $scope.totalPages) {
                $scope.currentPage++;
                $scope.updateDisplayedUsers();
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
            $scope.getUsers();
            $scope.isHiddenShow = !$scope.isHiddenShow;
            $scope.isHiddenCreate = true;
            $scope.isHiddenUpdate = true;
            $scope.isHiddenDelete = true;
            $scope.isHiddenAdd = true;
        }
        $scope.openAdd = function () {
            $scope.getUsers();
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