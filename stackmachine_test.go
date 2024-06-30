package main

import (
	"errors"
	"testing"
)

/*func TestSplitString(t *testing.T){
	String:= "1 2 3"
	want:= []string("5","6","8")

} */

func TestSplitEmptyString(t *testing.T){
	String:= ""

	_, GotErr:=StackMachine(String)

	if GotErr== nil{
		t.Error("Expected error when splitting empty string")
	}
}

// Write your own TDD tests here as you develop

func TestEmptyError(t *testing.T){
	_, GotErr:= StackMachine("")

	WantErr:= errors.New("empty error")

	if GotErr.Error()!=WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}

}

func TestInvalidCommand(t *testing.T){
	_, GotErr:= StackMachine("DOGBANA")

	WantErr:= errors.New("Invalid Command")

	if GotErr.Error()!=WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
}																			

func TestInvalidCommandContainingCorrectSymbol(t *testing.T){

	_, GotErr:= StackMachine("+hello-")

	WantErr:= errors.New("Invalid Command")

	if GotErr.Error()!=WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
}													


/*func TestValidCommand(t *testing.T){
	_, GotErr:= StackMachine("-")

	want:= errors.New("Valid Command")     //This test will likley have th change as valid command error will not be used

	if GotErr.Error()!=want.Error(){
		t.Error("Wanted",want,"But didnt get Valid Command prompt")
	}
} */


func TestPopSingleElement(t *testing.T){
	result, GotErr:=StackMachine("1 POP")
	want:= errors.New("Empty Stack")

	if GotErr.Error()!= want.Error(){
		t.Error("Expected error due to empty stack")
	}
	if result!=0{
		t.Error("Expected 0 got",result)
	}
	
}

func TestPopAfterMultipleNumbers(t *testing.T){
	result, GotErr:= StackMachine("1 2 3 POP")
	want:= 2
	if GotErr!=nil {
		t.Error("Unexpected Error", GotErr.Error())
	}
	if result != want{
		t.Error("Expected",want,"Got",result)
	}
}

func TestPopTwice(t *testing.T){
	result, GotErr:= StackMachine("1 2 3 POP POP")
	want:= 1
	if GotErr!= nil{
		t.Error("Unexpected Error", GotErr.Error())
	}
	if result != want{
		t.Error("Expected",want,"Got",result)
	}
}

func TestPopEmptyStack(t *testing.T){
	_, GotErr:= StackMachine("POP")
	WantErr:= errors.New("Empty Stack")

	if GotErr.Error()!= WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
}

func TestPOPEmpty(t *testing.T){
	result, GotErr:= StackMachine("1 2 POP POP")
	WantErr:= errors.New("Empty Stack")

	if GotErr.Error()!= WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
	if result!=0{
		t.Error("Expected 0 got",result)
	}
}

func TestPOPAlone(t *testing.T){
	result,GotErr:= StackMachine("POP")
	WantErr:= errors.New("Empty Stack")
	if GotErr.Error()!= WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
	if result!=0{
		t.Error("Expected 0 got",result)
	}
}

func TestDupSingleElement(t *testing.T) {
    result, GotErr := StackMachine("1 DUP")
    if GotErr != nil {
		t.Error("Unexpected Error", GotErr.Error())
    }
    want := 1
    if result != want {
        t.Error("Expected",want,"Got", GotErr)
    }
}

func TestDupMultipleElements(t *testing.T) {
    result, GotErr := StackMachine("1 2 3 DUP")
    if GotErr != nil {
		t.Error("Unexpected Error", GotErr.Error())
    }
    want := 3
    if result != want {
		t.Error("Expected",want,"Got", GotErr)
    }
}

func TestDupEmptyStack(t *testing.T) {
    _, GotErr := StackMachine("DUP")
    WantErr := errors.New("Empty Stack")
    if GotErr.Error() != WantErr.Error() {
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
    }
}

func TestAddWithElements(t *testing.T){
	result, GotErr:= StackMachine("1 2 +")
	if GotErr!=nil{
		t.Error("Unexpected Error", GotErr.Error())
	}
	want:=3
	if result!=want{
		t.Error("Expected",want,"Got",result)
	}
}

func TestAddTooFewElements(t *testing.T){
	result, GotErr:= StackMachine("99 +")
	WantErr:= errors.New("Too Few Elements")
	if GotErr.Error()!= WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
	
	if result!=0{
		t.Error("Expected 0 Got",result)
	}
}


func TestMinusTooFewElements(t *testing.T){
	result, GotErr:= StackMachine("99 -")
	WantErr:= errors.New("Too Few Elements")
	if GotErr.Error()!= WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
	
	if result!=0{
		t.Error("Expected 0 Got",result)
	}
}


func TestMultiplyTooFewElements(t *testing.T){
	result, GotErr:= StackMachine("99 *")
	WantErr:= errors.New("Too Few Elements")
	if GotErr.Error()!= WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
	
	if result!=0{
		t.Error("Expected 0 Got",result)
	}
}

//ADD POP DUP TESTS TOGETHER

func TestPopAndDup(t *testing.T) {
    result, GotErr := StackMachine("1 2 3 POP DUP")
    if GotErr != nil {
		t.Error("Unexpected Error", GotErr.Error())
    }
    want := 2
    if result != want {
		t.Error("Expected",want,",Got",result)
    }
}




func TestPopAndDupWithSingleElement(t *testing.T) { 
	result, GotErr := StackMachine("1 POP DUP") 
	want := errors.New("Empty Stack") 
	if GotErr.Error()!= want.Error(){
		t.Error("Expected",want,"Got the error",GotErr) 
		}
		 if result != 0 { 
			t.Error("Expected 0 Got",result) 
		}
		 }

func TestDupWithMultipleElements(t *testing.T) {
    result, GotErr := StackMachine("1 2 3 DUP")
    if GotErr != nil {
		t.Error("Unexpected Error", GotErr.Error())
    }
    want := 3 
	if result!=want{
		t.Error("Expected",want,"Got",result)
	}
}

/*func TestAddDupPop(t *testing.T){
	result, GotErr := StackMachine("1 2 DUP + POP")
	if GotErr!= nil{
		t.Error("Unexpected Error")
	}
	want:= 4										//Need to fix
	if result!= want {
		t.Error("Expected",want,"Got",result)
	}
} */

func TestAddPopDup(t *testing.T){
	result, GotErr := StackMachine("1 2 + POP DUP")
	WantErr:= errors.New("Empty Stack")

	if GotErr.Error()!=WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
	if result!=0{
		t.Error("Expected 0 Got",result)
	}
}

func TestAddOperation(t *testing.T){
	result, GotErr:= StackMachine("3 4 +")

	if GotErr != nil {
		t.Error("Expected no Error got",GotErr.Error() )
	}
	want:=7
	if result !=want{
		t.Error("Expected",want,"Got",result,)
	}
}


func TestAddWithNegativeElements(t *testing.T){
	result, GotErr:= StackMachine("3 -1 +")

	if GotErr != nil{
		t.Error("Unexpected Error", GotErr.Error())    // shouldnt be able to use negative
	}
	want:=2
	if result!= want{
		t.Error("Expected",want,"Got",result)
	}
}




func TestEmptyStackAfterClear(t *testing.T){
	result, GotErr:= StackMachine("99 CLEAR")
	WantErr:= errors.New("Empty Stack")

	if GotErr.Error() != WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
	if result != 0{
		t.Error("Expected 0 Got",result) 
	}

}

func TestSumSingleValue(t *testing.T){
	result, GotErr:= StackMachine("99 SUM")
	want:= 99

	if GotErr != nil {
		t.Error("Expected no Error got",GotErr.Error() )
	}
	if result!= want{
		t.Error("Expected",want,"Got",result)
	}
}

func TestSumEmpty(t *testing.T){
	result, GotErr:= StackMachine("SUM")
	WantErr:= errors.New("Empty Stack")

	if GotErr.Error() != WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
	if result != 0{
		t.Error("Expected 0 Got",result) 
	}

}

func TestAddMultiply(t *testing.T){
	result, GotErr:= StackMachine("5 6 + 2 *")
	want:= 22

	if GotErr != nil {
		t.Error("Expected no Error got",GotErr.Error())
	}
	if result!= want{
		t.Error("Expected",want,"Got",result)
	}
}

func TestClearTooFew(t *testing.T){
	result, GotErr:= StackMachine("1 2 3 4 + CLEAR 12 +")
	WantErr:= errors.New("Empty Stack")

	if GotErr.Error() != WantErr.Error(){
		t.Error("Expected",WantErr.Error(),"Got",GotErr.Error())
	}
	if result != 0{
		t.Error("Expected 0 Got",result) 
	}

}

func TestSingleInteger(t *testing.T){
	result, GotErr:= StackMachine("15")

	if GotErr != nil {
		t.Error("Expected no Error got",GotErr.Error() )
	}
	want:=15
	if result !=want{
		t.Error("Expected",want,"Got",result,)
	}
}


//func TestAddMinusMultiply(t *testing.T){}

//func TestOverflow(t *testing.T){}

func TestMinus(t *testing.T){
	result, GotErr:= StackMachine("2 5 -")

	if GotErr != nil {
		t.Error("Expected no Error got",GotErr.Error(), result )
	}
	want:=3
	if result !=want{
		t.Error("Expected",want,"Got",result,)
	}
}

																//tests to implement will also look back through code as its currently a mess




func TestMultiply(t *testing.T){
	result, GotErr:= StackMachine("3 4 *")

	if GotErr != nil {
		t.Error("Expected no Error got",GotErr.Error() )
	}
	want:=12
	if result !=want{
		t.Error("Expected",want,"Got",result,)
	}
}

func TestSingleNumberOverflow(t *testing.T) {
    _, GotErr := StackMachine("50001")
    WantErr := errors.New("Overflow Error")
    if GotErr.Error() != WantErr.Error() {
        t.Error("Expected", WantErr.Error(), "Got", GotErr.Error())
    }
}



func TestSumOverflow(t *testing.T) {
    _, GotErr := StackMachine("25000 25001 SUM")
    WantErr := errors.New("Overflow Error")
    if GotErr.Error() != WantErr.Error() {
        t.Error("Expected", WantErr.Error(), "Got", GotErr.Error())
    }
}

func TestMultiplyOverflow(t *testing.T) {
    _, GotErr := StackMachine("250 201 *")
    WantErr := errors.New("Overflow Error")
    if GotErr.Error() != WantErr.Error() {
        t.Error("Expected", WantErr.Error(), "Got", GotErr.Error())
    }
}

func TestAddOverflow(t *testing.T) {
    _, GotErr := StackMachine("25000 25001 +")
    WantErr := errors.New("Overflow Error")
    if GotErr.Error() != WantErr.Error() {
        t.Error("Expected", WantErr.Error(), "Got", GotErr.Error())
    }
}


//func TestUnderflowMinus(t *testing.T){}









/*func TestMultipleOperations(t *testing.T){
	result, GotErr:=StackMachine("1 2 DUP + 3 + POP")
	if GotErr!= nil{
		t.Error("Unexpected Error", GotErr.Error())
	}
	want:= 6
	if result!=want {
		t.Error("Expected",want,"Got",result)
	}
} */ 
																					//need to fix
/*func TestPopAfterAdd(t *testing.T){
	result, GotErr := StackMachine("1 2 + POP")

	if GotErr!=nil{
		t.Error("Unexpected Error", GotErr.Error())
	}
	want:=3
	if result!= want{
		t.Error("Expected",want,"Got",result)
	}
} */

func TestAddAfterDup(t *testing.T){
	result, GotErr :=StackMachine("1 DUP +")
	if GotErr!=nil{
		t.Error("Unexpected Error", GotErr.Error())
	}
	want:=2
	if result!= want{
		t.Error("Expected",want,"Got",result)
	}
}

func TestUnderflow(t *testing.T) {
	_, GotErr := StackMachine("5 2 -")
    WantErr := errors.New("Underflow Error")
    if GotErr.Error() != WantErr.Error() {
        t.Error("Expected", WantErr.Error(), "Got", GotErr.Error())
    }
}




/*func TestIfTwoSymbolsValidCommand(t *testing.T){

	_, GotErr:= StackMachine("+ -")

	want:= errors.New("Invalid Command")

	if GotErr.Error()!=want.Error(){
		t.Error("Wanted",want,"But didnt get Invalid Command prompt")
	}
}

func TestCommandSymbolAndWord(t *testing.T){

	_, GotErr:= StackMachine("+ DUP")

	want:= errors.New("Invalid Command")

	if GotErr.Error()!=want.Error(){
		t.Error("Wanted",want,"But didnt get Invalid Command prompt")
	}
}

*/

/*func TestTwoFewAdd(t *testing.T){
	_, gotErr:= StackMachine("99 +")

	want:= errors.New("Too few elements to add")

	if gotErr.Error()!=want.Error(){
		t.Error("Expected error due to test having few elements to add")
	}
} 

 
/*
  All these tests must pass for completion
*/
func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name string
		commands string
		expected int
		expectedErr error		
	}{
		{name: "empty error", commands: "", expected: 0, expectedErr: errors.New("empty error")},
		{name: "add overflow", commands: "50000 DUP +", expected: 0, expectedErr: errors.New("Overflow Error")},
		{name: "too few add", commands: "99 +", expected: 0, expectedErr: errors.New("Too Few Elements")},
		{name: "too few minus", commands: "99 -", expected: 0, expectedErr: errors.New("Too Few Elements")},
		{name: "too few multiply", commands: "99 *", expected: 0, expectedErr: errors.New("Too Few Elements")},
		{name: "empty stack", commands: "99 CLEAR", expected: 0, expectedErr: errors.New("Empty Stack")},
		{name: "sum single value", commands: "99 SUM", expected: 99, expectedErr: nil},
		{name: "sum empty", commands: "SUM", expected: 0, expectedErr: errors.New("Empty Stack")},
		{name: "normal +*", commands: "5 6 + 2 *", expected: 22, expectedErr: nil},
		{name: "clear too few", commands: "1 2 3 4 + CLEAR 12 +", expected: 0, expectedErr: errors.New("Empty Stack")},
		{name: "normal after clear", commands: "1 CLEAR 2 3 +", expected: 5, expectedErr: nil},
		{name: "single integer", commands: "9876", expected: 9876, expectedErr: nil},
		{name: "invalid command", commands: "DOGBANANA", expected: 0, expectedErr: errors.New("Invalid Command")},
		{name: "normal +-*", commands: "5 9 DUP + + 43 - 3 *", expected: 60, expectedErr: nil},
		{name: "minus", commands: "2 5 -", expected: 3, expectedErr: nil},
		{name: "underflow minus", commands: "5 2 -", expected: 0, expectedErr: errors.New("Underflow Error")},
		{name: "at overflow limit", commands: "25000 DUP +", expected: 50000, expectedErr: nil},
		{name: "at overflow limit single value", commands: "50000 0 +", expected: 50000, expectedErr: nil},
		{name: "overflow plus", commands: "50000 1 +", expected: 0, expectedErr: errors.New("Overflow Error")},
		{name: "overflow single value", commands: "50001", expected: 0, expectedErr: errors.New("Overflow Error")},
		{name: "times zero at overflow limit", commands: "50000 0 *", expected: 0, expectedErr: nil},
		{name: "too few at first", commands: "1 2 3 4 5 + + + + * 999", expected: 0, expectedErr: errors.New("Too Few Elements")},
		{name: "normal simple", commands: "1 2 - 99 +", expected: 100, expectedErr: nil},
		{name: "at overflow minus to zero", commands: "50000 50000 -", expected: 0, expectedErr: nil},
		{name: "clear empties stack", commands: "CLEAR", expected: 0, expectedErr: errors.New("Empty Stack")},
		{name: "normal sum", commands: "3 4 3 5 5 1 1 1 SUM", expected: 23, expectedErr: nil},
		{name: "sum after clear stack", commands: "3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8, expectedErr: nil},
		{name: "sum then too few", commands: "3 4 3 5 5 1 1 1 SUM -", expected: 0, expectedErr: errors.New("Too Few Elements")},
		{name: "fibonacci", commands: "1 2 3 4 5 * * * *", expected: 120, expectedErr: nil},
		
	}

	for _, test := range tests {
			
		got, err := StackMachine(test.commands)

		if (test.expectedErr != nil) {
			if err == nil {
				t.Errorf("%s (%s) Expected error, but received nil", test.name, test.commands)
			} else if err.Error() != test.expectedErr.Error()  {
				t.Errorf("%s (%s) got error %v, want %v", test.name, test.commands, err, test.expectedErr)
			}
		}  else if got != test.expected {
			t.Errorf("%s (%s) got %v, want %v", test.name, test.commands, got, test.expected)
		}
	}
} 