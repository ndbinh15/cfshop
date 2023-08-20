angular.module('cfshop')
    .controller('userOrderedController', ['$scope', '$http', '$timeout', '$window', 'dataService',
        function ($scope, $http, $timeout, $window, dataService) {

            //link page /home/user/
            $scope.linkUserProfile = "/home/user/profile"
            $scope.linkUserCart = "/home/user/order"
            $scope.linkUserBuy = "/home/user/listproduct"
            $scope.linkUserOrdered = "/home/user/ordered"
            $scope.linkUserTracking = "/commingsoon"

            $scope.order = [];
            $scope.product = [];
            $scope.totalPrice = 0;
            $scope.totalBill = 0;
            $scope.errorID = true;
            $scope.getProduct = function (id, buyQuantity, cartID) {
                $http.get('/products?id=' + id)
                    .then(function (response) {
                        $scope.totalPrice = math.multiply(response.data.price, buyQuantity).toFixed(2);
                        $scope.product.push({ product: response.data, quantity: buyQuantity, price: $scope.totalPrice });
                        $scope.totalBill = math.add($scope.totalBill, $scope.totalPrice);
                        $scope.storeValue($scope.product, $scope.totalPrice, $scope.totalBill, cartID);
                    })
                    .catch(function (error) {
                        console.error('Error retrieving products:', error);
                        $window.location.href = '/error'
                    });
            };
            //$scope.getUserInfo = function (userID) {
            //    $http.get('/users?id=' + userID)
            //        .then(function (response) {
            //            console.log(response);
            //            $scope.userInfo = response.data;
            //        })
            //        .catch(function (error) {
            //            console.error('Error retrieving products:', error);
            //        });
            //}
            //$scope.getUserInfo(dataService.getData('userID'));
            $scope.getOrdered = function () {
                $http.get('/orders/get?id=' + dataService.getData('userID'))
                    .then(function (response) {
                        console.log(response)
                        if (response.data != null) {

                            $scope.errorID = false;
                            $scope.order = response.data;
                        }

                        //for (let i = 0; i < $scope.order.items.length; i++) {
                        //    const item = $scope.order.items[i];
                        //    $scope.getProduct(item.productId, item.quantity, $scope.order["_id"]);
                        //}
                    })
                    .catch(function (error) {
                        console.error('Error retrieving order:', error);
                    });
            };
            $scope.getOrdered();

            $scope.goBackCart = function () {
                $window.location.href = '/home/user/cart'
            }
        }])
