# Snap But Doesnt Suck
SnapChat but it doesn't suck.

# Server
The Server component runs the back end for the Snap But Doesnt Suck app.

Since Gradle is already used as the build system for the Android app it is used for the Server.

All server code is located in the `server` directory.

## Database
On first setup run the `dbCreate` task to create Postgres Docker container for development. 
You can use the `dbDestroy` task to delete this Docker container.

After first setup you can use the `dbStart` and `dbStop` task to manage the Postgres Docker 
container.

## Dependency Vendoring
To vendor dependencies run the `goVendor` task.

To lock dependency versions run the `goLock` task.

## Building
To build the server run the `goBuild` task.

You will find the output in `server/.gogradle/<Platform_Name>_server`

## Running
To run the server run the `run` task.

This will automatically build the server, set run permissions on the 
executable, and run the executable.

# Android App
The Android app is developed with Android studio. Simply import this repository into 
Android Studio and develop as you normally would.
