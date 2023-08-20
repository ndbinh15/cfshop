angular.module('cfshop')
    .controller('userCartController', ['$scope', '$http', '$timeout', '$window', 'dataService',
        function ($scope, $http, $timeout, $window, dataService) {

            //link page /home/user/
            $scope.linkUserProfile = "/home/user/profile"
            $scope.linkUserCart = "/home/user/cart"
            $scope.linkUserBuy = "/home/user/listproduct"
            $scope.linkUserOrdered = "/home/user/ordered"
            $scope.linkUserTracking = "/commingsoon"

            $scope.cart = [];
            $scope.product = [];
            $scope.totalPrice = 0;
            $scope.totalBill = 0;
            $scope.errorID = true;
            $scope.getProduct = function (id, buyQuantity, cartID) {
                $http.get('/products?id=' + id)
                    .then(function (response) {
                        $scope.totalPrice = math.multiply(response.data.price, buyQuantity).toFixed(2);
                        $scope.product.push({ product: response.data, quantity: buyQuantity, price: $scope.totalPrice });
                        $scope.totalBill = math.add($scope.totalBill, $scope.totalPrice).toFixed(2);
                        $scope.storeValue($scope.product, $scope.totalPrice, $scope.totalBill, cartID);
                    })
                    .catch(function (error) {
                        console.error('Error retrieving products:', error);
                        $window.location.href = '/error'
                    });
            };
            $scope.getUserInfo = function (userID) {
                $http.get('/users?id=' + userID)
                    .then(function (response) {
                        console.log(response);
                        $scope.userInfo = response.data;
                    })
                    .catch(function (error) {
                        console.error('Error retrieving products:', error);
                    });
            }
            $scope.getUserInfo(dataService.getData('userID'));
            $scope.getCart = function () {
                $http.get('/cart/get?id=' + dataService.getData('userID'))
                    .then(function (response) {
                        console.log(response)
                        $scope.errorID = false;
                        $scope.cart = response.data;

                        for (let i = 0; i < $scope.cart.items.length; i++) {
                            const item = $scope.cart.items[i];
                            $scope.getProduct(item.productId, item.quantity, $scope.cart["_id"]);
                        }
                    })
                    .catch(function (error) {
                        console.error('Error retrieving products:', error);
                    });
            };
            $scope.getCart();

            $scope.order = {}

            $scope.storeValue = function (product, totalPrice, totalBill, cartID) {
                $scope.order.product = product;
                $scope.order.totalPrice = totalPrice;
                $scope.order.totalBill = totalBill;
                $scope.order.cartID = cartID;
                $scope.order.userID = dataService.getData('userID')
            }

            // add order
            $scope.successNoti = false;
            $scope.errorNoti = false;
            $scope.createOrder = function () {
                console.log($scope.order)
                $http.post('/orders/create', $scope.order)
                    .then(function (response) {
                        console.log(response);
                        if (response.data && response.data.success) {
                            $scope.successMess = "order created successfully";
                            $scope.successNoti = true;
                            $timeout(function () {
                                $window.location.href = $scope.linkUserOrdered;
                            }, 2000)
                        } else {
                            $scope.errorNoti = true;
                            $scope.errorMess = "Failed to create order:" + response.data.message;
                        }
                    })
                    .catch(function (error) {
                        console.log(error)
                        $scope.errorNoti = true;
                        $scope.errorMess = 'Error creating order:' + error?.data;
                    });
                $timeout(function () {
                    $scope.successNoti = false;
                    $scope.errorNoti = false;
                }, 2000);
            };


            $scope.goBackListProduct = function () {
                $window.location.href = '/home/user/listproduct'
            }
        }])
