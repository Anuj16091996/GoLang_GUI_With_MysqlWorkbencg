package View

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

func StartGui2(){
	widgets.NewQApplication(len(os.Args), os.Args)

	loginWindow()
	widgets.QApplication_Exec()
}

func loginWindow(){
	var(

		/*Heading=widgets.NewQLabel2("Welcome To XYZ Bank",nil,0)*/
		userLabel = widgets.NewQLabel2("Customer Id",nil,0)
		userTextBox = widgets.NewQLineEdit(nil)
		pwdLabel = widgets.NewQLabel2("Password",nil,0)
		pwdTextBox = widgets.NewQLineEdit(nil)
		button2 = widgets.NewQPushButton2("Forget ?",nil)
		button = widgets.NewQPushButton2("Login",nil)
		loginLayout = widgets.NewQGridLayout2()
		loginGroup = widgets.NewQGroupBox(nil)
	)
	pwdTextBox.SetEchoMode(widgets.QLineEdit__Password)
	loginLayout.AddWidget2(userLabel,0,1,0)
	loginLayout.AddWidget2(userTextBox,0,2,0)
	loginLayout.AddWidget2(pwdLabel,1,1,0)
	loginLayout.AddWidget2(pwdTextBox,1,2,0)
	loginLayout.AddWidget2(button2,2,1,0)
	loginLayout.AddWidget2(button,2,2,0)
/*	loginLayout.AddWidget3(button,2,0,1,2,core.Qt__AlignJustify)
	loginLayout.AddWidget3(button2,2,2,1,2,core.Qt__AlignJustify)*/
	loginGroup.SetLayout(loginLayout)
	loginGroup.SetContentsMargins(10,10,10,10)
	loginGroup.SetFixedSize2(400,200)

	var(
		window = widgets.NewQMainWindow(nil,0)
		centralWidget = widgets.NewQWidget(window,0)
		layout = widgets.NewQGridLayout2()
	)
	window.SetWindowTitle("Welcome To My bank")
	window.SetFixedSize2(600,400)
	layout.AddWidget2(loginGroup,0,0,0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)
	window.Show()






}