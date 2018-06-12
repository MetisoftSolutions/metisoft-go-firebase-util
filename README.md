# metisoft-go-firebase-util

This module provides an abstraction over the `firebase.google.com/go` module to better suit Metisoft server applications written in Go.

## Installation

1. In your Go application, add `firebaseUtil "github.com/MetisoftSolutions/metisoft-go-firebase-util"` to your imports list.
2. Make a call to `godep get`, or whatever other dependency manager you use to retrieve the `metisoft-go-firebase-util` module.
3. In your server's initialization code, make a call to `firebaseUtil.Init()`, which requires at least `PathToServiceAccountKey` and `DatabaseUrl` to be set in the options argument.

`PathToServiceAccountKey` should be the **full** path to the `serviceAccountKey.json` file that you should get from the Firebase Console. You can get this key from Project Settings -> Service Accounts -> Firebase Admin SDK -> Generate New Private Key.

`DatabaseUrl` is the URL to the Firebase database for your app. This can also be found on the Firebase Admin SDK page.

## Usage

### Push notifications

You can send push notifications using `firebaseUtil.SendPushNotification()`. The first argument is some form of user ID. By default, this is a Firebase token. However, if you supply a function to the configuration options object's `FnGetFirebaseTokenForUser` key in the `Init()` call, then you can pass in a user ID native to your user authentication system. This relies on you defining your `FnGetFirebaseTokenForUser` function to return a Firebase token when given a user ID.