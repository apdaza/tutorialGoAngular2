// Creación del módulo
var tutorialApp = angular.module('tutorialApp', ['ngRoute']);

// Configuración de las rutas
tutorialApp.config(function($routeProvider) {

    $routeProvider
        .when('/home', {
            templateUrl : 'views/home.html',
            controller  : 'homeController'
        })
        .when('/listarContactos', {
            templateUrl : 'views/listarContactos.html',
            controller  : 'contactosController'
        })
        .when('/agregarContacto', {
            templateUrl : 'views/agregarContacto.html',
            controller  : 'addContactoController'
        })
        .when('/about', {
            templateUrl : 'views/about.html',
            controller  : 'aboutController'
        })
        .otherwise({
            redirectTo: '/home'
        });
});

tutorialApp.controller('homeController', function($scope) {
  $scope.title = 'Contactos';
  $scope.message = 'Versión 0.0.1';
});

tutorialApp.controller('contactosController', function($scope, $http) {
  $scope.title = 'Contactos';
  $scope.message = 'Listado de Contactos';

  $http.get('http://localhost:8080/api/contactos/contactos')
  .then(function(response) {
      $scope.data = response.data;
  });
});

tutorialApp.controller('addContactoController', function($scope, $http) {
  $scope.title = 'Contactos';
  $scope.message = 'Agregar Contacto';

  $scope.reset = function(){
    $scope.contacto = {};
  }

  $scope.add = function(){
    if($scope.contacto.nombre == null){
      return;
    }
    if($scope.contacto.apellido == null){
      return;
    }
    if($scope.contacto.celular == null){
      return;
    }
    if($scope.contacto.mail == null){
      return;
    }
    var data = {
        nombre: $scope.contacto.nombre,
        apellido: $scope.contacto.apellido,
        celular: $scope.contacto.celular,
        mail: $scope.contacto.mail
    };
    $http.post('http://localhost:8080/api/contactos/contactos',data);

    $scope.contacto = {};
  };
});

tutorialApp.controller('aboutController', function($scope) {
  $scope.title = 'Contactos';
  $scope.message = 'Esta es la página "Acerca de"';
});

tutorialApp.controller('parserMenu', function($scope, $http) {
  $http.get('models/menu.json')
       .then(function(res){
          $scope.menu = res.data.menu;
        });
});
