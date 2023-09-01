(function () {
    var app = angular.module('cfshop', ['ngRoute']);


    app.service('dataService', [function () {

        return {
            setData: function (key, data) {
                localStorage.setItem(key, JSON.stringify(data));
            },
            getData: function (key) {
                var storedData = localStorage.getItem(key);
                return storedData ? JSON.parse(storedData) : null;
            },
            clearData: function () {
                localStorage.clear();
            }
        };
    }]);

    //app.factory('DataService', function () {
    //    var EXPIRATION_KEY = 'dataExpiration';
    //    var DATA_KEY = 'myDataKey';

    //    function storeData(data, expirationMinutes) {
    //        var expirationTimestamp = new Date().getTime() + (expirationMinutes * 60 * 1000);
    //        localStorage.setItem(EXPIRATION_KEY, expirationTimestamp);
    //        localStorage.setItem(DATA_KEY, JSON.stringify(data));
    //    }

    //    function getData() {
    //        var expirationTimestamp = parseInt(localStorage.getItem(EXPIRATION_KEY));
    //        var currentTime = new Date().getTime();

    //        if (expirationTimestamp && currentTime < expirationTimestamp) {
    //            return JSON.parse(localStorage.getItem(DATA_KEY));
    //        } else {
    //            // Data has expired or doesn't exist
    //            return null;
    //        }
    //    }

    //    return {
    //        storeData: storeData,
    //        getData: getData
    //    };
    //});

    app.config(["$routeProvider", "$locationProvider", function ($routeProvider, $locationProvider) {

        $locationProvider.html5Mode({
            enabled: true,
            requireBase: false
        });

        $routeProvider
            .when("/", {
                templateUrl: window.config.viewLocation + "/welcome.html",
                controllerUrl: window.config.controllerLocation + "/welcome.js"
            })
            .when("/signup", {
                templateUrl: window.config.viewLocation + "/signup.html",
                controllerUrl: window.config.controllerLocation + "/signup.js"
            })
            .when("/login", {
                templateUrl: window.config.viewLocation + "/login.html",
                controllerUrl: window.config.controllerLocation + "/login.js"
            })
            .when("/forgotpss", {
                templateUrl: window.config.viewLocation + "/fgpass.html",
                controllerUrl: window.config.controllerLocation + "/fgpass.js"
            })
            .when("/home", {
                templateUrl: window.config.viewLocation + "/home.html",
                controllerUrl: window.config.controllerLocation + "/home.js",
                //resolve: {
                //    isAdmin: ['$http', function ($http) {
                //        return $http.get('/check-role').then(function (response) {
                //            if (response.data.success) {
                //                return true;
                //            } else {
                //                return false;
                //            }
                //        });
                //    }]
                //}
            })
            .when("/home/admin", {
                templateUrl: window.config.viewLocation + "/admin.html",
                controllerUrl: window.config.controllerLocation + "/admin.js"
            })
            .when("/home/admin/products", {
                templateUrl: window.config.viewLocation + "/adProducts.html",
                controllerUrl: window.config.controllerLocation + "/adProducts.js"
            })
            .when("/home/admin/users", {
                templateUrl: window.config.viewLocation + "/adUsers.html",
                controllerUrl: window.config.controllerLocation + "/adUsers.js"
            })
            .when("/home/admin/orders", {
                templateUrl: window.config.viewLocation + "/adOrders.html",
                controllerUrl: window.config.controllerLocation + "/adOrders.js"
            })
            .when("/home/user", {
                templateUrl: window.config.viewLocation + "/user.html",
                controllerUrl: window.config.controllerLocation + "/user.js"
            })
            .when("/home/user/listproduct", {
                templateUrl: window.config.viewLocation + "/listproduct.html",
                controllerUrl: window.config.controllerLocation + "/listproduct.js"
            })
            .when("/home/user/cart", {
                templateUrl: window.config.viewLocation + "/userCart.html",
                controllerUrl: window.config.controllerLocation + "/userCart.js"
            })
            .when("/home/user/ordered", {
                templateUrl: window.config.viewLocation + "/userOrdered.html",
                controllerUrl: window.config.controllerLocation + "/userOrdered.js"
            })
            .when("/home/user/profile", {
                templateUrl: window.config.viewLocation + "/userProfile.html",
                controllerUrl: window.config.controllerLocation + "/userProfile.js"
            })
            .when("/home/user/listproduct/product", {
                templateUrl: window.config.viewLocation + "/product.html",
                controllerUrl: window.config.controllerLocation + "/product.js"
            })
            .when("/commingsoon", {
                templateUrl: window.config.baseLocation + "/commingsoon.html",
                controllerUrl: window.config.baseLocation + "/commingsoon.js"
            })
            .otherwise({
                templateUrl: window.config.baseLocation + "/error.html",
                controllerUrl: window.config.baseLocation + "/error.js"
            })
    }]);
}());