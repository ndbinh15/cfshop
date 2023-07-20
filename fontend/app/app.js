(function () {
    var app = angular.module('cfshop', ['ngRoute']);

    app.config(["$routeProvider", "$locationProvider", function ($routeProvider, $locationProvider) {

        $locationProvider.html5Mode({
            enabled: true,
            requireBase: false
        });

        $routeProvider
            .when("/", {
                templateUrl: window.config.viewLocation + "/login.html",
                controllerUrl: window.config.controllerLocation + "/login.js"
            })
            .when("/home", {
                templateUrl: window.config.viewLocation + "/home.html",
                controllerUrl: window.config.controllerLocation + "/home.js"
            })
            .when("/home/admin", {
                templateUrl: window.config.viewLocation + "/admin.html",
                controllerUrl: window.config.controllerLocation + "/admin.js"
            })
            .when("/home/admin/products", {
                templateUrl: window.config.viewLocation + "/products.html",
                controllerUrl: window.config.controllerLocation + "/products.js"
            })
            .when("/home/user", {
                templateUrl: window.config.viewLocation + "/user.html",
                controllerUrl: window.config.controllerLocation + "/user.js"
            })
            .otherwise({
                templateUrl: window.config.viewLocation + "/login.html",
                controllerUrl: window.config.controllerLocation + "/login.js"
            })
    }]);
}());