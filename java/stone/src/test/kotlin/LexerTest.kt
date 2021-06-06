import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.flow.flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.runBlocking
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals


internal class LexerTest {
    @Test
    fun givenACodeStringItWillParseAsTokenList() = runBlocking {
        val InputCode = """while i < 10 {
                sum = sum + i
                i=i+1
            }
            sum"""
        val l = Lexer(InputCode.reader())


        val tokenList = flow<String> {
            do {
                val t: Token = l.read()
                emit(t.getText())
            } while (t != Token.EOF)
        }.toList()

        assertEquals(
            listOf(
                "while",
                "i",
                "<",
                "10",
                "{",
                "\\n",
                "sum",
                "=",
                "sum",
                "+",
                "i",
                "\\n",
                "i",
                "=",
                "i",
                "+",
                "1",
                "\\n",
                "}",
                "\\n",
                "sum",
                "\\n",
                ""   // EOF text
            ),
            tokenList
        )
    }

}