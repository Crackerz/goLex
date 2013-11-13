package main

import(
	"fmt"
	"bytes"
	"log"
	"io/ioutil"
)

const(
	MAX_LEXEME_LENGTH = 128
)

func main() {
	str := getProgram("example.wb")
	fmt.Printf("Input:\n%s\nOutput:\n",str)

	program := bytes.NewReader(str)
	for i:=0;i<program.Len();i++ {
		lexeme,token:=lex(program)
		fmt.Printf("Next token is: %s Next lexeme is: %s\n",getTokenString(token),lexeme)
	}
}

func lex(program *bytes.Reader) (string,int) {
	lexeme:=make([]rune,0,128)
	index:=0

	//Get next Non-Whitespace Character
	char,_,err:=program.ReadRune()
	for ;charClass(char)==WHITESPACE; {
		errCheck(err)
		char,_,err = program.ReadRune()
	}

	switch charClass(char) {
	case LETTER: //IDENT or KEYWORD
		for ;charClass(char)==LETTER||charClass(char)==DIGIT;char,_,err=program.ReadRune() {
			errCheck(err)
			lexeme=lexeme[:index+1]
			lexeme[index]=char
			index++
		}
		program.UnreadRune()
		result:=string(lexeme)
		return result,getKeyword(result)
	case DIGIT:
		for ;charClass(char)==DIGIT;char,_,err=program.ReadRune() {
			errCheck(err)
			lexeme=lexeme[:index+1]
			lexeme[index]=char
			index++
		}
		program.UnreadRune()
		return string(lexeme),DIGIT
	}
	lexeme=lexeme[:index+1]
	lexeme[index]=char
	return string(lexeme),charClass(char)
}

func getProgram(filename string) (contents []byte) {
	contents,err:=ioutil.ReadFile(filename)
	errCheck(err)
	return
}

func errCheck(err error) {
	if err!=nil {
		log.Fatal(err.Error())
	}
}
