# How to package the program?

To package the program first you need to install the `fyne` cmd tool by
executing the following:

> go install fyne.io/fyne/v2/cmd/fyne@latest

Once, You have installed the `fyne` cmd tool you next need to have logo
file and rename it as Icon.png.

Now all you need to do is run the following command in the terminal:

> fyne package

If you want to specify custom icon file name: 

> fyne package -icon <iconFile>.png

You can also compile the application for various other operating
systems by executing the command like such:

> fyne package -os windows

Android:

fyne package -os android -appID com.example.myapp -icon mobileIcon.png
