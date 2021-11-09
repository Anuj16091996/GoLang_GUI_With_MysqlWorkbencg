package View

import (
	"Final_Project/Database"
	"Final_Project/Model"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
	"strconv"
	"unicode"
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

		button = widgets.NewQPushButton2("Login",nil)
		button2=widgets.NewQPushButton2("Sign Up",nil)
		loginLayout = widgets.NewQGridLayout2()
		loginGroup = widgets.NewQGroupBox(nil)
	)
	pwdTextBox.SetEchoMode(widgets.QLineEdit__Password)
	loginLayout.AddWidget2(userLabel,0,1,0)
	loginLayout.AddWidget2(userTextBox,0,2,0)
	loginLayout.AddWidget2(pwdLabel,2,1,0)
	loginLayout.AddWidget2(pwdTextBox,2,2,0)

	loginLayout.AddWidget2(button2,3,1,0)

	loginLayout.AddWidget2(button,3,2,0)
	/*loginLayout.AddWidget3(button2,2,2,1,2,core.Qt__AlignJustify)*/
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
	button.ConnectClicked(func(checked bool) {
		customer_id:=userTextBox.Text()
		password:=pwdTextBox.Text()
		IsValid:=Database.CheckLoginCredt(customer_id,password)
		if IsValid{
			ClinetWindow(customer_id)
			window.Close()
		}else {
			showDialog("Invalid","Invalid User Name or Password")
		}
	})

	button2.ConnectClicked(func(checked bool) {
		window.Close()
		SignUpWindow()

	})


}



func showDialog(title string, message string) {
	var (
		dialog = widgets.NewQDialog(nil,0)
		textMsg = widgets.NewQLabel2(message,nil,0)
		dialogBtns = widgets.NewQDialogButtonBox4(widgets.QDialogButtonBox__Close,core.Qt__Horizontal,nil)
		layout  =widgets.NewQGridLayout2()
	)
	dialogBtns.ConnectClicked(func(button *widgets.QAbstractButton) {
		if button.Text()=="Close"{
			dialog.Close()
		}
	})
	layout.AddWidget2(textMsg,0,0,0)
	layout.AddWidget2(dialogBtns,1,0,0)
	dialog.SetLayout(layout)
	dialog.SetWindowTitle(title)
	dialog.Exec()

}



func ClinetWindow(Customer_id string){

	Client_data:=Database.CustomerWindow(Customer_id)
	var(
		window = widgets.NewQMainWindow(nil,0)
		centralWidget = widgets.NewQWidget(window,0)
		layout = widgets.NewQGridLayout2()
		userLabel = widgets.NewQLabel2("Current Balance ",nil,0)
		TextBoz=widgets.NewQLineEdit(nil)

		userTextBox = widgets.NewQLineEdit2( strconv.Itoa(Client_data.Balace),nil)
		loginLayout = widgets.NewQGridLayout2()
		loginGroup = widgets.NewQGroupBox(nil)
		Deposit=widgets.NewQPushButton2("Deposit",nil)
		Withdraw=widgets.NewQPushButton2("Withdraw",nil)
		Logout=widgets.NewQPushButton2("LogOut",nil)

		)
		userTextBox.SetReadOnly(true)
	loginLayout.AddWidget2(userLabel,0,0,0)
	loginLayout.AddWidget2(userTextBox,0,4,0)
	loginLayout.AddWidget2(TextBoz,1,2,0)
	TextBoz.SetFixedSize2(180,100)
	loginLayout.AddWidget2(Deposit,3,0,0)
	loginLayout.AddWidget2(Withdraw,3,4,0)
	loginLayout.AddWidget2(Logout,4,2,0)
	loginGroup.SetLayout(loginLayout)
	loginGroup.SetContentsMargins(5,5,5,5)
	loginGroup.SetFixedSize2(400,250)
	window.SetWindowTitle("Welcome "+Client_data.First_Name+" "+ Client_data.Last_Name)
	window.SetFixedSize2(600,400)
	layout.AddWidget2(loginGroup,0,0,0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)
	window.Show()

	Deposit.ConnectClicked(func(checked bool) {

		Value:=TextBoz.Text()
		n, err :=strconv.Atoi(Value)
		if  n >=0  && err == nil{
			FinalAmount:=Client_data.Balace+n
			userTextBox.SetText(strconv.Itoa(FinalAmount) )
			Client_data.Balace=FinalAmount
			showDialog("Sucess","Money Deposited" );
		}else {
			showDialog("Invalid","Invalid Input")
		}
		TextBoz.Clear()

	})


	Withdraw.ConnectClicked(func(checked bool) {
		Value:=TextBoz.Text()
		n, err :=strconv.Atoi(Value)
		if  n >=0  && err == nil{
			if n>Client_data.Balace{
				showDialog("Gareeb","No Money-Gareeb" );
			}else {
				FinalAmount:=Client_data.Balace-n
				userTextBox.SetText(strconv.Itoa(FinalAmount) )
				Client_data.Balace=FinalAmount
				showDialog("Sucess","Money Withdraw" );
			}


		}else {
			showDialog("Invalid","Invalid Input")
		}
		TextBoz.Clear()


	})
	
	
	Logout.ConnectClicked(func(checked bool) {
		Database.UpdateBankBalance(Client_data)
		showDialog("Thank You","Have A Nice Day Ahead" );
		window.Close()
		loginWindow()
	})

}

func SignUpWindow(){
	var(

	FirstName = widgets.NewQLabel2("First Name",nil,0)
	Fname = widgets.NewQLineEdit(nil)

		LastName = widgets.NewQLabel2("Last Name",nil,0)
		Lname = widgets.NewQLineEdit(nil)

		DepostText=widgets.NewQLabel2("How Much You want to Deposit-: ",nil,0)
		Depost=widgets.NewQLineEdit(nil)

		Phone=widgets.NewQLabel2("Enter Your Number ",nil,0)
		Number=widgets.NewQLineEdit(nil)

	pwdLabel = widgets.NewQLabel2("Password",nil,0)
	pwdTextBox = widgets.NewQLineEdit(nil)

		pwdLabel2 = widgets.NewQLabel2("Confirm Password",nil,0)
		pwdTextBox2 = widgets.NewQLineEdit(nil)

	button2=widgets.NewQPushButton2("Resgiter",nil)
	loginLayout = widgets.NewQGridLayout2()
	loginGroup = widgets.NewQGroupBox(nil)
	)

	pwdTextBox.SetEchoMode(widgets.QLineEdit__Password)
	pwdTextBox2.SetEchoMode(widgets.QLineEdit__Password)

	loginLayout.AddWidget2(FirstName,0,1,0)
	loginLayout.AddWidget2(Fname,0,2,0)

	loginLayout.AddWidget2(LastName,1,1,0)
	loginLayout.AddWidget2(Lname,1,2,0)

	loginLayout.AddWidget2(pwdLabel,2,1,0)
	loginLayout.AddWidget2(pwdTextBox,2,2,0)

	loginLayout.AddWidget2(pwdLabel2,3,1,0)
	loginLayout.AddWidget2(pwdTextBox2,3,2,0)

	loginLayout.AddWidget2(Phone,4,1,0)
	loginLayout.AddWidget2(Number,4,2,0)



	loginLayout.AddWidget2(DepostText,5,1,0)
	loginLayout.AddWidget2(Depost,5,2,0)
	loginLayout.AddWidget3(button2,6,1,1,2,0)

	loginGroup.SetLayout(loginLayout)
	loginGroup.SetContentsMargins(10,10,10,10)
	loginGroup.SetFixedSize2(400,300)

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

	button2.ConnectClicked(func(checked bool) {

			var Details Model.Body
			AllValue:=0
			if Fname.Text()=="" || Lname.Text()==""||Depost.Text()==""||pwdTextBox.Text()=="" ||Number.Text()=="" {
				showDialog("Invalid","Invalid Input Please Provaide All Information")
			}else {
				FnameC:=Fname.Text()
				Result:=CheckString(FnameC)

				if Result==false{
					showDialog("Invalid","Invalid Input On First Name")
					Fname.Clear()
				}else {
					Details.First_Name=Fname.Text()
					AllValue=AllValue+1
				}

				LnameC:=Lname.Text()
				Result1:=CheckString(LnameC)
				if Result1==false{
					showDialog("Invalid","Invalid Input On Last Name")
					Lname.Clear()
				}else {
					Details.Last_Name=Lname.Text()
					AllValue=AllValue+1
				}

				if pwdTextBox.Text()!=pwdTextBox2.Text(){
					showDialog("Invalid","Password Doesnt Match")
					pwdTextBox2.Clear()
					pwdTextBox.Clear()
				}else {
					Details.Online_Password=pwdTextBox.Text()
					AllValue=AllValue+1
				}

				Value:=Depost.Text()
				n, err :=strconv.Atoi(Value)
				if  n >=0  && err == nil{

					Details.Balace= n
					AllValue=AllValue+1
				}else {
					showDialog("Invalid", "Invalid Input on Deposit Balance")
					Depost.Clear()
				}



				Value2:=Number.Text()
				n2, err2 :=strconv.Atoi(Value2)
				if  n2 >999999999 && n2<=9999999999  && err2 == nil{

					value:=Database.CheckNumber(n2)

					if value==false{
						Details.MobileNumber= n2
						AllValue=AllValue+1
					}else {
						showDialog("Invalid", "Number Already Register With Us")
						Number.Clear()
					}



				}else {
					showDialog("Invalid", "Invalid Input on Phone Number")
					Number.Clear()
				}

				if AllValue==5{
					Message:=Database.InsertUser(Details)
					Fname.Clear()
					Lname.Clear()
					pwdTextBox2.Clear()
					pwdTextBox.Clear()
					Depost.Clear()
					Number.Clear()
				window.Close()
					showDialog("SucessFully Sign Up", "Your Customer ID is "+ Message+" Please Note It down" +
						" this is not a team project, so i didnt fetch any API to send a text message, and you know your" +
						"Password, \nThank you so much for using this application and\n enjoy your pizza")

					loginWindow()
				}
			}




	})


}

func CheckString(s string) bool{
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true

}

