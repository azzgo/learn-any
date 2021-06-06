abstract class Token(var lineNumber: Int) {
    companion object {
        val EOF = object : Token(-1) {}
        val EOL = "\\n"
    }

    open fun isIdentifier(): Boolean = false
    open fun isNumber(): Boolean = false
    open fun isString(): Boolean = false
    open fun getNumber(): Number {
        throw StoneException("not number Token")
    }

    open fun getText(): String = ""
}