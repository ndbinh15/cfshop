angular.module('cfshop')
    .controller('productController', ['$scope', '$http', '$window', '$timeout', 'dataService', function ($scope, $http, $window, $timeout, dataService) {

        //link page /home/user/
        $scope.linkUserProfile = "/commingsoon"
        $scope.linkUserCart = "/commingsoon"
        $scope.linkUserBuy = "/home/user/listproduct"
        $scope.linkUserOrdered = "/commingsoon"
        $scope.linkUserTracking = "/commingsoon"

        const id = new URLSearchParams(window.location.search).get("id");
        $scope.errorID = true;
        $scope.product = [];
        $scope.getProduct = function (id) {
            $http.get('/products?id=' + id)
                .then(function (response) {
                    console.log(response);
                    $scope.product = response.data;
                    $scope.errorID = false;
                })
                .catch(function (error) {
                    console.error('Error retrieving products:', error);
                    $window.location.href = '/error'
                });
        };
        $scope.getProduct(id);

        $scope.buyQuantity = 1;
        $scope.addToCart = function () {
            if ($scope.buyQuantity != undefined && $scope.buyQuantity != null && $scope.buyQuantity > 0) {
                var cartRequest = {
                    userId: dataService.getData('userID'),
                    items: [
                        { productId: id, quantity: $scope.buyQuantity }
                        // Add more items if needed
                    ]
                };
                console.log(cartRequest)
                $http.post('/cart/add', cartRequest)
                    .then(function (response) {
                        $scope.successMess = "Add to cart successfully";
                        $scope.successNoti = true;
                        console.log(response);
                    })
                    .catch(function (error) {
                        $scope.errorNoti = true;
                        $scope.errorMess = "Failed to add to cart: " + response.data.message;
                        console.log('Error:', error);
                    });
                $timeout(function () {
                    $scope.successNoti = false;
                    $scope.errorNoti = false;
                }, 2000);
            } else {
                $scope.errorNoti = true;
                $scope.errorMess = "Invalid quantity";
                $timeout(function () {
                    $scope.errorNoti = false;
                }, 2000);
            }

        };


        $scope.goBackListProduct = function () {
            $window.location.href = '/home/user/listproduct'
        }
        $scope.goToCart = function () {
            $window.location.href = '/home/user/cart'
        }
    }])
