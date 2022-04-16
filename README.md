Sakara is MVC structure for GO to be easily to initialize your own project based on GIN framework.

# Routing 
In routes directory you will find two file that you can use them to registering your own routes.

`ExposedRoutes` : You can use it to set your public routes that are no need for make authentication  for them.

`AuthRoutes` : You can use it to set your routes that need authentication.

`Note` You can create your own route files with same package name `Routes`

After setting your routes you must to call theses function in Routing() in RouteApp.


# Middleware 

Middleware provides a simple way to filter HTTP requests as they enter your application. Sakara, for example, contains middleware that checks your application's user is authenticated. The middleware will respond with unauthenticated if the user is not authenticated however, the middleware will allow the request to continue inside the application.

# Controllers 

Instead of defining all of your request handling logic as closures in your route files, you may wish to organize this behavior using "controller" classes. 

You can create your own controller under the controllers directory 


# Database
By now default database available database is MYSQL,
And in .env file you can set your database credentials in connection.go file you will find makeConnection() you can change the database connection in it.


# Models 
Sakara uses ORM library called GORM with no limitation to change based on your needs

# Validations
Under validations directory you can see rules.go, the place that you can set your validation rules to be generic to use with one or more model.


# Resources
When building an API, you may need a transformation layer that sits between your Eloquent models and the JSON responses that are actually returned to your application's users.
So Sakara offers Resources that gives you the ability to control your returned data.

# Request Life Cycle

The main entry point is main.go, Sakara has main point to initialize the application with all required dependencies and configurations and run the application.
in Bootstrap.go sakara make sure that all required dependencies is run successfully like database connection

in main.go file you can see the calling of `Routes.RouterApp{app}` that all routes will be available to use.

When the request hits the service the controller will be available to handle it, and based on your logic the controller will use one of available Responses defined by default. 


`Note` You can see the whole example based on user model to guide you throw this journey.