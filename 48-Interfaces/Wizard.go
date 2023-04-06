package Wizard

 // Create Wizard struct
 type Wizard struct{}
   
 // Wizard Receiver Function (method) - This is how we want the Wizard to Defend()
 func (w Wizard) Defend() string { 
	 return "Expelliarmus" 
 }
 
 // Wizard Receiver Function (method) - A Wizard has the unique ability to cast a Forget() spell
 func (w Wizard) Forget() string {
	 return "Obliviate"
 }


func Wizard()