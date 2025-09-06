// Import Firebase compat scripts
importScripts('https://www.gstatic.com/firebasejs/9.22.2/firebase-app-compat.js');
importScripts('https://www.gstatic.com/firebasejs/9.22.2/firebase-messaging-compat.js');

// Initialize Firebase inside Service Worker
firebase.initializeApp({
  apiKey: "AIzaSyC9slEJJveJecsnjOuLg1UpyKDV146xvQY",
  authDomain: "push--notification-task.firebaseapp.com",
  projectId: "push--notification-task",
  storageBucket: "push--notification-task.firebasestorage.app",
  messagingSenderId: "213486219690",
  appId: "1:213486219690:web:dc2063ff36f658576d7bd3"
});

// Retrieve messaging object
const messaging = firebase.messaging();

  // Handle background messages
    messaging.onBackgroundMessage((payload) => {
    console.log('Background message received:', payload);

   const notificationTitle = payload.notification?.title || 'New Message';
   const notificationOptions = {
     body: payload.notification?.body || 'You Got a new Notification!',
     icon: payload.notification?.icon || '/firebase-logo.png'
   };

 self.registration.showNotification(notificationTitle, notificationOptions);
});
