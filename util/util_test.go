package util_test

import (
	"database/sql"
	"testing"
	"time"

	u "github.com/bruno5200/CSM/util"
	"github.com/google/uuid"
)

func TestStringToInt(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		s    string
		want int
		test string
	}{
		{
			"",
			0,
			"empty string",
		},
		{
			"1",
			1,
			"valid string to int",
		},
		{
			"2b",
			0,
			"invalid number",
		},
		{
			"2b%",
			0,
			"invalid string with special characters",
		},
		{
			"A2",
			0,
			"invalid string with letters",
		},
		{
			"1000",
			1000,
			"valid",
		},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.test)
		got, err := u.StringToInt(test.s)
		if err != nil {
			t.Logf("StringToInt(%s) = %v, want %v", test.s, got, test.want)
		}
		if got != test.want {
			t.Errorf("StringToInt(%s) = %v, want %v", test.s, got, test.want)
		}
	}
}

func TestValidAddress(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		email string
		ok    bool
		test  string
	}{
		{
			"",
			false,
			"empty address",
		},
		{
			"bruno",
			false,
			"invalid email",
		},
		{
			"bruno@",
			false,
			"invalid email format",
		},
		{
			"bruno bruno",
			false,
			"invalid email format",
		},
		{
			"bruno@bruno",
			false,
			"invalid email format",
		},
		{
			"bgutierrezdatec.com.bo",
			false,
			"valid email",
		},
		{
			"bgutierrez@datec.com.bo",
			true,
			"valid email",
		},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.test)
		ok := u.ValidAddress(test.email)
		if ok != test.ok {
			t.Errorf("ValidAddress(%s) = %v, want %v", test.email, ok, test.ok)
		}
	}
}

func TestValidCompleteName(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		name string
		ok   bool
		test string
	}{
		{
			"",
			false,
			"empty name",
		},
		{
			"John Doe Smith",
			true,
			"valid name",
		},
		{
			"John",
			false,
			"invalid name",
		},
		{
			"John  Doe",
			false,
			"invalid name with 2 spaces",
		},
		{
			"JohnDoeDoe",
			false,
			"invalid name without space",
		},
		{
			"John Doe^&*%",
			false,
			"invalid name with special characters",
		},
		{
			"John2344 Smith Doe",
			false,
			"invalid name with numbers",
		},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.test)
		ok := u.ValidCompleteName(test.name)
		if ok != test.ok {
			t.Errorf("ValidCompleteName(%s) = %v, want %v", test.name, ok, test.ok)
		}
	}
}

func TestValidPhone(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		phone string
		ok    bool
		err   string
	}{
		{
			"",
			false,
			"empty phone number",
		},
		{
			"123",
			false,
			"is too short",
		},
		{
			"12345678",
			true,
			"valid phone number",
		},
		{
			"12345678901234567890",
			false,
			"is too long",
		},
		{
			"abcdefghijk",
			false,
			"invalid phone number",
		},
		{
			"12345 67890",
			false,
			"phone number with space",
		},
		{
			"12345      67890",
			false,
			"phone number with spaces",
		},
		{
			"12345-67890",
			false,
			"phone number with dash",
		},
		{
			"1234%^&*()12345",
			false,
			"phone number with special characters",
		},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.err)
		ok := u.ValidPhone(test.phone)
		if ok != test.ok {
			t.Errorf("ValidPhone(%s) = %v, want %v", test.phone, ok, test.ok)
		}
	}
}

func TestValidName(t *testing.T) {

	// Table Driven test
	var tests = []struct {
		name string
		ok   bool
		test string
	}{
		{
			"",
			false,
			"empty name",
		},
		{
			"123",
			false,
			"numbers",
		},
		{
			"b",
			false,
			"too short",
		},
		{
			"abcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijk",
			false,
			"too long",
		},
		{
			"abcdef  ghijk",
			false,
			"contains double spaces",
		},
		{
			"abcdefghijk kjdbvaibds",
			true,
			"valid name",
		},
		{
			"afg^&*%",
			false,
			"invalid name",
		},
		{
			"sf2344 jhsvfyu",
			false,
			"invalid name",
		},
		{
			"Lucy Ecos",
			true,
			"valid name",
		},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.test)
		ok := u.ValidName(test.name)
		if ok != test.ok {
			t.Errorf("ValidName(%s) = %v, want %v", test.name, ok, test.ok)
		}
	}
}

func TestValidUser(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		user string
		ok   bool
		test string
	}{
		{
			"",
			false,
			"empty user",
		},
		{
			"123",
			false,
			"numbers",
		},
		{
			"b",
			false,
			"too short",
		},
		{
			"abcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijk",
			false,
			"too long",
		},
		{
			"abcdef  ghijk",
			false,
			"contains double spaces",
		},
		{
			"abcdefghijk kjdbvaibds",
			false,
			"valid user",
		},
		{
			"afg^&*%",
			false,
			"invalid user",
		},
		{
			"sf2344 jhsvfyu",
			false,
			"invalid user",
		},
		{
			"lecos",
			true,
			"valid user",
		},
		{
			"CGF-SCZ-VIP-001",
			true,
			"valid generated user",
		},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.test)
		ok := u.ValidUser(test.user)
		if ok != test.ok {
			t.Errorf("ValidUser(%s) = %v, want %v", test.user, ok, test.ok)
		}
	}
}

func TestValidBussines(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		bussiness string
		ok        bool
		test      string
	}{
		{
			"",
			false,
			"empty bussiness",
		},
		{
			"123",
			true,
			"numbers",
		},
		{
			"b",
			false,
			"too short",
		},
		{
			"abcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijkabcdefghijk",
			false,
			"too long",
		},
		{
			"abcdef  ghijk",
			false,
			"contains double spaces",
		},
		{
			"abcdefghijk kjdbvaibds",
			true,
			"valid bussiness",
		},
		{
			"afg^&*%",
			false,
			"invalid bussiness",
		},
		{
			"sf2344 jhsvfyu",
			true,
			"invalid bussiness",
		},
		{
			"lucy_ecos",
			true,
			"valid bussiness",
		},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.test)
		ok := u.ValidBussines(test.bussiness)
		if ok != test.ok {
			t.Errorf("ValidBussiness(%s) = %v, want %v", test.bussiness, ok, test.ok)
		}
	}
}

func TestValidPassword(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		pass string
		ok   bool
		test string
	}{
		{
			"",
			false,
			"empty password",
		},
		{
			"1SalchichaDeSanJuan*",
			true,
			"valid password",
		},
		{
			"2SalchichaDeSanJuan*",
			true,
			"2nd valid password",
		},
		{
			"ABCDEF",
			false,
			"too short with uppercase",
		},
		{
			"abc",
			false,
			"too short with lowecase",
		},
		{
			"123",
			false,
			"too short with numbers",
		},
		{
			"1234567890ABCDEF",
			false,
			"no special characters",
		},
		{
			"1234567890abCDEF",
			false,
			"no special characters",
		},
		{
			"a123456789012345^&*()ABCDEF",
			false,
			"too long",
		},
		{
			"a1234567890^",
			false,
			"no uppercase",
		},
		{
			"1234567890!$ABC",
			false,
			"no lowercase",
		},
		{
			"a1234567890 1^A",
			false,
			"contains space",
		},
	}
	for _, test := range tests {
		ok := u.ValidPassword(test.pass)
		t.Logf("Testing %s", test.test)
		if ok != test.ok {
			t.Errorf("ValidPassword(%s) = %v, want %v", test.pass, ok, test.ok)
		}
	}
}

func TestValidDate(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		date string
		ok   bool
		test string
	}{
		{"21/12/2018", true, "Valid Date"},
		{"21/12/2018 12:00:00", false, "Invalid Date"},
		{"21/12/2018 12:00", false, "Invalid Date"},
		{"21-12-2018 12:00:00", false, "Invalid Date"},
		{"21-12-2018 12:00", false, "Invalid Date"},
		{"veintiuno de diciembre", false, "Invalid Date"},
		{"doce/diciembre/dosmildieciocho", false, "Invalid Date"},
		{"@!/!@/@)!*", false, "Valid Date"},
	}

	for _, test := range tests {
		ok := u.ValidDate(test.date)
		t.Logf("Testing %s", test.test)
		if ok != test.ok {
			t.Errorf("ValidDate(%s) = %v, want %v", test.date, ok, test.ok)
		}
	}
}

func TestValidTime(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		Time string
		ok   bool
		test string
	}{
		{"12:00", true, "Valid Time"},
		{"12:00:00", false, "Invalid Time"},
		{"12000", false, "Invalid Time wihout :"},
		{"doce:0", false, "Invalid Time with characters"},
		{"!@:))", false, "Invalid Time with symbols"},
	}
	for _, test := range tests {
		ok := u.ValidTime(test.Time)
		t.Logf("Testing %s", test.test)
		if ok != test.ok {
			t.Errorf("ValidTime(%s) = %v, want %v", test.Time, ok, test.ok)
		}
	}
}

func TestValidSKU(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		sku  string
		ok   bool
		test string
	}{
		{
			"",
			false,
			"empty sku",
		},
		{
			"123",
			false,
			"invalid sku",
		},
		{
			"12345678901234567890",
			false,
			"too long",
		},
		{
			"JNBF-12",
			true,
			"valid sku",
		},
		{
			"afg^&*%",
			false,
			"invalid sku",
		},
		{
			"34256-NGJ",
			true,
			"valid sku",
		},
		{
			"Lucy Ecos",
			false,
			"invalid sku",
		},
	}

	for _, test := range tests {
		t.Logf("Testing %s", test.test)
		ok := u.ValidSKU(test.sku)
		if ok != test.ok {
			t.Errorf("ValidSKU(%s) = %t, want %t", test.sku, ok, test.ok)
		}
	}
}

func TestGetOTP(t *testing.T) {
	otp := u.GetOTP()
	t.Logf("OTP: %s", otp)
	if len(otp) != 6 {
		t.Errorf("GetOtp() = %s, want %s and get %d digits", otp, "6 digits", len(otp))
	}
}

func TestHashPassword(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		pass string
		hash string
		test string
	}{
		{
			"bTQG$EM3DKEMN",
			"$2a$04$uL6Hg5wgoG7OUBGrhajjuOe.VLcIP2lfIkm96Mot.yWtOz6Hr7NSa",
			"Valid Password User",
		},
		{
			"bTQG$EM3DKEMZ",
			"$2a$04$Tpl2DipuWkgRxDhvaUgx/uKVQ9ENO9zLa2DLDQPdoVWp2M1jVXOZa",
			"Valid Password Clerk",
		},
		{
			"bTQG$EM3DKEMA",
			"$2a$04$uyr9ApQ1oT6IVRuIy.pdZebNy.7Uvt2n8LZk1Sm.sK1Dmjb8P0iLa",
			"Valid Password Admin",
		},
		{
			"hAbECKPS6Q6W3i",
			"$2a$14$slbtnKMHfC9GlklkndY1XeoIJIFtIby7ggTnPXYhHMCRNfES01EXa",
			"Secure Password",
		},
		{
			"bPGCDqd9",
			"$2a$14$y9Eb7Jm9nKEE4EJmURX7ZO54uFa/gyjF6HrtsB03K2OhhiXvpeWiK",
			"Maria Suarez",
		},
		{
			"I5PTtdNm",
			"$2a$14$FHIPOLGwqYGKJJE3sTqqpOy1AjxXb942v7fqydnHhdlgIltaJDm9O",
			"Ximena Correa",
		},
		{
			"HsL2h5n*DD",
			"$2a$12$Gf6NdTsp40kCTjOc62.wk.Gzu6E4wNM4UoDVGYrxQqRuV.E0dnbqe",
			"CGF-SCZ-001",
		},
	}
	for _, test := range tests {
		var err error
		test.hash, err = u.HashPassword(test.pass)
		if err != nil {
			t.Errorf("HashPassword(%s) = %v, want %v", test.pass, err, nil)
		}
		t.Logf("password: %s, hash %v", test.pass, test.hash)
		t.Logf("Testing %s", test.test)
		if !u.CheckPasswordHash(test.pass, test.hash) {
			t.Errorf("HashPassword = %s, returned %v", test.pass, test.hash)
		}
	}
}

func TestHashPasswordAPP(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		pass string
		hash string
		test string
	}{
		{"dev", "$2a$04$wr1QEAtXtdB6XyIHV2t2BO1EXh3m0pbOwCywLDms233kelG3WSQ2S", "dev"},
		{"test", "$2a$04$qaUypaRj1Z/bgJFg.9KytOFez/XjJMWRhxd37vpzFPD2nYz9bEAQW", "test"},
		{"test1", "$2a$04$yAKI9cSJyBtR6P7bt8yO7uV/L4dNbhux49tUQRmPmDf9Z2E/RVCiO", "test1"},
		{"mier1", "$2a$04$pq.ckFLbnLtuHMkRAt1zeOkis/yxiwWxvKQaszOz9MHkuk5lFwfu6", "mier1"},
		{"mier2", "$2a$04$2AvZVzgKujV/G307iCrTH.2Agivdk.V4KSDUQBJA4Az5aHbVNx0Zi", "mier2"},
		{"virt1", "$2a$04$so0XIAgkNuda3kr0XGQrHugbcPfZ4wncoAW63UP.zwPoXHePJH5gi", "virt1"},
		{"dsoruco", "$2a$04$jZMAqsPEIwcgrmUNUrU5F.EfW8PX5cRin.HF8z7dodgosysMdl9A2", "diego"},
		{"nyañez", "$2a$04$ybX8n/i6mg0q2X57f/rGcuk9zlJPBvvW39kNqq6fAM669pN2sQ3QC", "nahely"},
		{"ShowdeMier", "$2a$04$hakBdsmHR2q8EFtqR0yTO.WJAW7X5p0o74P9q6Pe6BE36xRAJNWJ2", "showdemier"},
		{"pL6U$4a271", "$2a$04$JLnnpQLSheJad/J5anHWIuRJRdW9qqHOuHcWVLy3xEqhJoaFBr6bO", "pL6U$4a271"},
		{"P*Lk94MYq3", "$2a$04$EWCph0xFvZoYH9MPl/MiSeojctpAbKhECEYItBvkng1ShX8oSZN26", "P*Lk94MYq3"},
		{"L1h7Z7pD$X", "$2a$04$fPBaDUTDNeJce8e71u3/6.x68eMzOim8sWAztnQ4pjsr4untruvg2", "L1h7Z7pD$X"},
		{"3s7TJ@4lXy", "$2a$04$Hk5g5IEgSHcIfCgd9cF1ReXCrCsskLNAPL4r2xPxxCm9ptxOzGUcW", "3s7TJ@4lXy"},
		{"P68zFu6I*h", "$2a$04$hB6nnu1/AeA5wXbHVMDjG.hCFf4F/rZKQqZ.1s6P3oJQAG5tQcaX6", "P68zFu6I*h"},
		{"B8z01Ng6@g", "$2a$04$SuKy4waK3oMFjeA4h.e4Ruxanb1EXACER0yY0gVntPONDizyfdHE6", "B8z01Ng6@g"},
	}
	for _, test := range tests {
		var err error
		test.hash, err = u.HashPassword(test.pass)
		if err != nil {
			t.Errorf("Error %s on HashPassword %s", err, test.pass)
		}
		t.Logf("%s, %s", test.pass, test.hash)
		// t.Logf("Testing %s", test.test)
		if !u.CheckPasswordHash(test.pass, test.hash) {
			t.Errorf("HashPassword = %s, returned %s", test.pass, test.hash)
		}
	}
}

func TestGenereatePassword(t *testing.T) {
	var tests = []struct {
		pass   string
		length int
	}{
		{
			length: 10,
		},
	}
	for _, test := range tests {
		test.pass = u.GeneratePassword(test.length)
		t.Logf("Generated Password: %s", test.pass)
		if len(test.pass) != test.length {
			t.Errorf("GeneratePassword() = %s, want %s and get %d digits", test.pass, "10-18 digits", len(test.pass))
		}
	}
}

func TestGenerateQrCode(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id       uuid.UUID
		stringid string
	}{
		{
			id:       uuid.MustParse("c0a80121-7ac0-11e8-9c9c-2d42b21b1a3e"),
			stringid: "c0a80121-7ac0-11e8-9c9c-2d42b21b1a3e",
		},
	}

	for _, test := range tests {
		err := u.GenerateQRCode(test.stringid)
		t.Logf("Generated Qr: %s", test.id)
		if err != nil {
			t.Errorf("GenerateQr(%s Error: %v", test.id, err)
		}
	}
}

func TestGenerateWhiteQrCode(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id       uuid.UUID
		stringid string
	}{
		{
			id:       uuid.MustParse("c0a80121-7ac0-11e8-9c9c-2d42b21b1a3e"),
			stringid: "c0a80121-7ac0-11e8-9c9c-2d42b21b1a3e",
		},
	}

	for _, test := range tests {
		err := u.GenerateWhiteQRCode(test.stringid)
		t.Logf("Generated Qr: %s", test.id)
		if err != nil {
			t.Errorf("GenerateQr(%s Error: %v", test.id, err)
		}
	}
}

func TestGenerateGreyQrCode(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id       uuid.UUID
		stringid string
	}{
		{
			id:       uuid.MustParse("c0a80121-7ac0-11e8-9c9c-2d42b21b1a3e"),
			stringid: "c0a80121-7ac0-11e8-9c9c-2d42b21b1a3e",
		},
	}

	for _, test := range tests {
		err := u.GenerateGreyQRCode(test.stringid)
		t.Logf("Generated Qr: %s", test.id)
		if err != nil {
			t.Errorf("GenerateQr(%s Error: %v", test.id, err)
		}
	}
}

func TestValidOrderName(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		order string
		ok    bool
		test  string
	}{
		{
			"#1200",
			true,
			"Valid Order Name",
		},
		{
			"#0000",
			true,
			"Valid Zero Order Name",
		},
		{
			"#9999",
			true,
			"Invalid Max Order Name",
		},
		{
			"#1000097457",
			false,
			"Invalid 11 digits Order Name",
		},
		{
			"#10000$",
			false,
			"Invalid Order Name With Simbol",
		},
		{
			"#asdb$%",
			false,
			"Invalid Order Name With characters",
		},
		{
			"ascb",
			false,
			"Invalid Order Name Without #",
		},
		{
			" ",
			false,
			"Invalid Order Name With Space",
		},
	}

	for _, test := range tests {
		ok := u.ValidOrderName(test.order)
		t.Logf("Testing %s", test.test)
		if ok != test.ok {
			t.Errorf("ValidOrderName(%s) = %v, want %v", test.order, ok, test.ok)
		}
	}
}

func TestEncodeQueryParams(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		params map[string]string
		query  string
	}{
		{
			map[string]string{
				"filter": `{"payment_status":"paid", "date_from":"2023-06-01 00:00:00", "date_to":"2023-06-19 23:59:59"}`,
			},
			`filter=%7B%22payment_status%22%3A%22paid%22%2C%20%22date_from%22%3A%222023-06-01%2000%3A00%3A00%22%2C%20%22date_to%22%3A%222023-06-19%2023%3A59%3A59%22%7D`,
		},
		{
			map[string]string{
				"filter": `{"payment_status":"paid", "date_from":"2023-06-01 00:00:00", "date_to":"2023-06-19 23:59:59"}`,
			},
			`filter=%7B%22payment_status%22%3A%22paid%22%2C%20%22date_from%22%3A%222023-06-01%2000%3A00%3A00%22%2C%20%22date_to%22%3A%222023-06-19%2023%3A59%3A59%22%7D`,
		},
	}
	for _, test := range tests {
		query := u.EncodeQueryParams(test.params)
		t.Logf("Testing %s", test.query)
		if query != test.query {
			t.Errorf("EncodeQueryParams(%v) = \n%s, \nwant \n%s", test.params, query, test.query)
		}
	}
}

func TestStringtoInt64(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		s    string
		want int64
		test string
	}{
		{
			"",
			0,
			"empty string",
		},
		{
			"1",
			1,
			"valid string to int",
		},
		{
			"2b",
			0,
			"invalid number",
		},
		{
			"2b%",
			0,
			"invalid string with special characters",
		},
		{
			".",
			0,
			"invalid string with point",
		},
		{
			"2.0",
			20,
			"invalid string number with point",
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.test, func(t *testing.T) {
			t.Parallel()
			got := u.StringToInt64(test.s)
			if got != test.want {
				t.Errorf("StringToInt(%s) = %v, want %v", test.s, got, test.want)
			}
			t.Logf("Result %d", got)
		})
	}
}

func TestFormatDateWithYear(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		date string
		want string
	}{
		{
			date: "21/12/2018",
			want: "21 de Diciembre, 2018",
		},
		{
			date: "14/11/2023",
			want: "14 de Noviembre, 2023",
		},
		{
			date: "25/10/2023",
			want: "25 de Octubre, 2023",
		},
		{
			date: "17/09/2023",
			want: "17 de Septiembre, 2023",
		},
		{
			date: "10/08/1997",
			want: "10 de Agosto, 1997",
		},
		{
			date: "01/07/2004",
			want: "01 de Julio, 2004",
		},
		{
			date: "02/06/2005",
			want: "02 de Junio, 2005",
		},
		{
			date: "03/05/2006",
			want: "03 de Mayo, 2006",
		},
		{
			date: "04/04/2007",
			want: "04 de Abril, 2007",
		},
		{
			date: "05/03/2008",
			want: "05 de Marzo, 2008",
		},
		{
			date: "06/02/2009",
			want: "06 de Febrero, 2009",
		},
		{
			date: "07/01/2010",
			want: "07 de Enero, 2010",
		},
	}

	for i := range tests {
		test := tests[i]

		t.Run(test.date, func(t *testing.T) {
			t.Parallel()

			date, _ := time.Parse("02/01/2006", test.date)
			got := u.FormatDateWithYear(date)
			t.Logf("Testing %s", test.date)
			if got != test.want {
				t.Errorf("FormatDate(%s) = %s, want %s", test.date, got, test.want)
			}
		})
	}
}

func TestFormatDateWithoutYear(t *testing.T) {

	// Table Driven Test
	var tests = []struct {
		date string
		want string
	}{
		{
			date: "21/12/2018",
			want: "21 de Diciembre",
		},
		{
			date: "14/11/2023",
			want: "14 de Noviembre",
		},
		{
			date: "25/10/2023",
			want: "25 de Octubre",
		},
		{
			date: "17/09/2023",
			want: "17 de Septiembre",
		},
		{
			date: "10/08/1997",
			want: "10 de Agosto",
		},
		{
			date: "01/07/2004",
			want: "01 de Julio",
		},
		{
			date: "02/06/2005",
			want: "02 de Junio",
		},
		{
			date: "03/05/2006",
			want: "03 de Mayo",
		},
		{
			date: "04/04/2007",
			want: "04 de Abril",
		},
		{
			date: "05/03/2008",
			want: "05 de Marzo",
		},
		{
			date: "06/02/2009",
			want: "06 de Febrero",
		},
		{
			date: "07/01/2010",
			want: "07 de Enero",
		},
	}

	for i := range tests {
		test := tests[i]

		t.Run(test.date, func(t *testing.T) {
			t.Parallel()

			date, _ := time.Parse("02/01/2006", test.date)
			got := u.FormatDateWithoutYear(date)
			t.Logf("Testing %s", test.date)
			if got != test.want {
				t.Errorf("FormatDate(%s) = %s, want %s", test.date, got, test.want)
			}
		})
	}
}

func TestFormatDateTime(t *testing.T) {
	date := time.Now()
	t.Log(u.FormatDateTime(date))
	t.Log(date.Format("2006-01-02 15:04:05+00"))
}

func TestFirstIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "111a9044",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "5192a55d",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "81617e7a",
		},
	}
	for _, test := range tests {
		got := u.FisrtIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func TestSecondIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "33c9",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "1db0",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "ce55",
		},
	}
	for _, test := range tests {
		got := u.SecondIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func TestThirdIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "4306",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "4042",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "42ae",
		},
	}
	for _, test := range tests {
		got := u.ThirdIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func TestFourthIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "9feb",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "bd28",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "b440",
		},
	}
	for _, test := range tests {
		got := u.FourthIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func TestFifthIdentifier(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		id   uuid.UUID
		want string
	}{
		{
			id:   uuid.MustParse("111a9044-33c9-4306-9feb-f7501cabca29"),
			want: "f7501cabca29",
		},
		{
			id:   uuid.MustParse("5192a55d-1db0-4042-bd28-ab9b6b2e7744"),
			want: "ab9b6b2e7744",
		},
		{
			id:   uuid.MustParse("81617e7a-ce55-42ae-b440-ea6916ae943a"),
			want: "ea6916ae943a",
		},
	}
	for _, test := range tests {
		got := u.FifthIdentifier(test.id)
		t.Logf("Testing %s", test.id)
		if got != test.want {
			t.Errorf("ShortIdentifier(%s) = %s, want %s", test.id, got, test.want)
		}
	}
}

func TestNullToString(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name  string
		value sql.NullString
		want  string
	}{
		{
			name: "Hello",
			value: sql.NullString{
				String: "Hello",
				Valid:  true,
			},
			want: "Hello",
		},
		{
			name: "Empty",
			value: sql.NullString{
				String: "",
				Valid:  false,
			},
			want: "",
		},
	}
	for _, test := range tests {
		got := u.NullToString(test.value)
		t.Logf("Testing %s", test.name)
		if got != test.want {
			t.Errorf("NullToString(%v) = %s, want %s", test.value, got, test.want)
		}
	}
}

func TestNullString(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name  string
		value string
		want  sql.NullString
	}{
		{
			name:  "Hello",
			value: "Hello",
			want: sql.NullString{
				String: "Hello",
				Valid:  true,
			},
		},
		{
			name:  "Empty",
			value: "",
			want: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
	}
	for _, test := range tests {
		got := u.NullString(test.value)
		t.Logf("Testing %s", test.name)
		if got != test.want {
			t.Errorf("NullString(%s) = %v, want %v", test.value, got, test.want)
		}
	}
}

func TestNullToBool(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name  string
		value sql.NullBool
		want  bool
	}{
		{
			name: "True",
			value: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
			want: true,
		},
		{
			name: "False",
			value: sql.NullBool{
				Bool:  false,
				Valid: true,
			},
			want: false,
		},
		{
			name: "Empty",
			value: sql.NullBool{
				Bool:  false,
				Valid: false,
			},
			want: false,
		},
	}
	for _, test := range tests {
		got := u.NullToBool(test.value)
		t.Logf("Testing %s", test.name)
		if got != test.want {
			t.Errorf("NullToBool(%v) = %t, want %t", test.value, got, test.want)
		}
	}
}

func TestNullBool(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name  string
		value bool
		want  sql.NullBool
	}{
		{
			name:  "True",
			value: true,
			want: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		},
		{
			name:  "False",
			value: false,
			want: sql.NullBool{
				Bool:  false,
				Valid: true,
			},
		},
	}
	for _, test := range tests {
		got := u.NullBool(test.value)
		t.Logf("Testing %s", test.name)
		if got != test.want {
			t.Errorf("NullBool(%t) = %v, want %v", test.value, got, test.want)
		}
	}
}

func TestNullToInt64(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name  string
		value sql.NullInt64
		want  int64
	}{
		{
			name: "1",
			value: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			want: 1,
		},
		{
			name: "0",
			value: sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
			want: 0,
		},
		{
			name: "Empty",
			value: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
			want: 0,
		},
	}
	for _, test := range tests {
		got := u.NullToInt64(test.value)
		t.Logf("Testing %s", test.name)
		if got != test.want {
			t.Errorf("NullToInt(%v) = %d, want %d", test.value, got, test.want)
		}
	}
}

func TestNullInt64(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name  string
		value int64
		want  sql.NullInt64
	}{
		{
			name:  "1",
			value: 1,
			want: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		},
		{
			name:  "0",
			value: 0,
			want: sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
		},
	}
	for _, test := range tests {
		got := u.NullInt64(test.value)
		t.Logf("Testing %s", test.name)
		if got != test.want {
			t.Errorf("NullInt(%d) = %v, want %v", test.value, got, test.want)
		}
	}
}

func TestNullToTime(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name  string
		value sql.NullTime
		want  time.Time
	}{
		{
			name: "2018-06-21 00:00:00",
			value: sql.NullTime{
				Time:  time.Date(2018, 6, 21, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
			want: time.Date(2018, 6, 21, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "0001-01-01 00:00:00",
			value: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			want: time.Time{},
		},
	}
	for _, test := range tests {
		got := u.NullToTime(test.value)
		t.Logf("Testing %s", test.name)
		if got != test.want {
			t.Errorf("NullToTime(%v) = %v, want %v", test.value, got, test.want)
		}
	}
}

func TestNullTime(t *testing.T) {
	// Table Driven Test
	var tests = []struct {
		name  string
		value time.Time
		want  sql.NullTime
	}{
		{
			name:  "2018-06-21 00:00:00",
			value: time.Date(2018, 6, 21, 0, 0, 0, 0, time.UTC),
			want: sql.NullTime{
				Time:  time.Date(2018, 6, 21, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		},
		{
			name:  "0001-01-01 00:00:00",
			value: time.Time{},
			want: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
		},
	}
	for _, test := range tests {
		got := u.NullTime(test.value)
		t.Logf("Testing %s", test.name)
		if got != test.want {
			t.Errorf("NullTime(%v) = %v, want %v", test.value, got, test.want)
		}
	}
}
