import java.io.IOException
import java.io.LineNumberReader
import java.io.Reader
import java.lang.StringBuilder
import java.util.regex.Matcher
import java.util.regex.Pattern

private const val regexPat =
    "\\s*((//.*)|([0-9]+)|(\"(\\\\\"|\\\\\\\\|\\\\n|[^\"])*\")" + "|[A-Za-z][A-Za-z0-9]*|==|<=|>=|&&|\\|\\||\\p{Punct})?"

open class Lexer(
    r: Reader
) {
    private var hasMore: Boolean = true
    private val reader: LineNumberReader = LineNumberReader(r)
    private val queue = mutableListOf<Token>()


    companion object {
        private val pattern = Pattern.compile(regexPat)
        private fun toStringLiteral(s: String): String {
            val sb = StringBuilder()
            val len = s.length - 1
            var i = 1
            while (i < len) {
                var c = s[i]
                if (c == '\\' && i + 1 < len) {
                    val c2 = s[i + 1]
                    if (c2 == '"' || c2 == '\\') {
                        c = s[++i]
                    } else if (c2 == 'n') {
                        ++i
                        c = '\n'
                    }
                }
                sb.append(c)
                i++
            }
            return sb.toString()
        }
    }

    @Throws(ParseException::class)
    fun read(): Token {
        return if (fillQueue(0)) {
            queue.removeAt(0)
        } else {
            return Token.EOF
        }
    }

    @Throws(ParseException::class)
    fun peek(i: Int): Token {
        return if (fillQueue(i)) {
            queue[i]
        } else Token.EOF

    }

    @Throws(ParseException::class)
    private fun fillQueue(i: Int): Boolean {
        while (i >= queue.size) if (hasMore) {
            readLine()
        } else {
            return false
        }
        return true
    }

    @Throws(ParseException::class)
    protected fun readLine() {
        val line: String?
        try {
            line = reader.readLine()
        } catch (e: IOException) {
            throw ParseException(e)
        }

        if (line == null) {
            hasMore = false
            return
        }

        val lineNo = reader.lineNumber
        val matcher = pattern.matcher(line)
        matcher.useTransparentBounds(true).useAnchoringBounds(false)

        var pos = 0
        val endPos = line.length

        while (pos < endPos) {
            matcher.region(pos, endPos)
            if (matcher.lookingAt()) {
                addToken(lineNo, matcher)
                pos = matcher.end()
            } else throw ParseException("bad token at line $lineNo")
        }
        queue.add(IdToken(lineNo, Token.EOL))
    }

    private fun addToken(lingNo: Int, matcher: Matcher) {
        val m = matcher.group(1)
        if (m != null) {
            if (matcher.group(2) == null) {
                val token: Token = when {
                    matcher.group(3) != null -> {
                        NumToken(lingNo, Integer.parseInt(m))
                    }
                    matcher.group(4) != null -> {
                        StrToken(lingNo, toStringLiteral(m))
                    }
                    else -> {
                        IdToken(lingNo, m)
                    }
                }
                queue.add(token)
            }
        }
    }

    private inner class NumToken(line: Int, val value: Int) : Token(line) {
        override fun isNumber() = true
        override fun getText() = value.toString()
        override fun getNumber(): Number = value
    }

    private inner class StrToken(line: Int, val literal: String) : Token(line) {
        override fun isString(): Boolean = true
        override fun getText(): String = literal
    }

    private inner class IdToken(line: Int, val id: String) : Token(line) {
        override fun isIdentifier(): Boolean = true
        override fun getText(): String = id
    }
}