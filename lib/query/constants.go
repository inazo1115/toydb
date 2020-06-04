package query

//
// LexTokenID
//

type LexTokenID int

const (
	// TokenKEY is a key which isn't keyword but like foo or bar.
	TokenKEY LexTokenID = iota
	// TokenVALUE is a value like "foo", 'foo' or 100.
	TokenVALUE

	// Keywords
	TokenALTER
	TokenAND
	TokenAS
	TokenBETWEEN
	TokenBY
	TokenCREATE
	TokenDATABASE
	TokenDELETE
	TokenDISTINCT
	TokenDROP
	TokenEXISTS
	TokenFROM
	TokenGROUP
	TokenHAVING
	TokenIF
	TokenIN
	TokenINDEX
	TokenINNER
	TokenINSERT
	TokenINTO
	TokenJOIN
	TokenLEFT
	TokenLIKE
	TokenLIMIT
	TokenNOT
	TokenORDER
	TokenPASSWORD
	TokenRIGHT
	TokenSELECT
	TokenSET
	TokenSHOW
	TokenTABLE
	TokenTRUNCATE
	TokenUNION
	TokenUPDATE
	TokenUSE
	TokenUSER
	TokenVALUES
	TokenWHERE

	// Symbols
	TokenAT
	TokenGEQ
	TokenLEQ
	TokenNEQ
	TokenEQ
	TokenGT
	TokenLT
	TokenLPAREN
	TokenRPAREN
	TokenLBRACK
	TokenRBRACK
	TokenCOMMA
	TokenPLUS
	TokenMINUS
	TokenTIMES
	TokenDIV
	TokenDIVINT
	TokenMODINT

	// Keywords(Type)
	TokenSTRING
	TokenINT
)
