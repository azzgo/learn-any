import java.io.IOException
import java.lang.Exception

class ParseException : Exception {

    constructor(msg: String, t: Token) : super("syntax error around ${location(t)}. $msg")
    constructor(msg: String): super(msg)
    constructor(e: IOException): super(e)


    companion object {
        fun location(t: Token):String {
           return if (t == Token.EOF)  {
               "the last line"
           } else {
              "\"${t.getText()}\" at line ${t.lineNumber}"
           }
        }
    }
}