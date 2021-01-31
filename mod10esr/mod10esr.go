package mod10esr

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"math"
	"strconv"
	"strings"
)

//Flogo declaraion functions and initilization

type mod10esr struct {
}

func init() {
	 _ = function.Register(&mod10esr{})
}

func (s *mod10esr) Name() string {
	return "mod10esr"
}

func (s *mod10esr) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString,data.TypeString, data.TypeString}, false
}

//Eval(BILLING_ESR_MOD10 string, BILLING_ESR_PREFIX string, customerNumber string, invoiceNumber string) (string, error)
func (s *mod10esr) Eval(params ...interface{}) (interface{}, error) {
	BILLING_ESR_MOD10,err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Input paramater must be specified, BILLING_ESR_MOD10 [%+v] must be string", params[0])
	}
	BILLING_ESR_PREFIX,err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("Input paramater must be specified, BILLING_ESR_PREFIX [%+v] must be string", params[1])
	}
	customerNumber,err := coerce.ToString(params[2])
	if err != nil {
		return nil, fmt.Errorf("Input paramater must be specified, customerNumber [%+v] must be string", params[2])
	}
	invoiceNumber,err := coerce.ToString(params[3])
	if err != nil {
		return nil, fmt.Errorf("Input paramater must be specified, invoiceNumber [%+v] must be string", params[3])
	}

	fmt.Println("mod10esrValue function has been executed for the following values, BILLING_ESR_MOD10=",BILLING_ESR_MOD10, " , and BILLING_ESR_PREFIX=", BILLING_ESR_PREFIX,
		", CustomerNumber=",customerNumber , "invoiceNumber=",invoiceNumber)
	fullString := BILLING_ESR_PREFIX + PadLeft(customerNumber,"0",10) + PadLeft(invoiceNumber,"0", 10)
	slicedBillingMod10Esr := CreateSliceEsrMod10(BILLING_ESR_MOD10)

	var esr = 0
	for i := 0; i < len(fullString); i++ {
	fullStringRune := []rune(fullString) //converted to rune just for the subsctring
	partOfFullString, err := strconv.Atoi(string(fullStringRune[i:i+1]))
	if err != nil {
	fmt.Println("Error during conversion of the FULL Constructed sting to Integer, error details: ", err)
}

	esr = int(math.Mod(float64(esr+partOfFullString), 10))
	fmt.Println("Calculated ESR Value--> ", esr)
	esr, err = strconv.Atoi(slicedBillingMod10Esr[esr])
	if err != nil {
	fmt.Println("Error during conversion of the FULL Constructed sting to Integer, error details: ", err)
}
	fmt.Println("esr value --> ", esr)
	if esr == 0 {
	esr = 0
}
}
	esr = int(math.Mod(10 - float64(esr),10))
	returnValue := fullString + strconv.Itoa(esr)
	fmt.Println("Final invoice calculation with the ESR digit is --> ", returnValue)
	return returnValue,nil
}

func CreateSliceEsrMod10(BILLING_ESR_MOD10 string) []string {
	// hardcoded the value 10 as according to modulo 10 esr algorithim, the max filler/carry is 10
	billingEsrMod10 := strings.SplitN(BILLING_ESR_MOD10,"",10)
	return billingEsrMod10
}

func PadLeft(str, pad string, lenght int) string {
	for {
		if len(str) >= lenght {
			return str[0:lenght]
		} else {
			str = pad + str
		}
	}
}