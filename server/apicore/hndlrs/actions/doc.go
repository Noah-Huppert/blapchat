/*
Package actions is a package which holds the Creator, Getter, Updater, and Deleter
interfaces. These interfaces all define a single function which follows the
HandleFn signature. This function will be called to handle events. The response
will be returned to the client.

Each of these defined functions match directly with one of the 4 allowed Event
actions:

	- Create => Creator
	- Get => Getter
	- Update => Updater
	- Delete => Deleter

By implementing one of the above interfaces, a type signifies that it would like
that specific handler function to be registered under its corresponding action.

This allows one to register Handlers in a much simpler and expressive way. By
simply implementing the interfaces of the actions you wish your resource to
provide.

A Loader, named the ResourceLoader, can be used to automatically register any
Creators, Getters, Updaters, and / or Deleters implemented by a type, with a
Registry.
*/
package actions
